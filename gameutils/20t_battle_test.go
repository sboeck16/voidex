package gameutils

import (
	"math/rand"
	"testing"
	"time"
)

func TestBattle(t *testing.T) {
	deb("RUNNING: test battle")

	// create game
	board := NewBoard(testWidth, testHeight)
	game := NewGame(testWidth, testHeight)
	game.AddBoard(board)

	// add enemies
	fighterAmount := 1
	fdist := 10
	yCoords := 200

	// stresstest
	if !displayGame {
		fighterAmount *= 10
	}

	for i := 0; i < fighterAmount; i++ {
		s1 := getTestFighter(50, float64(yCoords)+float64(i*fdist), OwnerPlayer)
		s2 := getTestFighter(550, float64(yCoords)+float64(i*fdist), OwnerLost)
		board.RegisterObj(s1, s1, s1, s1, nil)
		board.RegisterObj(s2, s2, s2, s2, nil)
	}
	if displayGame {
		game.StartGame("TEST Battle")
	} else {
		s := time.Now()
		for i := 0; i < testTicks; i++ {
			board.Tick()
		}
		dur := time.Since(s)
		deb("-- #ticks, #ftr, #duration, #per tick",
			testTicks, fighterAmount, dur, dur/time.Duration(testTicks))
	}
}

func getTestFighter(x, y float64, faction int) *FightingObject {
	eFaction := OwnerPlayer
	if faction == OwnerPlayer {
		eFaction = OwnerLost
	}
	ret := &FightingObject{
		health:        5000,
		maxhealth:     30,
		accel:         0.05,
		speed:         1,
		corr:          .95,
		minDist:       200,
		agility:       10,
		scanRange:     10,
		targetFaction: eFaction,
		targetType:    ObjFighter,
	}
	ret.x = x
	ret.y = y
	ret.dx = (0.5 - rand.Float64()) * 2
	ret.dy = (0.5 - rand.Float64()) * 2
	ret.faction = faction
	ret.objType = ObjFighter
	ret.mass = 1
	ret.collSize = 10
	ret.myCollObj = &Collision{
		mass:     1,
		collSize: 10,
		faction:  faction,
		objType:  ObjFighter,
		foRef:    ret,
	}
	ret.SetImg(FactionGraphics[faction][ObjFighter])
	launcher := GetDefaultRocketLauncher(10, 10)
	ret.AddWeapon(launcher)
	return ret
}
