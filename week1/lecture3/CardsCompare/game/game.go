package game

import (
	"errors"
	"fmt"

	"week1_lecture3/cardSuit"
)

func CompareCards(cardOneVal int64, cardOneSuit cardSuit.CardSuit, cardTwoVal int64, cardTwoSuit cardSuit.CardSuit) (int64, error) {
	if cardOneVal < 1 || cardOneVal > 13 {
		fmt.Printf("The value of card is outside of range! Please, enter value from 1 to 13!")
		return cardOneVal, errors.New("The value of card one is outside of range! Please, enter value from 1 to 13!")
	}
	if cardTwoVal < 1 || cardTwoVal > 13 {
		fmt.Printf("The value of card is outside of range! Please, enter value from 1 to 13!")
		return cardTwoVal, errors.New("The value of card two is outside of range! Please, enter value from 1 to 13!")
	}
	if cardOneSuit < 1 || cardOneSuit > 4 {
		fmt.Printf("The value of suit is not correct!")
		return cardOneSuit, errors.New("The value of card one suit is not correct!")
	}
	if cardTwoSuit < 1 || cardTwoSuit > 4 {
		fmt.Printf("The value of suit is not correct!")
		return cardTwoSuit, errors.New("The value of card two suit is not correct!")
	}

	if cardOneVal < cardTwoVal {

		return -1, nil
	} else if cardOneVal > cardTwoVal {
		return 1, nil
	} else {
		if cardOneSuit < cardTwoSuit {
			return -1, nil
		} else if cardOneSuit > cardTwoSuit {
			return 1, nil
		} else {
			return 0, nil
		}
	}
}
