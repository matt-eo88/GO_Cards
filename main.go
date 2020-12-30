package main

import (
	"fmt"
	"log"
)

func main() {

	var cards deck

	fmt.Println("======== CARDS ========")
	fmt.Println("Welcome to your virtual deck of cards!")
	fmt.Println("=======================")

	for {
		fmt.Println("Choose from the following options:")
		fmt.Println("1 - Create a new deck")
		fmt.Println("2 - See the cards")
		fmt.Println("3 - Shuffle!")
		fmt.Println("4 - Save your cards to file")
		fmt.Println("5 - Load your cards from file")
		fmt.Println("0 - Exit")

		var input string
		fmt.Scanln(&input)

		if input == "0" {
			fmt.Println("BYE !")
			break
		}

		switch input {
		case "1":
			cards = newDeck()
			fmt.Print("=== Congratulations, you have a brand new deck of cards! ===\n")

		case "2":
			if len(cards) > 0 {
				fmt.Print("=== YOUR CARDS ===\n")
				cards.print()
				fmt.Print("==================\n")
			} else {
				fmt.Println("YOUR DECK IS EMPTY! ===")
			}
		case "3":
			if len(cards) > 0 {
				cards.shuffle()
				fmt.Println("==== Cards shuffled ! ====")
			} else {
				fmt.Println("YOUR DECK IS EMPTY! ===")
			}
		case "4":
			fmt.Println("=== Please choose a filename ===")

			var fileName string
			fmt.Scanln(&fileName)

			if cards != nil {
				err := cards.saveToFile(fileName)

				if err != nil {
					log.Fatal("There was an issue while saving your file")
				}
			}

			fmt.Println("=== File Saved !!! ===")

		case "5":
			fmt.Println("=== What's the filename? ===")

			var fileName string
			fmt.Scanln(&fileName)

			cards = newDeckFromFile(fileName)
			fmt.Println("=== Success ! ===")

		default:
			fmt.Print("Error: Unknown command :(\n")
		}
	}
}
