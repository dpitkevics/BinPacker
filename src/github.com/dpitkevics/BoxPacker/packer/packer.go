package Packer

import (
	"github.com/dpitkevics/BoxPacker/packer/box_item"
	"github.com/dpitkevics/BoxPacker/packer/box"
	"github.com/dpitkevics/BoxPacker/packer/packed_box"
	"errors"
)

type Packer struct {
	items *box_item.ItemList
	boxes *box.BoxList
}

func NewPacker() *Packer {
	return &Packer{
		items: box_item.NewItemList(),
		boxes: box.NewBoxList(),
	}
}

func (packer *Packer) AddItem(item *box_item.Item) {
	packer.items.Insert(item)
}

func (packer *Packer) AddBox(box *box.Box) {
	packer.boxes.Insert(box)
}

func (packer *Packer) Pack() *packed_box.PackedBoxList {
	packedBoxes, _ := packer.doVolumePack()

	return packedBoxes
}

func (packer *Packer) doVolumePack() (*packed_box.PackedBoxList, error) {
	packedBoxes := packed_box.NewPackedBoxList()

	for packer.items.Count() > 0 {
		boxesToEvaluate := packer.boxes.Clone()
		packedBoxesIteration := packed_box.NewPackedBoxList()

		for !boxesToEvaluate.IsEmpty() {
			extractedBox, _ := boxesToEvaluate.Extract()
			volumePacker := NewVolumePacker(extractedBox, packer.items.Clone())
			packedBox := volumePacker.Pack()

			if packedBox.GetItems().Count() > 0 {
				packedBoxesIteration.Insert(packedBox)

				if packedBox.GetItems().Count() == packer.items.Count() {
					break
				}
			}
		}

		if packedBoxesIteration.IsEmpty() {
			return nil, errors.New("Item is too large to fit any box")
		}

		bestBox, _ := packedBoxesIteration.Top()
		bestBoxItems := bestBox.GetItems().Clone()

		unpackedItems := packer.items.Clone().Items
		for _, packedItem := range bestBoxItems.Items {
			for i, unpackedItem := range unpackedItems {
				if packedItem == unpackedItem {
					unpackedItems = append(unpackedItems[:i], unpackedItems[i+1:]...)
					break
				}
			}
		}

		unpackedItemList := box_item.NewItemList()
		for _, unpackedItem := range unpackedItems {
			unpackedItemList.Insert(unpackedItem)
		}
		packer.items = unpackedItemList
		packedBoxes.Insert(bestBox)
	}

	return packedBoxes, nil
}