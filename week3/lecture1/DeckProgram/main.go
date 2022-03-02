package main

import (
	"fmt"
	"math/rand"
)

type Deck struct {
	LastCard *Card
}
type Card struct {
	Suit     CardSuit
	Value    int
	PrevCard *Card
}

type CardSuit = int

const (
	SuitClub CardSuit = iota + 1
	SuitDiamond
	SuitHeart
	SuitSpade
)

func (deck *Deck) New() {
	for i := 2; i < 15; i++ {
		for s := 1; s < 5; s++ {
			newCard := Card{Suit: s, Value: i, PrevCard: deck.LastCard}
			deck.LastCard = &newCard
		}
	}
}
func (deck *Deck) toSlice() []Card {

	slice := []Card{}
	currentElement := deck.LastCard

	for currentElement != nil {
		slice = append(slice, *currentElement)
		currentElement = currentElement.PrevCard
	}

	return slice
}
func (deck *Deck) Shuffle() [52]*Card {
	arr := [52]*Card{}

	temp := deck.LastCard

	for i := 0; i < 52; i++ {
		arr[i] = temp
		temp = temp.PrevCard
	}

	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	deck.LastCard = arr[0]
	temp = deck.LastCard

	for i := range arr {
		temp = arr[i]

		if i == 51 {
			temp.PrevCard = nil
		} else {
			temp.PrevCard = arr[i+1]
		}

		temp = temp.PrevCard
	}

	return arr
}
func (deck *Deck) Deal() *Card {
	lastCard := deck.LastCard

	if deck.LastCard != nil {
		deck.LastCard = deck.LastCard.PrevCard
	}

	return lastCard
}

func main() {
	deck := &Deck{}
	deck.New()
	fmt.Println("Before the shuffle: \n", deck.toSlice())
	deck.Shuffle()
	fmt.Print("\n")
	fmt.Println("After the shuffle: \n", deck.toSlice())
	fmt.Print("\n")
	deck.Deal()
	fmt.Println("After Deal:\n ", deck.toSlice())

}
