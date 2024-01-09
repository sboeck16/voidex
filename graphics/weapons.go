package graphics

import (
	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

/*
GenBeam generates a beam image.
*/
func GenBeam(sx, sy, tx, ty, str float32, clr *Col) *eb.Image {
	w := float32(0)
	h := float32(0)
	if sx < tx {
		w = tx - sx
	} else {
		w = sx - tx
	}
	if sy < ty {
		h = ty - sy
	} else {
		h = sy - ty
	}
	img := eb.NewImage(int(w), int(h))
	vector.StrokeLine(img, 0, 0, w, h, str, clr, false)
	return img
}

/*
GenProj generates a projectile image.
*/
func GenProj(size float32, clr *Col) *eb.Image {
	img := eb.NewImage(int(size), int(size))
	vector.DrawFilledCircle(img, size/2, size/2, size/2, clr, false)
	return img
}

/*
GenRocket generates a a rocket image.
*/
func GenRocket(size, str float32, clr *Col) *eb.Image {
	img := eb.NewImage(int(size), int(size))
	vector.StrokeLine(img, size/2, 0, 0, size, str, clr, false)
	vector.StrokeLine(img, size/2, 0, size, size, str, clr, false)
	return img
}
