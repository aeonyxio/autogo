package point

type Point struct {
	X, Y int
}

func New(x int, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}