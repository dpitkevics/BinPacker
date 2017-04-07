package runner

import (
	"github.com/dpitkevics/BoxPacker/packer/packed_box"
	"github.com/dpitkevics/BoxPacker/packer/redistributor"
	"github.com/dpitkevics/BoxPacker/packer"
)

func Pack(packer *Packer.Packer) (*packed_box.PackedBoxList, error) {
	packedBoxes, err := packer.DoVolumePacking()

	if err != nil {
		return nil, err
	}

	if packedBoxes.Count() > 1 && packedBoxes.Count() < 20 {
		weightRedistributor := redistributor.NewWeightRedistributor(packer.GetBoxes())
		packedBoxes = weightRedistributor.RedistributeWeight(packedBoxes)
	}

	return packedBoxes, nil
}
