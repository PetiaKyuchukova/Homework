package main

import (
	"cardGame/cardDraw"
	"cardGame/cardGame"
	"fmt"
	"log"
)

func main() {
	deck := &cardGame.Deck{}
	//deck.New()

	sliceB := deck.ToSlice()
	fmt.Println("%", sliceB)

	slice, err := cardDraw.DrawAllCards(deck)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("%", slice)

}
