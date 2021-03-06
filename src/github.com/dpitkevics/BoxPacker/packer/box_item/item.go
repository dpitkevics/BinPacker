package box_item

type Item struct {
	Identifier  string
	Description string
	Length      float64
	Width       float64
	Height      float64
	Weight      float64
	Volume      float64
}

func NewItem(identifier string, description string, length float64, width float64, height float64, weight float64) *Item {
	return &Item{
		Identifier:  identifier,
		Description: description,
		Length:      length,
		Width:       width,
		Height:      height,
		Weight:      weight,
		Volume:      length * width * height,
	}
}
