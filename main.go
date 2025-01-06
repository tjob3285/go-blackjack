package main

import (
	"bufio"
	"fmt"
	"go-blackjack/game"
	"os"
	"strings"
)

func main() {
	game := game.NewGame("Player 1")
	game.Start()
	//game.Dealer.DealerDraws(game.Deck)
	fmt.Printf("Player score %d", game.Player.Score)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("What would you like to do?")
		fmt.Println("Type 'hit' to draw a card, 'stand' to hold, or 'quit' to exit the game.")

		scanner.Scan()
		input := strings.ToLower(scanner.Text())

		switch input {
		case "hit":
			fmt.Println("You chose to 'hit' — drawing a card...")
			game.Player.AddCard(game.Deck.DealCard())
			fmt.Printf("Player score %d", game.Player.Score)
		case "stand":
			fmt.Println("You chose to 'stand' — ending your turn.")
			game.DetermineWinner()
		case "quit":
			fmt.Println("Exiting the game...")
			return
		default:
			fmt.Println("Invalid option, please type 'hit', 'stand', or 'quit'.")
		}
	}
}
