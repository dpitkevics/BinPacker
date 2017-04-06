package packed_box

import (
	"github.com/dpitkevics/BoxPacker/packer/box"
	"github.com/dpitkevics/BoxPacker/packer/box_item"
	"math"
)

type PackedBox struct {
	box *box.Box
	items *box_item.ItemList
	weight float64
	remainingWidth float64
	remainingLength float64
	remainingHeight float64
	remainingWeight float64
	usedWidth float64
	useLength float64
	usedHeight float64
}

func NewPackedBox(
box *box.Box,
items *box_item.ItemList,
remainingWidth float64,
remainingLength float64,
remainingHeight float64,
remainingWeight float64,
usedWidth float64,
usedLength float64,
usedHeight float64,
) *PackedBox {
	return &PackedBox{
		box: box,
		items: items,
		remainingWidth: remainingWidth,
		remainingLength: remainingLength,
		remainingHeight: remainingHeight,
		remainingWeight: remainingWeight,
		usedWidth: usedWidth,
		useLength: usedLength,
		usedHeight: usedHeight,
	}
}

func (packedBox *PackedBox) GetBox() *box.Box {
	return packedBox.box
}

func (packedBox *PackedBox) GetItems() *box_item.ItemList {
	return packedBox.items
}

func (packedBox *PackedBox) GetWeight() float64 {
	if packedBox.weight > 0 {
		return packedBox.weight
	}

	packedBox.weight = packedBox.box.EmptyWeight
	items := packedBox.items.Clone()
	for _, item := range items.Items {
		packedBox.weight += item.Weight
	}

	return packedBox.weight
}

func (packedBox *PackedBox) GetRemainingWidth() float64 {
	return packedBox.remainingWidth
}

func (packedBox *PackedBox) GetRemainingLength() float64 {
	return packedBox.remainingLength
}

func (packedBox *PackedBox) GetRemainingHeight() float64 {
	return packedBox.remainingHeight
}

func (packedBox *PackedBox) GetRemainingWeight() float64 {
	return packedBox.remainingWeight
}

func (packedBox *PackedBox) GetUsedWidth() float64 {
	return packedBox.usedWidth
}

func (packedBox *PackedBox) GetUsedLength() float64 {
	return packedBox.useLength
}

func (packedBox *PackedBox) GetUsedHeight() float64 {
	return packedBox.usedHeight
}

func (packedBox *PackedBox) GetVolumeUtilisation() float64 {
	itemVolume := 0.0

	items := packedBox.items.Clone()
	for _, item := range items.Items {
		itemVolume += item.Volume
	}

	return math.Floor((itemVolume / packedBox.box.InnerVolume * 100) + .5)
}
