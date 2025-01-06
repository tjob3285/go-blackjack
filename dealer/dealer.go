package dealer

import (
	"fmt"
	"go-blackjack/card"
	"go-blackjack/deck"
	"go-blackjack/player"
)

type Dealer struct {
	player.Player
}

func NewDealer() *Dealer {
	return &Dealer{
		Player: player.Player{
			Name:     "Dealer",
			Hand:     []*card.Card{},
			IsDealer: true,
		},
	}
}

func (d *Dealer) DealerDraws(deck *deck.Deck) {
	for d.Score < 17 {
		card := deck.DealCard()
		if card != nil {
			d.AddCard(card)
			fmt.Printf("%s drew a %s of %s. Score: %d\n\n", d.Name, card.Rank, card.Suit, d.Score)
		}
	}
}
