package main

import "testing"

func TestNewCard(t *testing.T) {
	expectedCard := Card{Suit: 3, Value: 3}
	card := newCard(3, 3)

	if expectedCard != *card {
		t.Error("ima erroro")
	}
}
func TestCheckTheInput_Case1(t *testing.T) {
	cardOne := Card{Suit: 6, Value: 1}
	cardTwo := Card{Suit: 6, Value: 1}
	expectedResult := false

	result := CheckTheInput(cardOne, cardTwo)
	if expectedResult != result {
		t.Errorf("Expected result: %t", expectedResult)
	}
}
func TestCheckTheInput_Case2(t *testing.T) {
	cardOne := Card{Suit: 0, Value: 15}
	cardTwo := Card{Suit: 0, Value: 15}
	expectedResult := false

	result := CheckTheInput(cardOne, cardTwo)
	if expectedResult != result {
		t.Errorf("Expected result: %t", expectedResult)
	}
}
func TestCompareCards_Case1(t *testing.T) {
	cardOne := Card{Suit: 3, Value: 3}
	cardTwo := Card{Suit: 3, Value: 3}

	expectedResult := 0

	result := CompareCards(cardOne, cardTwo)

	if expectedResult != result {
		t.Errorf("Expected result: %d", expectedResult)
	}

}
func TestCompareCards_Case2(t *testing.T) {
	cardOne := Card{Suit: 4, Value: 3}
	cardTwo := Card{Suit: 3, Value: 3}

	expectedResult := 1

	result := CompareCards(cardOne, cardTwo)

	if expectedResult != result {
		t.Errorf("Expected result: %d", expectedResult)
	}

}
func TestCompareCards_Case3(t *testing.T) {
	cardOne := Card{Suit: 3, Value: 3}
	cardTwo := Card{Suit: 4, Value: 3}

	expectedResult := -1

	result := CompareCards(cardOne, cardTwo)

	if expectedResult != result {
		t.Errorf("Expected result: %d", expectedResult)
	}

}
func TestCompareCards_Case4(t *testing.T) {
	cardOne := Card{Suit: 0, Value: 1}
	cardTwo := Card{Suit: 0, Value: 19}

	expectedResult := -2

	result := CompareCards(cardOne, cardTwo)

	if expectedResult != result {
		t.Errorf("Expected result: %d", expectedResult)
	}

}
func TestMaxCard(t *testing.T) {
	cardOne := Card{Suit: 4, Value: 3}
	cardTwo := Card{Suit: 4, Value: 8}
	cards := make([]Card, 3)
	cards[0] = cardOne
	cards[1] = cardTwo
	expectedResult := cardTwo

	result := maxCard(cards)
	if expectedResult != result {
		t.Errorf("Expected result: %d", expectedResult)
	}

}
