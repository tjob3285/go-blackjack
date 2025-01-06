package game

import (
	"fmt"
	"go-blackjack/dealer"
	"go-blackjack/deck"
	"go-blackjack/player"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Deck   *deck.Deck
	Player *player.Player
	Dealer *dealer.Dealer
}

func NewGame(playerName string) *Game {
	deck := deck.NewDeck()
	playerTokens := LoadTokens()
	player := player.NewPlayer(playerName, false, playerTokens)
	dealer := dealer.NewDealer()
	return &Game{
		Deck:   deck,
		Player: player,
		Dealer: dealer,
	}
}

func LoadTokens() int {
	tokensTxt := "tokens.txt"
	data, err := os.ReadFile(tokensTxt)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	currTokens, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		log.Fatalf("Error converting file content to integer: %v", err)
	}

	return currTokens
}

func (g *Game) Start() {
	g.Player.AddCard(g.Deck.DealCard())
	g.Dealer.AddCard(g.Deck.DealCard())
	g.Player.AddCard(g.Deck.DealCard())
	g.Dealer.AddCard(g.Deck.DealCard())

	for _, card := range g.Player.Hand {
		fmt.Printf("Player's hand: %s\n\n", card.Rank)
	}

	fmt.Printf("Player's score %d\n\n", g.Player.Score)
	fmt.Printf("Dealer's face-up card: %s\n\n", g.Dealer.Hand[0].Rank)
}

func (g *Game) DetermineWinner(betAmount int) {
	if g.Player.Score > 21 {
		g.Player.UpdateTokens(betAmount, false)
		fmt.Println("Player busts!")
	} else if g.Dealer.Score > 21 {
		g.Player.UpdateTokens(betAmount, true)
		fmt.Println("Dealer busts! Player wins.")
	} else if g.Player.Score > g.Dealer.Score {
		g.Player.UpdateTokens(betAmount, true)
		fmt.Println("Player wins!")
	} else if g.Player.Score < g.Dealer.Score {
		g.Player.UpdateTokens(betAmount, false)
		fmt.Println("Dealer wins!")
	} else {
		fmt.Println("It's a tie!")
	}

	g.Player.ResetHand()
	g.Dealer.ResetHand()
}
