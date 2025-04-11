package main

import (
	"Project1/height"
	"fmt"
)

func main() {
	player := height.Height(5, 13)
	fmt.Println(player.GetInches())

}
