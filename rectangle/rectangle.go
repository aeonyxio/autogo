package rectangle

import (
	"image"

	"github.com/aeonyxio/autogo/point"
)

type Rectangle struct{
	X, Y, Width, Height int
}

func New(x int, y int, width int, height int) *Rectangle {
	return &Rectangle {
		X: x,
		Y: y,
		Width: width,
		Height: height,
	}
}


func CenterOf(rectangle *Rectangle) *point.Point {
	return &point.Point {
		X: rectangle.X + rectangle.Width/2,
		Y: rectangle.Y + rectangle.Height/2,
	}
}

func ImageRectangleToRectangle(imgRect *image.Rectangle) *Rectangle {
	rectangle := New(imgRect.Min.X, imgRect.Min.Y, imgRect.Max.X - imgRect.Min.X, imgRect.Max.Y - imgRect.Min.Y )
	return rectangle
}

func RectangleToImageRectangle(rectangle *Rectangle) *image.Rectangle {
	imgRect := image.Rect(int(rectangle.X), int(rectangle.Y), int(rectangle.X) + int(rectangle.Width), int(rectangle.Y) + int(rectangle.Height))
	return &imgRect
}

func (r *Rectangle) LeftX() int{
	return r.X
}

func (r *Rectangle) RightX() int{
	return r.X + r.Width
}

func (r *Rectangle) TopY() int{
	return r.Y
}

func (r *Rectangle) BottomY() int{
	return r.Y + r.Height
}
