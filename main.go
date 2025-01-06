package main

import "go-blackjack/game"

func main() {
	game := game.NewGame("Player 1")
	game.Start()
	game.Dealer.DealerDraws(game.Deck)
	game.DetermineWinner()
}
