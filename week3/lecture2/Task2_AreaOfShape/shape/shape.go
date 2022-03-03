package shape

type Shape interface {
	CalcArea() float64
}

func GetArea(s Shape) float64 {

	return s.CalcArea()

}

type Shapes []Shape

func (s Shapes) LargestArea() float64 {
	var maxArea float64

	for i := 0; i < len(s); i++ {

		curentArea := s[i].CalcArea()

		if curentArea > float64(maxArea) {
			maxArea = curentArea
		}
	}
	return maxArea

}
