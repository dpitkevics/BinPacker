package runner

import (
	"github.com/dpitkevics/BoxPacker/packer/packed_box"
	"github.com/dpitkevics/BoxPacker/packer/redistributor"
	"github.com/dpitkevics/BoxPacker/packer"
)

func Pack(packer *Packer.Packer) *packed_box.PackedBoxList {
	packedBoxes, _ := packer.DoVolumePacking()

	if packedBoxes.Count() > 1 && packedBoxes.Count() < 20 {
		redistributor := redistributor.NewWeightRedistributor(packer.GetBoxes())
		packedBoxes = redistributor.RedistributeWeight(packedBoxes)
	}

	return packedBoxes
}
