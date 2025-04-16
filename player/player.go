package player

import (
	"Project1/height"
	"fmt"
)

// Struct created for PlayerCreds
type PlayerCreds struct {
	Name   string
	Age    int
	Height height.PlayerHeight
}

// Function that creates the Player object
func Player(name string, age int, height height.PlayerHeight) PlayerCreds {
	p := PlayerCreds{
		Name:   name,
		Age:    age,
		Height: height,
	}
	return p
}

// Reciever function to Get Player name
func (p PlayerCreds) GetName() string {
	fs := fmt.Sprintf(`Name: %v`, p.Name)
	return fs
}

// Reciever function to Get Player age
func (p PlayerCreds) GetAge() float64 {

	return float64(p.Age)
}

// Reciever function to Get Player height
func (p PlayerCreds) GetPlayerHeight() string {
	fs := fmt.Sprintf(`%0.f' %0.f"`, p.Height.Feet, p.Height.Inches)
	return fs
}

// Reciver function to convert the data into a formatted string
func (p PlayerCreds) ToString() string {
	getHeight := p.GetPlayerHeight()
	fs := fmt.Sprintf(`Name: %s, Age: %d, %s`, p.Name, p.Age, getHeight)
	return fs
}
