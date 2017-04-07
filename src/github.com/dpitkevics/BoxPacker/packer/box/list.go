package box

import (
	"errors"
	"sort"
)

type BoxList struct {
	Boxes []*Box
}

func (boxList *BoxList) Len() int {
	return boxList.Count()
}

func (boxList *BoxList) Less(i, j int) bool {
	boxA := boxList.Boxes[i]
	boxB := boxList.Boxes[j]

	return boxB.InnerVolume > boxA.InnerVolume
}

func (boxList *BoxList) Swap(i, j int) {
	boxList.Boxes[i], boxList.Boxes[j] = boxList.Boxes[j], boxList.Boxes[i]
}

func NewBoxList() *BoxList {
	return &BoxList{}
}

func (boxList *BoxList) Insert(box *Box) {
	boxList.Boxes = append(boxList.Boxes, box)

	sort.Sort(boxList)
}

func (boxList *BoxList) Count() int {
	return len(boxList.Boxes)
}

func (boxList *BoxList) IsEmpty() bool {
	return boxList.Count() == 0
}

func (boxList *BoxList) Extract() (*Box, error) {
	if boxList.Count() == 0 {
		return nil, errors.New("No Boxes left in list")
	}

	box := boxList.Boxes[0]

	boxList.Boxes = append(boxList.Boxes[:0], boxList.Boxes[1:]...)

	return box, nil
}

func (boxList *BoxList) Clone() *BoxList {
	newList := NewBoxList()

	for _, box := range boxList.Boxes {
		newList.Insert(NewBox(
			box.Reference,
			box.OuterLength,
			box.OuterWidth,
			box.OuterHeight,
			box.EmptyWeight,
			box.InnerLength,
			box.InnerWidth,
			box.InnerHeight,
			box.MaxWeight,
		))
	}

	return newList
}

