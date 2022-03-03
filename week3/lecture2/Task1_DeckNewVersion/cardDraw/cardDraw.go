package cardDraw

import "cardGame/cardGame"

type dealer interface {
	Deal() *cardGame.Card
	ToSlice() []cardGame.Card
}

func DrawAllCards(dealer dealer) []cardGame.Card {
	dealer.Deal()
	return dealer.ToSlice()
}
