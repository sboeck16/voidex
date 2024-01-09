package gameutils

import (
	"math"
)

/*
MovingObject holds information needed to reposition the object (a speed vector
and border from its own board chunk)
*/
type MovingObject struct {
	PosObject

	// chunkBorders for boundary leap detection
	chunkMinX float64
	chunkMinY float64
	chunkMaxX float64
	chunkMaxY float64

	// movement
	dx float64
	dy float64
}

/*
UpdatePosition moves an object
*/
func (m *MovingObject) UpdatePosition() {

	m.x += m.dx
	m.y += m.dy
	if m.x < m.chunkMinX || m.x > m.chunkMaxX ||
		m.y < m.chunkMinY || m.y > m.chunkMaxY {
		m.myBoard.LeftChunk(m)
	}
}

/*
RotMove provides an update function that uses target coordinates an acceleration
a maximum speed and a correction factor to recalculate the dx and dy values.
*/
func (m *MovingObject) RotMove(tx, ty, accel, max, corr float64) {
	// reduce drift
	m.dx = m.dx * corr
	m.dy = m.dy * corr

	// diff vektor
	cx := tx - m.x
	cy := ty - m.y

	clen := math.Sqrt(cx*cx + cy*cy)
	normalx := cx / clen
	normaly := cy / clen

	m.dx += normalx * accel
	m.dy += normaly * accel

	speed := math.Sqrt(m.dx*m.dx + m.dy*m.dy)
	if max > 0 && speed > max {
		m.dx = max * normalx
		m.dy = max * normaly
	}

	m.rotation = math.Pi - math.Atan2(m.dx, m.dy)
}
