package box_item

import "errors"

type ItemList struct {
	Items []*Item
}

func NewItemList() *ItemList {
	return &ItemList{}
}

func (itemList *ItemList) Insert(item *Item) {
	itemList.Items = append(itemList.Items, item)
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
	newList.Items = itemList.Items

	return newList
}
