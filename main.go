package main

import (
	"fmt"
	"go-blackjack/game"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
)

func main() {
	game := game.NewGame("Player 1")

	var input string
	for {
		game.Player.ResetHand()
		game.Dealer.ResetHand()

		fmt.Printf("\n💰 Starting a new round... You currently have $%d.\n", game.Player.Tokens)
		fmt.Printf("🤑 Your balance: $%d. What's your bet? ", game.Player.Tokens)

		if game.Player.Tokens == 0 {
			fmt.Printf("🚫 Insufficient funds! You have $%d left.\n", game.Player.Tokens)
			return
		}

		var betInput string
		huh.NewInput().
			Title("Enter your bet:").
			Value(&betInput).
			Run()

		betAmount, err := strconv.Atoi(strings.TrimSpace(betInput))
		if err != nil || betAmount <= 0 {
			fmt.Println("⚠️ Invalid bet. Please enter a number greater than 0.")
			continue
		}
		if betAmount > game.Player.Tokens {
			fmt.Printf("🚫 You don't have enough tokens to bet that much! You have $%d.\n", game.Player.Tokens)
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
					huh.NewOption("Split", "split"),
				).
				Value(&input).
				Run()

			switch input {
			case "hit":
				fmt.Println("🎯 You chose to 'Hit' — drawing a card...")
				game.Player.AddCard(game.Deck.DealCard())
				fmt.Printf("🃏 You drew the %s of %s!\n", game.Player.Hand[len(game.Player.Hand)-1].Rank, game.Player.Hand[len(game.Player.Hand)-1].Suit)
				fmt.Printf("🔥 Your score: %d\n\n", game.Player.Score)
				if game.Player.Score > 21 {
					game.DetermineWinner(betAmount)
					goto NextRound
				}
			case "stand":
				fmt.Println("🛑 You chose to 'Stand' — ending your turn.")
				time.Sleep(2 * time.Second)
				fmt.Printf("💥 Dealer's visible card: %s of %s — Score: %d\n", game.Dealer.Hand[1].Rank, game.Dealer.Hand[1].Suit, game.Dealer.Score)
				game.Dealer.DealerDraws(game.Deck)
				game.DetermineWinner(betAmount)
				goto NextRound
			case "split":
				if len(game.Player.Hand) == 2 && game.Player.Hand[0].Rank == game.Player.Hand[1].Rank {
					splitHands, err := game.Player.SplitHand()
					if err != nil {
						fmt.Println(err)
						continue
					}

					var additionalBetInput string
					huh.NewInput().
						Title("Enter your second bet for the split hand:").
						Value(&additionalBetInput).
						Run()

					additionalBetAmount, err := strconv.Atoi(strings.TrimSpace(additionalBetInput))
					if err != nil || additionalBetAmount <= 0 || additionalBetAmount > game.Player.Tokens {
						fmt.Println("⚠️ Invalid bet for split hand. Please enter a valid number.")
						continue
					}

					fmt.Println("You have split your hand! Now you have two separate hands to play.")

					for _, splitHand := range splitHands {
						fmt.Printf("\nPlaying for Hand 1: %v\n", splitHand.Hand)
						for splitHand.Score < 21 || input == "double" {
							huh.NewSelect[string]().
								Title("What would you like to do? (Hand 1)").
								Options(
									huh.NewOption("Hit", "hit"),
									huh.NewOption("Double", "double"),
									huh.NewOption("Stand", "stand"),
								).
								Value(&input).
								Run()

							switch input {
							case "hit":
								fmt.Println("🎯 You chose to 'Hit' — drawing a card...")
								splitHand.AddCard(game.Deck.DealCard())
								fmt.Printf("🃏 You drew the %s of %s!\n", splitHand.Hand[len(splitHand.Hand)-1].Rank, splitHand.Hand[len(splitHand.Hand)-1].Suit)
								fmt.Printf("🔥 Your score: %d\n\n", splitHand.Score)
								if splitHand.Score > 21 {
									fmt.Println("❌ Hand 1 went over 21!")
									goto NextRound
								}
							case "stand":
								fmt.Println("🛑 You chose to 'Stand' — ending your turn for Hand 1.")
								continue
							case "double":
								if additionalBetAmount*2 > splitHand.Tokens {
									fmt.Println("🚫 You don't have enough tokens to double down on Hand 1.")
									continue
								}
								fmt.Println("💪 You chose to 'Double Down' — drawing one more card!")
								splitHand.AddCard(game.Deck.DealCard())
								continue
							}
						}

						fmt.Printf("\nPlaying for Hand 2: %v\n", splitHand.Hand)
					}

				} else {
					fmt.Println("⚠️ You can only split if you have two cards of the same rank!")
					continue
				}
			case "double":
				if betAmount*2 > game.Player.Tokens {
					fmt.Println("🚫 You don't have enough tokens to double down.")
					continue
				}
				fmt.Println("💪 You chose to 'Double Down' — drawing one more card!")
				game.Player.AddCard(game.Deck.DealCard())
				fmt.Printf("🃏 You drew the %s of %s!\n", game.Player.Hand[len(game.Player.Hand)-1].Rank, game.Player.Hand[len(game.Player.Hand)-1].Suit)
				fmt.Printf("🔥 Your score: %d\n\n", game.Player.Score)
				time.Sleep(2 * time.Second)
				game.Dealer.DealerDraws(game.Deck)
				game.DetermineWinner(betAmount * 2)
				goto NextRound
			default:
				fmt.Println("⚠️ Invalid option, please type 'hit', 'stand', or 'double'.")
			}
		}

		if game.Player.Score >= 21 {
			fmt.Println("🎉 You've reached 21 or higher! Let's see the results...")
			game.DetermineWinner(betAmount)
			goto NextRound
		}

	NextRound:
		var playAgain bool
		if game.Player.Tokens == 0 {
			playAgain = false
			fmt.Println("No Money left.")
		} else {
			huh.NewConfirm().
				Title("Do you want to play another round?").
				Affirmative("Yes!").
				Negative("No.").
				Value(&playAgain).
				Run()
		}

		if !playAgain {
			fmt.Println("🙏 Thanks for playing! Exiting the game.")
			break
		}
	}
}
