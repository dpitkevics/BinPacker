package packed_box

import (
	"github.com/dpitkevics/BoxPacker/packer/box"
	"github.com/dpitkevics/BoxPacker/packer/box_item"
)

type PackedBoxJson struct {
	Box *box.Box
	Items []*box_item.Item
	Weight float64
	RemainingWidth float64
	RemainingLength float64
	RemainingHeight float64
	RemainingWeight float64
	UsedWidth float64
	UsedLength float64
	UsedHeight float64
}
