package deck

import (
	"go-blackjack/card"
	"time"

	"math/rand"
)

// Deck
type Deck struct {
	Cards []*card.Card
}

func NewDeck() *Deck {
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	var cards []*card.Card

	for i := 0; i < 6; i++ {
		for _, suit := range suits {
			for _, rank := range ranks {
				value := getCardValue(rank)
				cards = append(cards, card.NewCard(rank, suit, value))
			}
		}
	}

	deck := &Deck{
		Cards: cards,
	}

	deck.shuffle()
	return deck
}

func getCardValue(rank string) int {
	switch rank {
	case "J", "Q", "K":
		return 10
	case "A":
		return 11 // Ace can be 11 or 1, but we'll adjust later
	default:
		return int(rank[0] - '0')
	}
}

func (d *Deck) shuffle() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(d.Cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

func (d *Deck) DealCard() *card.Card {
	if len(d.Cards) == 0 {
		return nil
	}
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}
