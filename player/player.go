package player

import "go-blackjack/card"

type Player struct {
	Name     string
	Hand     []*card.Card
	Score    int
	IsDealer bool
}

func NewPlayer(name string, isDealer bool) *Player {
	return &Player{
		Name:     name,
		Hand:     []*card.Card{},
		IsDealer: isDealer,
	}
}

func (p *Player) AddCard(card *card.Card) {
	p.Hand = append(p.Hand, card)
	p.updateScore()
}

func (p *Player) updateScore() {
	score := 0
	aceCount := 0

	for _, card := range p.Hand {
		score += card.Value
		if card.Rank == "A" {
			aceCount++
		}
	}

	for score > 21 && aceCount > 0 {
		score -= 10
		aceCount--
	}

	p.Score = score
}
