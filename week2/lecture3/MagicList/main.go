package main

import "fmt"

type Item struct {
	Value    int
	PrevItem *Item
}

type MagicList struct {
	LastItem *Item
}

func main() {
	l := &MagicList{}
	add(l, 10)
	fmt.Printf("Last element is: %d; Prev element is: %d\n", l.LastItem.Value, l.LastItem.PrevItem)
	add(l, 22)
	fmt.Printf("Last element is: %d; Prev element is: %d\n", l.LastItem.Value, l.LastItem.PrevItem.Value)
	add(l, 44)
	fmt.Printf("Last element is: %d; Prev element is: %d\n", l.LastItem.Value, l.LastItem.PrevItem.Value)
	add(l, 78)
	fmt.Printf("Last element is: %d; Prev element is: %d\n", l.LastItem.Value, l.LastItem.PrevItem.Value)

	fmt.Printf("The slice is: %d", toSlice(l))

}

func add(l *MagicList, value int) {

	i := Item{Value: value}

	if l.LastItem == nil {
		l.LastItem = &i
	} else {
		i.PrevItem = l.LastItem
		l.LastItem = &i
	}
}

func toSlice(l *MagicList) []int {

	slice := []int{}
	currentElement := l.LastItem

	for currentElement != nil {
		slice = append(slice, currentElement.Value)
		currentElement = currentElement.PrevItem
	}

	return slice
}
