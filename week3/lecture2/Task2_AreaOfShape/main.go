package main

import (
	"fmt"
	"shape/circle"
	"shape/shape"
	"shape/square"
)

func main() {
	myCircle := circle.Circle{Radius: 4.0}
	myCircle.NewCircle()
	myCircle.Area = shape.GetArea(myCircle)

	fmt.Printf("\nCircle area is %f", myCircle.Area)

	mySquare := square.Square{Side: 5.0}
	mySquare.NewSquare()
	mySquare.Area = shape.GetArea(mySquare)

	fmt.Printf("\nSquare are is %f", mySquare.Area)

	shapes := shape.Shapes{myCircle, mySquare}
	largestArea := shapes.LargestArea()
	fmt.Printf("\nThe largest area is %f", largestArea)

}
