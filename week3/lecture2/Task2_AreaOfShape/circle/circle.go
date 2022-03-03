package circle

import "math"

type Circle struct {
	Radius float64
	Area   float64
}

func (circle *Circle) NewCircle() *Circle {
	return &Circle{}
}
func (circle Circle) CalcArea() float64 {
	circle.Area = circle.Radius * circle.Radius * math.Pi

	return circle.Area
}
