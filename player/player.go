package player

import (
	"go-blackjack/card"
	"log"
	"os"
	"strconv"
)

type Player struct {
	Name     string
	Hand     []*card.Card
	Score    int
	IsDealer bool
	Tokens   int
}

func NewPlayer(name string, isDealer bool, tokens int) *Player {
	return &Player{
		Name:     name,
		Hand:     []*card.Card{},
		IsDealer: isDealer,
		Tokens:   tokens,
	}
}

func (p *Player) AddCard(card *card.Card) {
	p.Hand = append(p.Hand, card)
	p.updateScore()
}

func (p *Player) ResetHand() {
	p.Hand = []*card.Card{}
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

func (p *Player) UpdateTokens(betAmount int, outcome bool) error {
	tokensTxt := "tokens.txt"

	if outcome {
		p.Tokens += betAmount
	} else {
		p.Tokens -= betAmount
	}

	// Convert new token count to string
	newTokenStr := strconv.Itoa(p.Tokens)

	// Write the new value to the file
	err := os.WriteFile(tokensTxt, []byte(newTokenStr), 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
		return err
	}

	return nil
}
