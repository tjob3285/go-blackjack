package card

import (
	"testing"
)

// Test NewCard function
func TestNewCard(t *testing.T) {
	tests := []struct {
		rank   string
		suit   string
		value  int
		expect *Card
	}{
		{
			rank:  "Ace",
			suit:  "Spades",
			value: 1,
			expect: &Card{
				Rank:  "Ace",
				Suit:  "Spades",
				Value: 1,
			},
		},
		{
			rank:  "King",
			suit:  "Hearts",
			value: 13,
			expect: &Card{
				Rank:  "King",
				Suit:  "Hearts",
				Value: 13,
			},
		},
		{
			rank:  "10",
			suit:  "Diamonds",
			value: 10,
			expect: &Card{
				Rank:  "10",
				Suit:  "Diamonds",
				Value: 10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.rank+"-"+tt.suit, func(t *testing.T) {
			got := NewCard(tt.rank, tt.suit, tt.value)
			if got.Rank != tt.expect.Rank || got.Suit != tt.expect.Suit || got.Value != tt.expect.Value {
				t.Errorf("NewCard() = %v, want %v", got, tt.expect)
			}
		})
	}
}
