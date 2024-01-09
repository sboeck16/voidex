package gameutils

import (
	"testing"
	"voidex/graphics"
)

func TestCollide(t *testing.T) {

	deb("RUNNING: collide test")

	maxTicks := 2000

	// board, game
	board := NewBoard(testWidth, testHeight)
	game := NewGame(testWidth, testWidth)
	game.AddBoard(board)

	var s1 *testShip
	s1 = newTestBall(300, 100, -1, 0, 1)
	board.RegisterObj(s1, s1, s1, s1, nil)
	forImpTest := s1
	s1 = newTestBall(200, 100, 1, 0, 1)
	board.RegisterObj(s1, s1, s1, s1, nil)

	s1 = newTestBall(200, 200, 0, 0, 1)
	board.RegisterObj(s1, s1, s1, s1, nil)
	s1 = newTestBall(100, 205, 1, 0, 1)
	board.RegisterObj(s1, s1, s1, s1, nil)

	for i := 0; i < 2; i++ {
		s1 = newTestBall(200, 300+float64(i), 0, 0, 1)
		board.RegisterObj(s1, s1, s1, s1, nil)
	}

	if displayGame {
		game.StartGame("TEST Collision")
	} else {
		for i := 0; i <= maxTicks; i++ {
			board.Tick()
		}
	}
	s1 = forImpTest
	if int(s1.x*100) != 231799 || int(s1.dx*100) != 105 {
		t.Error("collsion moved different", int(s1.x*100), int(s1.dx*100))
	}

}

// #############################################################################
// 								Example Ship
// #############################################################################

type testShip struct {
	ClickObject // inherits CollidingObject
	moveTo      bool
	tx, ty      float64
	maxSpeed    float64
	accel       float64
	tick        int
}

func newTestBall(x, y, dx, dy, m float64) *testShip {
	ret := newTestO(x, y, dx, dy, m)
	ret.SetImg(graphics.GenCircle(20, 2, graphics.ColorGreen))
	return ret
}

func newTestShip(x, y, dx, dy, m float64, clr *graphics.Col) *testShip {
	ret := newTestO(x, y, dx, dy, m)
	ret.SetImg(graphics.GenShip(20, 2, clr))
	return ret
}

func newTestO(x, y, dx, dy, m float64) *testShip {
	ret := new(testShip)
	ret.mass = m
	ret.x = x
	ret.y = y
	ret.dx = dx
	ret.dy = dy
	ret.myCollObj = &Collision{mass: m, collSize: 10}
	ret.collSize = 10
	return ret
}

func (ts *testShip) Update() {
	ts.ResolveImpacts(0)
	ts.collideEvents = nil
	if ts.moveTo {
		if ts.tick%6 == 0 {
			ts.RotMove(ts.tx, ts.ty, ts.accel, ts.maxSpeed, 0.95)
		}
		ts.tick++
	}
}
