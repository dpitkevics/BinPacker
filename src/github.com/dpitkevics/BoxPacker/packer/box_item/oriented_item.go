package box_item

import "math"

type OrientedItem struct {
	item *Item
	length float64
	width float64
	height float64
}

func NewOrientedItem(item *Item, length float64, width float64, height float64) *OrientedItem {
	return &OrientedItem{
		item: item,
		length: length,
		width: width,
		height: height,
	}
}

func (orientedItem *OrientedItem) GetItem() *Item {
	return orientedItem.item
}

func (orientedItem *OrientedItem) GetLength() float64 {
	return orientedItem.length
}

func (orientedItem *OrientedItem) GetWidth() float64 {
	return orientedItem.width
}

func (orientedItem *OrientedItem) GetHeight() float64 {
	return orientedItem.height
}

func (orientedItem *OrientedItem) IsStable() bool {
	return orientedItem.GetHeight() <= math.Min(orientedItem.GetLength(), orientedItem.GetWidth())
}
