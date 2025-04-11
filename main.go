package main

import (
	"Project1/height"
	"Project1/player"
	"fmt"
)

func main() {
	height := height.Height(5, 13)
	player1 := player.Player("Adam", 23, height)

	fmt.Println(player1.ToString())
}
