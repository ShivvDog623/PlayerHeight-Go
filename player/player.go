package player

import (
	"Project1/height"
	"fmt"
)

type PlayerCreds struct {
	Name   string
	Age    int
	Height height.PlayerHeight
}

func Player(name string, age int, height height.PlayerHeight) PlayerCreds {
	p := PlayerCreds{
		Name:   name,
		Age:    age,
		Height: height,
	}
	return p
}

// reciever function to Get Player name
func (p PlayerCreds) GetName() string {
	fs := fmt.Sprintf(`Name: %v`, p.Name)
	return fs
}

// reciever function to Get Player age
func (p PlayerCreds) GetAge() string {
	fs := fmt.Sprintf(`Age: %v`, p.Age)
	return fs
}

// reciever function to Get Player height
func (p PlayerCreds) GetPlayerHeight() string {
	fs := fmt.Sprintf(`%0.f' %0.f"`, p.Height.Feet, p.Height.Inches)
	return fs
}

func (p PlayerCreds) ToString() string {
	getHeight := p.GetPlayerHeight()
	fs := fmt.Sprintf(`Name: %s, Age: %d, %s`, p.Name, p.Age, getHeight)
	return fs
}
