package main

type CardSuit = int64

const (
	SuitClub CardSuit = iota + 1
	SuitDiamond
	SuitHeart
	SuitSpade
)

type Card struct {
	Suit  CardSuit
	Value int
}

func newCard(cardSuit CardSuit, cardValue int) *Card {
	card := Card{Suit: cardSuit, Value: cardValue}
	return &card
}
func CheckTheInput(cardOne Card, cardTwo Card) bool {
	theInputIsCorrect := true
	if (cardOne.Value < 2 || cardOne.Value > 13) || (cardTwo.Value < 2 || cardTwo.Value > 13) || (cardOne.Suit < 1 || cardOne.Suit > 4) || (cardTwo.Suit < 1 || cardTwo.Suit > 4) {
		theInputIsCorrect = false
	}
	return theInputIsCorrect
}
func CompareCards(cardOne Card, cardTwo Card) int {
	if CheckTheInput(cardOne, cardTwo) == true {
		if cardOne.Value < cardTwo.Value || (cardOne.Value == cardTwo.Value && cardOne.Suit < cardTwo.Suit) {
			return -1
		} else if cardOne.Value > cardTwo.Value || (cardOne.Value == cardTwo.Value && cardOne.Suit > cardTwo.Suit) {
			return 1
		} else {
			return 0
		}
	} else {
		return -2
	}
}
func maxCard(cards []Card) Card {
	maxCard := Card{Suit: SuitDiamond, Value: 2}

	for i := range cards {
		if CompareCards(maxCard, cards[i]) == -1 {
			maxCard.Suit = cards[i].Suit
			maxCard.Value = cards[i].Value
		}
	}
	return maxCard
}
