package main

import (
	"Project1/height"
	"Project1/player"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Declaring variables
	var name string
	var age int
	var inches float64
	var feet float64
	var totalAge int

	// Slice to hold player values from input
	var players []player.PlayerCreds

	// Initializing the scanner
	scanner := bufio.NewScanner(os.Stdin)

	// For loop (Go While Loop) used to reprompt the user for more data
	for {
		fmt.Printf("Enter player's name, age, and height in feet and inches: ")

		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)

		// Breaks out of loop if line entered is empty
		if strings.TrimSpace(line) == "" {

			fmt.Println()
			break
		}

		// Checks to see if all value fields have an input
		if len(parts) < 4 {
			fmt.Println("Please enter all fields.")
			continue
		} else if len(parts) != 4 {
			fmt.Println("That is too many fields. Only 4!")
			continue
		}

		// Parsing all inputs from the line
		name = parts[0]
		ageStr := parts[1]
		feetStr := parts[2]
		inchesStr := parts[3]

		// Converts the parsed inputs from string to correct types
		age, _ = strconv.Atoi(ageStr)
		feet, _ = strconv.ParseFloat(feetStr, 64)
		inches, _ = strconv.ParseFloat(inchesStr, 64)

		// Initializes the height
		height := height.Height(feet, inches)

		// Initializes the player using the (height)
		player := player.Player(name, age, height)

		// Appends new player to the Players slice
		players = append(players, player)

		// Calulates the total Age
		totalAge += age

	}

	// Calculating average age
	averageAge := float64(totalAge) / float64(len(players))

	// Used to find the tallest player
	var tallestPlayer player.PlayerCreds
	maxHeight := 0
	found := false

	// Loops through the slice of players and checks to see who is tallest with age less than average
	for _, v := range players {
		if v.Height.ToInches() > float64(maxHeight) && v.GetAge() <= averageAge {
			maxHeight = int(v.Height.ToInches())
			tallestPlayer = v
			found = true
		}
	}

	// If found will then show the player that is the tallest with less than average age
	if found {
		fmt.Println("\nTallest player whose age is less than the average: ")
		fmt.Println(tallestPlayer.ToString())
	}

	// Used to show that no data is entered if players len is == 0
	if len(players) == 0 {
		fmt.Println("No player data entered")
	}

}
