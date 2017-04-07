package Packer

import (
	"github.com/dpitkevics/BoxPacker/packer/box"
	"github.com/dpitkevics/BoxPacker/packer/box_item"
	"github.com/dpitkevics/BoxPacker/packer/packed_box"
	"math"
	"errors"
)

type Pair struct {
	Key string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

type VolumePacker struct {
	box *box.Box
	items *box_item.ItemList
	lengthLeft float64
	widthLeft float64
	heightLeft float64
	remainingWeight float64
	usedLength float64
	usedWidth float64
	usedHeight float64
}

func NewVolumePacker(box *box.Box, items *box_item.ItemList) *VolumePacker {
	return &VolumePacker{
		box: box,
		items: items,
		heightLeft: box.InnerHeight,
		remainingWeight: box.MaxWeight - box.EmptyWeight,
		widthLeft: box.InnerWidth,
		lengthLeft: box.InnerLength,
	}
}

func (volumePacker *VolumePacker) Pack() *packed_box.PackedBox {
	packedItems := box_item.NewItemList()
	layerLength, layerWidth, layerHeight := 0.0, 0.0, 0.0

	var previousItem *box_item.OrientedItem

	for !volumePacker.items.IsEmpty() {
		itemToPack, _ := volumePacker.items.Extract()

		if itemToPack.Weight > volumePacker.remainingWeight {
			continue
		}

		var nextItem *box_item.Item
		if !volumePacker.items.IsEmpty() {
			nextItem, _ = volumePacker.items.Top()
		}

		orientatedItem, _ := volumePacker.findBestOrientation(itemToPack, previousItem, nextItem, volumePacker.widthLeft, volumePacker.lengthLeft, volumePacker.heightLeft)

		if orientatedItem != nil {
			packedItems.Insert(orientatedItem.GetItem())
			volumePacker.remainingWeight -= itemToPack.Weight

			volumePacker.lengthLeft -= orientatedItem.GetLength()

			layerLength += orientatedItem.GetLength()
			layerWidth = math.Max(orientatedItem.GetWidth(), layerWidth)
			layerHeight = math.Max(layerHeight, orientatedItem.GetHeight())

			volumePacker.usedLength = math.Max(volumePacker.usedLength, layerLength)
			volumePacker.usedWidth = math.Max(volumePacker.usedWidth, layerWidth)

			stackableHeight := layerHeight - orientatedItem.GetHeight()
			volumePacker.tryAndStackItemsIntoSpace(packedItems, previousItem, nextItem, orientatedItem.GetWidth(), orientatedItem.GetLength(), stackableHeight)

			previousItem = orientatedItem

			if nextItem != nil {
				volumePacker.usedHeight += layerHeight
			}
		} else {
			previousItem = nil

			if volumePacker.widthLeft >= math.Min(itemToPack.Width, itemToPack.Length) && volumePacker.isLayerStarted(layerWidth, layerLength, layerHeight) {
				volumePacker.lengthLeft += layerLength
				volumePacker.widthLeft -= layerWidth

				layerWidth = 0
				layerLength = 0

				volumePacker.items.Insert(itemToPack)

				continue
			} else if (volumePacker.lengthLeft < math.Min(itemToPack.Width, itemToPack.Length) || layerHeight == 0) {
				continue
			}

			volumePacker.widthLeft = volumePacker.box.InnerWidth
			if layerWidth > 0 {
				volumePacker.widthLeft = math.Min(math.Floor(layerWidth * 1.1), volumePacker.box.InnerWidth)
			}

			volumePacker.lengthLeft = volumePacker.box.InnerLength
			if layerLength > 0 {
				volumePacker.lengthLeft = math.Min(math.Floor(layerLength * 1.1), volumePacker.box.InnerLength)
			}

			volumePacker.heightLeft -= layerHeight
			volumePacker.usedHeight += layerHeight

			layerWidth = 0
			layerLength = 0
			layerHeight = 0

			volumePacker.items.Insert(itemToPack)
		}
	}

	return packed_box.NewPackedBox(
		volumePacker.box,
		packedItems,
		volumePacker.widthLeft,
		volumePacker.lengthLeft,
		volumePacker.heightLeft,
		volumePacker.remainingWeight,
		volumePacker.usedWidth,
		volumePacker.usedLength,
		volumePacker.usedHeight,
	)
}

func (volumePacker *VolumePacker) findBestOrientation(item *box_item.Item, previousItem *box_item.OrientedItem, nextItem *box_item.Item, widthLeft float64, lengthLeft float64, heightLeft float64) (*box_item.OrientedItem, error) {
	orientations := volumePacker.findPossibleOrientations(item, previousItem, widthLeft, lengthLeft, heightLeft)

	stableOrientations := []*box_item.OrientedItem{}
	unstableOrientations := []*box_item.OrientedItem{}

	for _, orientation := range orientations {
		if orientation.IsStable() {
			stableOrientations = append(stableOrientations, orientation)
		} else {
			unstableOrientations = append(unstableOrientations, orientation)
		}
	}

	orientationsToUse := []*box_item.OrientedItem{}

	if len(stableOrientations) > 0 {
		orientationsToUse = stableOrientations
	} else if len(unstableOrientations) > 0 {
		orientationsInEmptyBox := volumePacker.findPossibleOrientations(
			item,
			previousItem,
			volumePacker.box.InnerWidth,
			volumePacker.box.InnerLength,
			volumePacker.box.InnerHeight,
		)

		stableOrientationsInEmptyBox := []*box_item.OrientedItem{}
		for _, orientation := range orientationsInEmptyBox {
			if orientation.IsStable() {
				stableOrientationsInEmptyBox = append(stableOrientationsInEmptyBox, orientation)
			}
		}

		if nextItem == nil || len(stableOrientationsInEmptyBox) == 0 {
			orientationsToUse = unstableOrientations
		}
	}

	orientationFits := map[int]float64{}
	for i, orientation := range orientationsToUse {
		orientationFit := math.Min(widthLeft - orientation.GetWidth(), lengthLeft - orientation.GetLength())
		orientationFits[i] = orientationFit
	}

	if len(orientationFits) > 0 {
		bestKey := -1
		bestValue := 0.0

		for k, value := range orientationFits {
			if value > bestValue {
				bestKey = k
				bestValue = value
			}
		}

		if bestKey >= 0 {
			return orientationsToUse[bestKey], nil
		}
	}

	return nil, errors.New("No orientation fits")
}

func (volumePacker *VolumePacker) findPossibleOrientations(item *box_item.Item, previousItem *box_item.OrientedItem, widthLeft float64, lengthLeft float64, heightLeft float64) []*box_item.OrientedItem {
	var orientations []*box_item.OrientedItem

	if previousItem != nil && previousItem.GetItem() == item {
		orientations = append(orientations, box_item.NewOrientedItem(item, previousItem.GetLength(), previousItem.GetWidth(), previousItem.GetHeight()))
	} else {
		orientations = append(orientations, box_item.NewOrientedItem(item, item.Length, item.Width, item.Height))
		orientations = append(orientations, box_item.NewOrientedItem(item, item.Width, item.Length, item.Height))
		orientations = append(orientations, box_item.NewOrientedItem(item, item.Length, item.Height, item.Width))
		orientations = append(orientations, box_item.NewOrientedItem(item, item.Width, item.Height, item.Length))
		orientations = append(orientations, box_item.NewOrientedItem(item, item.Height, item.Length, item.Width))
		orientations = append(orientations, box_item.NewOrientedItem(item, item.Height, item.Width, item.Length))
	}

	filteredOrientations := orientations[:0]
	for _, orientation := range orientations {
		if orientation.GetWidth() <= widthLeft && orientation.GetLength() <= lengthLeft && orientation.GetHeight() <= heightLeft {
			filteredOrientations = append(filteredOrientations, orientation)
		}
	}

	return filteredOrientations
}

func (volumePacker *VolumePacker) tryAndStackItemsIntoSpace(packedItems *box_item.ItemList, previousItem *box_item.OrientedItem, nextItem *box_item.Item, maxWidth float64, maxLength float64, maxHeight float64) {
	topItem, _ := volumePacker.items.Top()
	for !volumePacker.items.IsEmpty() && volumePacker.remainingWeight >= topItem.Weight {
		stackedItem, _ := volumePacker.findBestOrientation(topItem, previousItem, nextItem, maxWidth, maxLength, maxHeight)
		if stackedItem != nil {
			volumePacker.remainingWeight -= topItem.Weight
			maxHeight -= stackedItem.GetHeight()

			extractedItem, _ := volumePacker.items.Extract()
			packedItems.Insert(extractedItem)
		} else {
			break
		}
	}
}

func (volumePacker *VolumePacker) isLayerStarted(layerWidth float64, layerLength float64, layerHeight float64) bool {
	return layerWidth > 0 && layerLength > 0 && layerHeight > 0
}
