package packed_box

import "errors"

type PackedBoxList struct {
	PackedBoxes []*PackedBox
}

func NewPackedBoxList() *PackedBoxList {
	return &PackedBoxList{}
}

func (packedBoxList *PackedBoxList) Insert(packedBox *PackedBox) {
	packedBoxList.PackedBoxes = append(packedBoxList.PackedBoxes, packedBox)
}

func (packedBoxList *PackedBoxList) Count() int {
	return len(packedBoxList.PackedBoxes)
}

func (packedBoxList *PackedBoxList) IsEmpty() bool {
	return packedBoxList.Count() == 0
}

func (packedBoxList *PackedBoxList) Top() (*PackedBox, error) {
	if packedBoxList.Count() == 0 {
		return nil, errors.New("No items left in list")
	}

	item := packedBoxList.PackedBoxes[packedBoxList.Count() - 1]

	return item, nil
}

func (packedBoxList *PackedBoxList) Clone() *PackedBoxList {
	newList := NewPackedBoxList()
	newList.PackedBoxes = packedBoxList.PackedBoxes

	return newList
}
