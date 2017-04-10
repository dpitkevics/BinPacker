package packed_box

import (
	"errors"
	"sort"
)

type PackedBoxList struct {
	PackedBoxes []*PackedBox
	meanWeight float64
}

func (packedBoxList *PackedBoxList) Len() int {
	return packedBoxList.Count()
}

func (packedBoxList *PackedBoxList) Less(i, j int) bool {
	boxA := packedBoxList.PackedBoxes[i]
	boxB := packedBoxList.PackedBoxes[j]

	choice := boxA.GetItems().Count() - boxB.items.Count()

	if choice == 0 {
		choice = int(boxB.GetBox().InnerVolume - boxA.GetBox().InnerVolume)
	}

	if choice == 0 {
		choice = int(boxA.GetWeight() - boxB.GetWeight())
	}

	return choice > 0
}

func (packedBoxList *PackedBoxList) Swap(i, j int) {
	packedBoxList.PackedBoxes[i], packedBoxList.PackedBoxes[j] = packedBoxList.PackedBoxes[j], packedBoxList.PackedBoxes[i]
}

func NewPackedBoxList() *PackedBoxList {
	return &PackedBoxList{}
}

func (packedBoxList *PackedBoxList) Insert(packedBox *PackedBox) {
	packedBoxList.PackedBoxes = append(packedBoxList.PackedBoxes, packedBox)

	sort.Sort(packedBoxList)
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

	item := packedBoxList.PackedBoxes[0]

	return item, nil
}

func (packedBoxList *PackedBoxList) Extract() (*PackedBox, error) {
	box, err := packedBoxList.Top()

	if err != nil {
		return nil, err
	}

	packedBoxList.PackedBoxes = append(packedBoxList.PackedBoxes[:0], packedBoxList.PackedBoxes[1:]...)

	return box, nil
}

func (packedBoxList *PackedBoxList) Clone() *PackedBoxList {
	newList := NewPackedBoxList()

	for _, packedBox := range packedBoxList.PackedBoxes {
		newList.Insert(NewPackedBox(
			packedBox.box,
			packedBox.items,
			packedBox.remainingWidth,
			packedBox.remainingLength,
			packedBox.remainingHeight,
			packedBox.remainingWeight,
			packedBox.usedWidth,
			packedBox.useLength,
			packedBox.usedHeight,
		))
	}

	return newList
}

func (packedBoxList *PackedBoxList) GetBestBox() (*PackedBox, error) {
	sort.Sort(packedBoxList)

	return packedBoxList.Top()
}

func (packedBoxList *PackedBoxList) GetMeanWeight() float64 {
	if packedBoxList.meanWeight != 0 {
		return packedBoxList.meanWeight
	}

	clonedBoxList := packedBoxList.Clone()
	for _, box := range clonedBoxList.PackedBoxes {
		packedBoxList.meanWeight += box.GetWeight()
	}

	packedBoxList.meanWeight = packedBoxList.meanWeight / float64(packedBoxList.Count())
	return packedBoxList.meanWeight
}

func (packedBoxList *PackedBoxList) ToJson() []*PackedBoxJson {
	var packedBoxJsonList []*PackedBoxJson

	for _, packedBox := range packedBoxList.PackedBoxes {
		packedBoxJson := &PackedBoxJson{
			Box: packedBox.box,
			Items: packedBox.items.Items,
			Weight: packedBox.weight,
			RemainingWidth: packedBox.remainingWidth,
			RemainingLength: packedBox.remainingLength,
			RemainingHeight: packedBox.remainingHeight,
			RemainingWeight: packedBox.remainingWeight,
			UsedWidth: packedBox.usedWidth,
			UsedLength: packedBox.useLength,
			UsedHeight: packedBox.usedHeight,
		}

		packedBoxJsonList = append(packedBoxJsonList, packedBoxJson)
	}

	return packedBoxJsonList
}
