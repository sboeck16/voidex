package gameutils

import (
	"testing"
	"voidex/graphics"

	eb "github.com/hajimehoshi/ebiten/v2"
)

var ()

func TestClicking(t *testing.T) {

	deb("RUNNING: click test")

	// crete board
	board := NewBoard(testWidth, testHeight)
	board.ClickHandle = boardclick

	// create game
	game := NewGame(testWidth, testHeight)
	game.AddBoard(board)

	if displayGame {
		deb("test click use right and left click, and close when ready")
		game.StartGame("Test Clicking")
	} else {
		board.BoardClicked(100, 100, eb.MouseButtonLeft)
		board.BoardClicked(300, 100, eb.MouseButtonLeft)
		for i := 0; i < 200; i++ {
			board.Tick()
		}
		if len(board.chunks) < 2 {
			t.Error("setting ships via click and init moving didnt work")
		}
	}
}

func boardclick(b *Board, x, y float64, but eb.MouseButton) {

	if but == eb.MouseButtonLeft {
		s := newTestSearcher(x, y, 0, 0, 1, graphics.ColorGreen)
		s.clickHandle = shipClick
		b.RegisterObj(s, s, s, s, s)
	}
	if but == eb.MouseButtonRight {
		if lastTarget != nil {
			lastTarget.removeMe = true
		}
		tar := new(PosObject)
		tar.faction = OnwerRobots
		tar.objType = ObjFighter
		tar.x = x
		tar.y = y
		b.RegisterObj(tar, nil, nil, nil, nil)
		lastTarget = tar
	}
}

func shipClick(cl *ClickObject, but eb.MouseButton) {
	deb(cl.x, cl.myChunk.x)
}

// #############################################################

type testSearcher struct {
	ClickObject // inherits CollidingObject
	maxSpeed    float64
	accel       float64
	tick        int
	target      PositionAble
}

func newTestSearcher(x, y, dx, dy, m float64, clr *graphics.Col) *testSearcher {
	ret := new(testSearcher)
	ret.SetImg(graphics.GenShip(40, 2, clr))
	ret.mass = m
	ret.x = x
	ret.y = y
	ret.dx = dx
	ret.dy = dy
	ret.accel = .5
	ret.maxSpeed = 5
	ret.myCollObj = &Collision{mass: m, collSize: 20}
	ret.collSize = 20
	return ret
}

func (tse *testSearcher) Update() {
	tse.ResolveImpacts(0)
	tse.collideEvents = nil

	if tse.tick%6 == 0 {
		if tse.target == nil || tse.target.Remove() {
			tse.target = tse.myBoard.GetNearestObject(
				tse.x, tse.y, OnwerRobots, ObjFighter, searchRange)
		}

		if tse.target != nil {
			tx, ty := tse.target.GetCoords()
			tse.RotMove(tx, ty, tse.accel, tse.maxSpeed, 0.95)
		}
	}
	tse.tick++
}
