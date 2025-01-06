package main

import (
	"fmt"
	"go-blackjack/game"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
)

func main() {
	game := game.NewGame("Player 1")

	var input string
	for {
		game.Player.ResetHand()
		game.Dealer.ResetHand()

		fmt.Printf("\nStarting a new round. You have $%d.\n", game.Player.Tokens)
		fmt.Printf("You have $%d. Enter your bet:", game.Player.Tokens)

		if game.Player.Tokens == 0 {
			fmt.Printf("Insufficient funds! You only have $%d.\n", game.Player.Tokens)
			return
		}

		var betInput string
		huh.NewInput().
			Title("Enter your bet:").
			Value(&betInput).
			Run()

		betAmount, err := strconv.Atoi(strings.TrimSpace(betInput))
		if err != nil || betAmount <= 0 {
			fmt.Println("Invalid bet amount. Please enter a valid number greater than 0.")
			continue
		}
		if betAmount > game.Player.Tokens {
			fmt.Printf("Insufficient funds! You only have $%d.\n", game.Player.Tokens)
			continue
		}

		game.Start()

		for game.Player.Score < 21 || input == "double" {

			huh.NewSelect[string]().
				Title("What would you like to do?").
				Options(
					huh.NewOption("Hit", "hit"),
					huh.NewOption("Double", "double"),
					huh.NewOption("Stand", "stand"),
				).
				Value(&input).
				Run()

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
			case "double":
				if betAmount*2 > game.Player.Tokens {
					fmt.Println("Cannot double down, not enough tokens")
					continue
				}
				fmt.Println("You chose to 'double down' — ending your turn.")
				game.Player.AddCard(game.Deck.DealCard())
				fmt.Printf("Player drew %s of %s \n\n", game.Player.Hand[len(game.Player.Hand)-1].Rank, game.Player.Hand[len(game.Player.Hand)-1].Suit)
				fmt.Printf("Player score %d \n\n", game.Player.Score)
				game.DetermineWinner(betAmount * 2)
				goto NextRound
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
		var playAgain bool
		if game.Player.Tokens == 0 {
			playAgain = false
		} else {
			huh.NewConfirm().
				Title("Do you want to play another round?").
				Affirmative("Yes!").
				Negative("No.").
				Value(&playAgain).
				Run()
		}

		if !playAgain {
			fmt.Println("Thank you for playing! Exiting the game.")
			break
		}
	}
}
