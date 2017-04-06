package box

type Box struct {
	Reference string
	OuterLength float64
	OuterWidth float64
	OuterHeight float64
	EmptyWeight float64
	InnerLength float64
	InnerWidth float64
	InnerHeight float64
	InnerVolume float64
	MaxWeight float64
}

func NewBox(reference string, outerLength float64, outerWidth float64, outerHeight float64, emptyWeight float64, innerLength float64, innerWidth float64, innerHeight float64, maxWeight float64) *Box {
	return &Box{
		Reference: reference,
		OuterLength: outerLength,
		OuterWidth: outerWidth,
		OuterHeight: outerHeight,
		EmptyWeight: emptyWeight,
		InnerLength: innerLength,
		InnerWidth: innerWidth,
		InnerHeight: innerHeight,
		InnerVolume: innerLength * innerWidth * innerHeight,
		MaxWeight: maxWeight,
	}
}
