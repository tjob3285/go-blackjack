package game

import (
	"fmt"
	"go-blackjack/dealer"
	"go-blackjack/deck"
	"go-blackjack/player"
)

type Game struct {
	Deck   *deck.Deck
	Player *player.Player
	Dealer *dealer.Dealer
}

func NewGame(playerName string) *Game {
	deck := deck.NewDeck()
	player := player.NewPlayer(playerName, false)
	dealer := dealer.NewDealer()
	return &Game{
		Deck:   deck,
		Player: player,
		Dealer: dealer,
	}
}

func (g *Game) Start() {
	g.Player.AddCard(g.Deck.DealCard())
	g.Dealer.AddCard(g.Deck.DealCard())
	g.Player.AddCard(g.Deck.DealCard())
	g.Dealer.AddCard(g.Deck.DealCard())

	for _, card := range g.Player.Hand {
		fmt.Printf("Player's hand: %d\n", card.Value)
	}

	fmt.Printf("Dealer's face-up card: %d\n", g.Dealer.Hand[0].Value)
}

func (g *Game) DetermineWinner() {
	if g.Player.Score > 21 {
		fmt.Println("Player busts!")
	} else if g.Dealer.Score > 21 {
		fmt.Println("Dealer busts! Player wins.")
	} else if g.Player.Score > g.Dealer.Score {
		fmt.Println("Player wins!")
	} else if g.Player.Score < g.Dealer.Score {
		fmt.Println("Dealer wins!")
	} else {
		fmt.Println("It's a tie!")
	}
}
