package square

type Square struct {
	Side float64
	Area float64
}

func (square *Square) NewSquare() *Square {
	return &Square{}
}

func (square Square) CalcArea() float64 {
	square.Area = square.Side * square.Side
	return square.Area
}
