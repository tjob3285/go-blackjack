package deck

import "testing"

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	// Check if the deck has 6 sets of 52 cards
	expectedCards := 6 * 52
	if len(d.Cards) != expectedCards {
		t.Errorf("Expected deck to have %d cards, but got %d", expectedCards, len(d.Cards))
	}

	// Check if the deck contains the expected card ranks and suits
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}

	// Validate that all ranks and suits are present
	for _, suit := range suits {
		for _, rank := range ranks {
			cardFound := false
			for _, card := range d.Cards {
				if card.Rank == rank && card.Suit == suit {
					cardFound = true
					break
				}
			}
			if !cardFound {
				t.Errorf("Card with rank %s and suit %s was not found in the deck", rank, suit)
			}
		}
	}
}

func TestDeckShuffle(t *testing.T) {
	d1 := NewDeck()
	d2 := NewDeck()

	// If shuffle works, the decks should not be in the same order after being shuffled
	sameOrder := true
	for i := range d1.Cards {
		if d1.Cards[i] != d2.Cards[i] {
			sameOrder = false
			break
		}
	}

	if sameOrder {
		t.Errorf("Deck was not shuffled, the order of cards is the same")
	}
}

func TestDealCard(t *testing.T) {
	d := NewDeck()

	// Initial deck size should be 6 * 52 cards
	initialDeckSize := len(d.Cards)

	// Deal the first card
	card := d.DealCard()
	if card == nil {
		t.Errorf("Expected to deal a card, but got nil")
	}

	// Deck size should decrease by 1 after dealing one card
	if len(d.Cards) != initialDeckSize-1 {
		t.Errorf("Expected deck size to be %d, but got %d", initialDeckSize-1, len(d.Cards))
	}

	// Check if the dealt card is actually removed from the deck
	if d.Cards[0] == card {
		t.Errorf("Card was not removed from the deck")
	}

	// Deal all remaining cards
	for i := 0; i < initialDeckSize-1; i++ {
		_ = d.DealCard()
	}

	// Deal one more card when deck is empty
	card = d.DealCard()
	if card != nil {
		t.Errorf("Expected to receive nil when deck is empty, but got a card")
	}
}

func TestGetCardValue(t *testing.T) {
	tests := []struct {
		rank  string
		value int
	}{
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
		{"9", 9},
		{"10", 10},
		{"J", 10},
		{"Q", 10},
		{"K", 10},
		{"A", 11},
	}

	for _, tt := range tests {
		t.Run(tt.rank, func(t *testing.T) {
			got := getCardValue(tt.rank)
			if got != tt.value {
				t.Errorf("getCardValue(%s) = %d; want %d", tt.rank, got, tt.value)
			}
		})
	}
}
