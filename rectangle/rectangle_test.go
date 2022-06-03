package rectangle

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRectangle(t *testing.T) {
	assert := assert.New(t)

	rectangle := New(10, 15, 500, 1000)

  	assert.Equal(rectangle.X, 10, "they should be equal")
  	assert.Equal(rectangle.Y, 15, "they should be equal")
  	assert.Equal(rectangle.Width, 500, "they should be equal")
  	assert.Equal(rectangle.Height, 1000, "they should be equal")
  	assert.Equal(rectangle.LeftX(), 10, "they should be equal")
  	assert.Equal(rectangle.RightX(), 510, "they should be equal")
  	assert.Equal(rectangle.TopY(), 15, "they should be equal")
  	assert.Equal(rectangle.BottomY(), 1015, "they should be equal")
}

func TestCenterOf(t *testing.T) {
	assert := assert.New(t)

	point := CenterOf(New(10, 15, 500, 1000))

  	assert.Equal(point.X, 260, "they should be equal")
  	assert.Equal(point.Y, 515, "they should be equal")
}

func TestRectangleToImageRectangle(t *testing.T) {
	assert := assert.New(t)

	imgRect := RectangleToImageRectangle(New(10, 15, 500, 1000))

  	assert.Equal(imgRect.Min.X, 10, "they should be equal")
  	assert.Equal(imgRect.Min.Y, 15, "they should be equal")
  	assert.Equal(imgRect.Max.X, 510, "they should be equal")
  	assert.Equal(imgRect.Max.Y, 1015, "they should be equal")
}

func TestImageRectangleToRectangle(t *testing.T) {
	assert := assert.New(t)
	imgRect := image.Rect(10, 15, 510, 1015)
	rectangle := ImageRectangleToRectangle(&imgRect)

	assert.Equal(rectangle.X, 10, "they should be equal")
	assert.Equal(rectangle.Y, 15, "they should be equal")
	assert.Equal(rectangle.Width, 500, "they should be equal")
	assert.Equal(rectangle.Height, 1000, "they should be equal")
}
