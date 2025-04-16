package height

import (
	"fmt"
	"math"
)

// Struct created for PlayerHeight
type PlayerHeight struct {
	Feet   float64
	Inches float64
}

// Function that creates the Height object
func Height(feet float64, inches float64) PlayerHeight {

	// Uesed to convert an inches above 12 into feet and leave the remaining inches
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

// Reciver function to convert the data into a formatted string
func (h PlayerHeight) ToString() string {
	fs := fmt.Sprintf(`%0.f' %0.f"`, h.Feet, h.Inches)
	return fs
}

// Reciever function to Get Feet
func (h PlayerHeight) GetFeet() float64 {
	return float64(h.Feet)
}

// Reciever function to Get Inches
func (h PlayerHeight) GetInches() float64 {
	return float64(h.Inches)
}

// Reciever that converts Feet to Inches
func (h PlayerHeight) ToInches() float64 {
	toInches := (h.Feet * 12) + h.Inches

	return toInches
}
