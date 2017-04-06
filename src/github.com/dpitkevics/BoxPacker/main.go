package main

import (
	"fmt"
	"github.com/dpitkevics/BoxPacker/packer"
	"github.com/dpitkevics/BoxPacker/packer/box"
	"github.com/dpitkevics/BoxPacker/packer/box_item"
)

func main() {
	packer := Packer.NewPacker()

	packer.AddBox(box.NewBox("Box 1", 32.0, 24.0, 24.0, 10, 32.0, 24.0, 24.0, 1000))
	packer.AddBox(box.NewBox("Box 2", 32.0, 24.0, 24.0, 10, 32.0, 24.0, 24.0, 1000))

	for i := 0; i < 1000; i++ {
		packer.AddItem(box_item.NewItem(fmt.Sprintf("Item %d", i), 11.0, 7.5, 0.5, 0.2))
	}

	packedBoxes := packer.Pack()
	fmt.Printf("%+v\n", packedBoxes)
}
