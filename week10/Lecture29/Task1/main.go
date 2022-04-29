package main

import "fmt"

type Order struct {
	Customer string
	Amount   int
}

func fetchOrders() []Order {
	return []Order{{Customer: "John", Amount: 1000}, {Customer: "Sara", Amount: 2000}, {Customer: "Sara", Amount: 1800}, {Customer: "John", Amount: 1200}}
}

func GroupBy[T any, U comparable](col []T, keyFn func(T) U) map[U][]T {
	outs := map[U][]T{}
	for _, v := range col {
		outs[keyFn(v)] = append(outs[keyFn(v)], v)
	}
	return outs
}

func main() {

	results := GroupBy(fetchOrders(), func(o Order) string { return o.Customer })

	fmt.Println(fetchOrders())
	fmt.Println(results)
	fmt.Println(results["Sara"])
	fmt.Println(results["John"])

}
