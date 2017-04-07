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

type ErrorMessage struct {
	Error string
}

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		errorJson, _ := json.Marshal(&ErrorMessage{
			Error: "Not enough arguments passed to script",
		})
		fmt.Printf(string(errorJson))

		os.Exit(0)
	}

	boxJson := args[0]
	itemJson := args[1]

	boxes := make([]*box.Box, 0)
	json.Unmarshal([]byte(boxJson), &boxes)

	for _, packerBox := range boxes {
		packerBox.RecalculateVolume()
	}

	items := make([]*box_item.Item, 0)
	json.Unmarshal([]byte(itemJson), &items)

	packer := Packer.NewPacker()
	packer.SetBoxes(boxes)
	packer.SetItems(items)

	packedBoxes, err := runner.Pack(packer)

	if err != nil {
		errorJson, _ := json.Marshal(&ErrorMessage{
			Error: err.Error(),
		})
		fmt.Printf(string(errorJson))

		os.Exit(0)
	}

	packedBoxesJson, _ := json.Marshal(packedBoxes.ToJson())
	fmt.Printf(string(packedBoxesJson))

	os.Exit(1)
}
