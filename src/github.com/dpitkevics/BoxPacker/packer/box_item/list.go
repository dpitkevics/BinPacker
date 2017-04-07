package box_item

import (
	"errors"
	"sort"
)

type ItemList struct {
	Items []*Item
}

func (itemList *ItemList) Len() int {
	return itemList.Count()
}

func (itemList *ItemList) Less(i, j int) bool {
	itemA := itemList.Items[i]
	itemB := itemList.Items[j]

	if itemA.Volume > itemB.Volume {
		return true
	} else if (itemA.Volume < itemB.Volume) {
		return false
	} else {
		return (itemA.Weight - itemB.Weight) > 0
	}
}

func (itemList *ItemList) Swap(i, j int) {
	itemList.Items[i], itemList.Items[j] = itemList.Items[j], itemList.Items[i]
}

func NewItemList() *ItemList {
	return &ItemList{}
}

func (itemList *ItemList) Insert(item *Item) {
	itemList.Items = append(itemList.Items, item)

	sort.Sort(itemList)
}

func (itemList *ItemList) Count() int {
	return len(itemList.Items)
}

func (itemList *ItemList) IsEmpty() bool {
	return itemList.Count() == 0
}

func (itemList *ItemList) Extract() (*Item, error) {
	if itemList.Count() == 0 {
		return nil, errors.New("No items left in list")
	}

	item := itemList.Items[0]

	itemList.Items = append(itemList.Items[:0], itemList.Items[1:]...)

	return item, nil
}

func (itemList *ItemList) Top() (*Item, error) {
	if itemList.Count() == 0 {
		return nil, errors.New("No items left in list")
	}

	item := itemList.Items[itemList.Count() - 1]

	return item, nil
}

func (itemList *ItemList) Clone() *ItemList {
	newList := NewItemList()

	for _, item := range itemList.Items {
		newList.Insert(NewItem(
			item.Id,
			item.Description,
			item.Length,
			item.Width,
			item.Height,
			item.Weight,
		))
	}

	return newList
}
