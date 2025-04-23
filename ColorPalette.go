package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Function to get a random color
func randomColor(colors []string) string {
	rand.Seed(time.Now().UnixNano()) // Initialize random seed
	randomIndex := rand.Intn(len(colors)) // Generate random index
	return colors[randomIndex] // Return a random color
}

// Placeholder function for saving colors
func saveColor(savedColors *[]string, color string) {
	*savedColors = append(*savedColors, color)
	fmt.Println("Color saved:", color)
}

// Function to display saved colors
func displaySavedColors(savedColors []string) {
	if len(savedColors) == 0 {
		fmt.Println("No colors saved yet!")
	} else {
		fmt.Println("Saved colors:")
		for _, color := range savedColors {
			fmt.Println(color)
		}
	}
}

func main() {
	// Define color groups
	colorGroups := map[string][]string{
		"red": {
			"Crimson: #DC143C", "Scarlet: #FF2400", "Ruby: #9B111E",
			"Rose Red: #C21E56", "Firebrick: #B22222",
		},
		"orange": {
			"Amber: #FFBF00", "Tangerine: #F28500", "Pumpkin: #FF7518",
			"Carrot Orange: #ED9121", "Mango: #FFC324",
		},
		"yellow": {
			"Gold: #FFD700", "Lemon: #FFF44F", "Dandelion: #F0E130",
			"Buttery Yellow: #FFFD96", "Sunflower Yellow: #FFDA03",
		},
		"green": {
			"Forest Green: #228B22", "Emerald Green: #50C878", "Mint Green: #98FF98",
			"Olive Green: #808000", "Spring Green: #00FF7F",
		},
		"blue": {
			"Steel Blue: #4682B4", "Sky Blue: #87CEEB", "Royal Blue: #4169E1",
			"Navy Blue: #000080", "Dodger Blue: #1E90FF",
		},
		"indigo": {
			"Indigo: #4B0082", "Persian Indigo: #32127A", "Slate Blue: #6A5ACD",
			"Blue-Violet: #8A2BE2", "Twilight Blue: #0C0D3C",
		},
		"violet": {
			"Lavender: #E6E6FA", "Amethyst: #9966CC", "Plum: #DDA0DD",
			"Fuchsia: #FF00FF", "Grape: #6F2DA8",
		},
	}

	// Slice for saved colors
	var savedColors []string

	for {
		var command string
		fmt.Println("\nWelcome to the Color Palette!")
		fmt.Println("Commands:")
		fmt.Println("getcolor - Get a random color shade")
		fmt.Println("specify - Specify a color group to get a random shade")
		fmt.Println("add - Save a color of your choice")
		fmt.Println("saved - Display saved colors")
		fmt.Println("exit - Exit the program")
		fmt.Print("Enter a command: ")
		fmt.Scan(&command)

		switch command {
		case "getcolor":
			// Combine all colors into a single slice
			var allColors []string
			for _, group := range colorGroups {
				allColors = append(allColors, group...)
			}
			random := randomColor(allColors)
			fmt.Println("Random color:", random)

		case "specify":
			var groupName string
			fmt.Print("Enter the color group (red, orange, yellow, green, blue, indigo, violet): ")
			fmt.Scan(&groupName)
			if colors, exists := colorGroups[groupName]; exists {
				fmt.Println("Random color from", groupName, "group:", randomColor(colors))
			} else {
				fmt.Println("Invalid color group. Try again!")
			}

		case "add":
			var customColor string
			fmt.Print("Enter the color and hex code to save (e.g., 'Cool Blue: #0011FF'): ")
			fmt.Scan(&customColor)
			saveColor(&savedColors, customColor)

		case "saved":
			displaySavedColors(savedColors)

		case "exit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid command. Please try again.")
		}
	}
}