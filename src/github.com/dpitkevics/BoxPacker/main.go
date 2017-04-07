package main

import (
	"fmt"
	"github.com/dpitkevics/BoxPacker/packer"
	"github.com/dpitkevics/BoxPacker/packer/box"
	"github.com/dpitkevics/BoxPacker/packer/box_item"
	"github.com/dpitkevics/BoxPacker/packer/runner"
	"os"
	"encoding/json"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf("Error: not enough arguments passed")
		os.Exit(0)
	}

	boxJson := args[0]
	itemJson := args[1]

	boxes := make([]*box.Box, 0)
	json.Unmarshal([]byte(boxJson), &boxes)

	items := make([]*box_item.Item, 0)
	json.Unmarshal([]byte(itemJson), &items)

	packer := Packer.NewPacker()
	packer.SetBoxes(boxes)
	packer.SetItems(items)

	packedBoxes := runner.Pack(packer)

	packedBoxesJson, _ := json.Marshal(packedBoxes.ToJson())
	fmt.Printf(string(packedBoxesJson))

	os.Exit(1)
}
