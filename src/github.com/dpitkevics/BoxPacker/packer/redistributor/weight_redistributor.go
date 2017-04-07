package redistributor

import (
	"github.com/dpitkevics/BoxPacker/packer/box"
	"github.com/dpitkevics/BoxPacker/packer/packed_box"
	"github.com/dpitkevics/BoxPacker/packer"

	"github.com/bradfitz/slice"
)

type WeightRedistributor struct {
	boxes *box.BoxList
}

func NewWeightRedistributor(boxes *box.BoxList) *WeightRedistributor {
	return &WeightRedistributor{
		boxes: boxes,
	}
}

func (redistributor *WeightRedistributor) RedistributeWeight(originalBoxes *packed_box.PackedBoxList) *packed_box.PackedBoxList {
	targetWeight := originalBoxes.GetMeanWeight()

	packedBoxes := packed_box.NewPackedBoxList()
	overWeightBoxes := []*packed_box.PackedBox{}
	underWeightBoxes := []*packed_box.PackedBox{}
	clonedOriginalBoxes := originalBoxes.Clone()

	for _, packedBox := range clonedOriginalBoxes.PackedBoxes {
		boxWeight := packedBox.GetWeight()
		if boxWeight > targetWeight {
			overWeightBoxes = append(overWeightBoxes, packedBox)
		} else if boxWeight < targetWeight {
			underWeightBoxes = append(underWeightBoxes, packedBox)
		} else {
			packedBoxes.Insert(packedBox)
		}
	}

	tryRepack := true
	for tryRepack {
		tryRepack = false

		OUTER:
		for u, underWeightBox := range underWeightBoxes {
			for o, overWeightBox := range overWeightBoxes {
				overWeightBoxItems := overWeightBox.GetItems().Items

				for oi, overWeightBoxItem := range overWeightBoxItems {
					if underWeightBox.GetWeight() + overWeightBoxItem.Weight > targetWeight {
						continue
					}

					newItemsForLighterBox := underWeightBox.GetItems().Clone()
					newItemsForLighterBox.Insert(overWeightBoxItem)

					newLighterBoxPacker := Packer.NewPacker()
					newLighterBoxPacker.SetBoxes(redistributor.boxes.Boxes)
					newLighterBoxPacker.SetItems(newItemsForLighterBox.Items)

					newLighterBoxes,_ := newLighterBoxPacker.DoVolumePacking()
					newLighterBox,_ := newLighterBoxes.Extract()

					if newLighterBox.GetItems().Count() == newItemsForLighterBox.Count() {
						overWeightBoxItems = append(overWeightBoxItems[:oi], overWeightBoxItems[oi+1:]...)

						newHeavierBoxPacker := Packer.NewPacker()
						newHeavierBoxPacker.SetBoxes(redistributor.boxes.Boxes)
						newHeavierBoxPacker.SetItems(overWeightBoxItems)

						newHeavierBoxes,_ := newHeavierBoxPacker.DoVolumePacking()
						if newHeavierBoxes.Count() > 1 {
							return originalBoxes
						}

						overWeightBoxes[o],_ = newHeavierBoxes.Extract()
						underWeightBoxes[u] = newLighterBox

						tryRepack = true
						slice.Sort(overWeightBoxes, func(i, j int) bool {
							boxA := overWeightBoxes[i]
							boxB := overWeightBoxes[j]

							choice := boxB.GetItems().Count() - boxA.GetItems().Count()

							if choice == 0 {
								choice = int(boxA.GetBox().InnerVolume - boxB.GetBox().InnerVolume)
							}

							if choice == 0 {
								choice = int(boxB.GetWeight() - boxA.GetWeight())
							}

							return choice > 0
						})
						slice.Sort(underWeightBoxes, func(i, j int) bool {
							boxA := overWeightBoxes[i]
							boxB := overWeightBoxes[j]

							choice := boxB.GetItems().Count() - boxA.GetItems().Count()

							if choice == 0 {
								choice = int(boxA.GetBox().InnerVolume - boxB.GetBox().InnerVolume)
							}

							if choice == 0 {
								choice = int(boxB.GetWeight() - boxA.GetWeight())
							}

							return choice > 0
						})
						break OUTER
					}
				}
			}
		}
	}

	for _, overWeightBox := range overWeightBoxes {
		packedBoxes.Insert(overWeightBox)
	}

	for _, underWeightBox := range underWeightBoxes {
		packedBoxes.Insert(underWeightBox)
	}

	return packedBoxes
}
