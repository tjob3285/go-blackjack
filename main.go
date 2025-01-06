package main

import (
	"bufio"
	"fmt"
	"go-blackjack/game"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	game := game.NewGame("Player 1")

	var input string
	for {
		game.Player.ResetHand()
		game.Dealer.ResetHand()

		fmt.Printf("\nStarting a new round. You have $%d.\n", game.Player.Tokens)
		fmt.Printf("You have $%d. Enter your bet:", game.Player.Tokens)
		scanner.Scan()
		betInput := strings.TrimSpace(scanner.Text())
		betAmount, err := strconv.Atoi(betInput)
		if err != nil || betAmount <= 0 {
			fmt.Println("Invalid bet amount. Please enter a valid number greater than 0.")
			return
		}
		if betAmount > game.Player.Tokens {
			fmt.Printf("Insufficient funds! You only have $%d.\n", game.Player.Tokens)
			return
		}

		game.Start()

		for game.Player.Score < 21 {

			fmt.Println("What would you like to do?")
			fmt.Println("Type 'hit' to draw a card, 'stand' to hold, or 'quit' to exit the game.")

			scanner.Scan()
			input = strings.ToLower(scanner.Text())

			switch input {
			case "hit":
				fmt.Println("You chose to 'hit' — drawing a card...")
				game.Player.AddCard(game.Deck.DealCard())
				fmt.Printf("Player drew %s of %s \n\n", game.Player.Hand[len(game.Player.Hand)-1].Rank, game.Player.Hand[len(game.Player.Hand)-1].Suit)
				fmt.Printf("Player score %d \n\n", game.Player.Score)
				if game.Player.Score > 21 {
					game.DetermineWinner(betAmount)
					goto NextRound
				}
			case "stand":
				fmt.Println("You chose to 'stand' — ending your turn.")
				fmt.Printf("Dealer shows %s of %s: %d: \n", game.Dealer.Hand[1].Rank, game.Dealer.Hand[1].Suit, game.Dealer.Score)
				game.Dealer.DealerDraws(game.Deck)
				game.DetermineWinner(betAmount)
				goto NextRound
			case "quit":
				fmt.Println("Exiting the game...")
				return
			default:
				fmt.Println("Invalid option, please type 'hit', 'stand', or 'quit'.")
			}
		}

		if game.Player.Score >= 21 {
			fmt.Println("You reached 21 or higher!")
			game.DetermineWinner(betAmount)
			goto NextRound
		}

	NextRound:
		fmt.Println("Do you want to play another round? (yes/no)")
		scanner.Scan()
		playAgain := strings.ToLower(scanner.Text())

		if playAgain != "yes" {
			fmt.Println("Thank you for playing! Exiting the game.")
			break
		}
	}
}
