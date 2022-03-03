package main

import (
	"cardGame/cardDraw"
	"cardGame/cardGame"
	"fmt"
)

func main() {
	deck := &cardGame.Deck{}
	deck.New()
	sliceB := deck.ToSlice()
	fmt.Println("%", sliceB)
	slice := cardDraw.DrawAllCards(deck)
	fmt.Println("%", slice)

}
