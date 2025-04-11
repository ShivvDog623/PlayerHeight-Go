package height

import (
	"fmt"
	"math"
)

type PlayerHeight struct {
	Feet   float64
	Inches float64
}

func Height(feet float64, inches float64) PlayerHeight {

	if inches > 12 {
		feet = math.Floor((inches / 12) + feet)
		inches = math.Floor(math.Mod(inches, 12.0))

	}

	h := PlayerHeight{
		Feet:   feet,
		Inches: inches,
	}
	return h
}

func (h PlayerHeight) ToString() string {
	fs := fmt.Sprintf(`%0.f' %0.f"`, h.Feet, h.Inches)
	return fs
}
