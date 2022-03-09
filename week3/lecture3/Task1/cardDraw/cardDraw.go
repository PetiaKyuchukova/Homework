package cardDraw

import (
	"cardGame/cardGame"
	"errors"
)

type dealer interface {
	Deal() *cardGame.Card
	ToSlice() []cardGame.Card
	IsEmpty() bool
}

func DrawAllCards(dealer dealer) ([]cardGame.Card, error) {
	if dealer.IsEmpty() == true {
		return nil, errors.New("The deck is empty")
	} else {
		dealer.Deal()
		return dealer.ToSlice(), nil
	}

}
