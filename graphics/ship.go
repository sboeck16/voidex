/*
Package graphics provides method to generate or load pictures, as well as the
picture directly
*/
package graphics

import (
	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

/*
GenShip generates a ship with "size", "str" stroke width and color "clr".
*/
func GenShip(size, str float32, clr *Col) *eb.Image {
	img := eb.NewImage(int(size), int(size))
	vector.StrokeLine(img, size/2, size/2, 0, size, str, clr, false)
	vector.StrokeLine(img, size/2, size/2, size, size, str, clr, false)
	vector.StrokeLine(img, size/2, 0, 0, size, str, clr, false)
	vector.StrokeLine(img, size/2, 0, size, size, str, clr, false)
	return img
}

/*
Generates a circle image
*/
func GenCircle(size, str float32, clr *Col) *eb.Image {
	img := eb.NewImage(int(size), int(size))
	vector.StrokeCircle(img, size/2, size/2, size/2, str, clr, false)
	return img
}
