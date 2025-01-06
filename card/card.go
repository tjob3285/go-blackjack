package card

// Card
type Card struct {
	Rank  string
	Suit  string
	Value int
}

func NewCard(rank, suit string, value int) *Card {
	return &Card{
		Rank:  rank,
		Suit:  suit,
		Value: value,
	}
}
