package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"week1_lecture3/cardSuit"
	"week1_lecture3/game"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter the value of the first card")
	scanner.Scan()

	cardOneVal, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("Enter the suit of the first card")
	scanner.Scan()
	var cardOneSuit cardSuit.CardSuit
	cardOneSuit, _ = strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Printf("Enter the value of the second card")
	scanner.Scan()
	cardTwoVal, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("Enter the suit of the second card")
	scanner.Scan()
	var cardTwoSuit cardSuit.CardSuit
	cardTwoSuit, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	cardTwoSuit = cardSuit.CardSuit(cardTwoSuit)

	output, _ := game.CompareCards(cardOneVal, cardOneSuit, cardTwoVal, cardTwoSuit)
	switch output {
	case -1:
		fmt.Printf("The first card has a lower strength than the second one! Result is %d", output)
	case 0:
		fmt.Printf("Both cards are equal! Result is %d", output)
	case 1:
		fmt.Printf("The second card has a greater strength than the first one! Result is %d", output)

	}

}
