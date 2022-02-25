package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter the value of the first card")
	scanner.Scan()

	cardOneValue, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Printf("Enter the suit of the first card")
	scanner.Scan()
	var cardOneSuit CardSuit
	cardOneSuit, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	cardOneSuit = CardSuit(cardOneSuit)

	fmt.Printf("Enter the value of the second card")
	scanner.Scan()
	cardTwoVal, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Printf("Enter the suit of the second card")
	scanner.Scan()
	var cardTwoSuit CardSuit
	cardTwoSuit, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	cardTwoSuit = CardSuit(cardTwoSuit)

	cardOne := newCard(cardOneSuit, int(cardOneValue))
	cardTwo := newCard(cardTwoSuit, int(cardTwoVal))
	sliceOfCards := []Card{*cardOne, *cardTwo}
	fmt.Println(sliceOfCards)

	output := maxCard(sliceOfCards)
	fmt.Printf("Result is %d", output)

	outputCom := CompareCards(*cardOne, *cardTwo)
	fmt.Printf("Result is %d", outputCom)

	// //switch output {
	// case -1:
	// 	fmt.Printf("The first card has a lower strength than the second one! Result is %d", output)
	// case 0:
	// 	fmt.Printf("Both cards are equal! Result is %d", output)
	// case 1:
	// 	fmt.Printf("The second card has a greater strength than the first one! Result is %d", output)
	// }

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
		return 000
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
