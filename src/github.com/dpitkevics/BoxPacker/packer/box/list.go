package box

import "errors"

type BoxList struct {
	boxes []*Box
}

func NewBoxList() *BoxList {
	return &BoxList{}
}

func (boxList *BoxList) Insert(box *Box) {
	boxList.boxes = append(boxList.boxes, box)
}

func (boxList *BoxList) Count() int {
	return len(boxList.boxes)
}

func (boxList *BoxList) IsEmpty() bool {
	return boxList.Count() == 0
}

func (boxList *BoxList) Extract() (*Box, error) {
	if boxList.Count() == 0 {
		return nil, errors.New("No boxes left in list")
	}

	box := boxList.boxes[0]

	boxList.boxes = append(boxList.boxes[:0], boxList.boxes[1:]...)

	return box, nil
}

func (boxList *BoxList) Clone() *BoxList {
	newList := NewBoxList()
	newList.boxes = boxList.boxes

	return newList
}

