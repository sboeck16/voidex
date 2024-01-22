package main

import (
	"voidex/gameutils"
)

var (
	// holds the battle board
	battleBoard *gameutils.Board
	// holds the battle wrapper for utility access
	battle = new(Battle)
)

/*
Battle wraps a gameutils.Board on which to fleets fight. Provides utility and
setup methods. MAYBE move it to gameutils?
*/
type Battle struct {
	shipsByFaction map[int][]*gameutils.FightingObject
	board          *gameutils.Board
}

/*
SetUp uses provided board (will be cleared of all other objects) and adds ships
to it.
*/
func (b *Battle) SetUp(
	brd *gameutils.Board, ships []*gameutils.FightingObject) {

	b.board = brd
	b.board.Reset()
	b.shipsByFaction = make(map[int][]*gameutils.FightingObject)

	for _, ship := range ships {
		b.board.RegisterObj(ship, ship, ship, ship, nil)
		fac, _ := ship.GetFactionAndType()

		b.shipsByFaction[fac] = append(b.shipsByFaction[fac], ship)
	}
}

/*
IsAlive checks if faction has at least one active object.
*/
func (b *Battle) IsAlive(faction int) bool {
	for _, ship := range b.shipsByFaction[faction] {
		if !ship.Remove() {
			return true
		}
	}
	return false
}

/*
Tick progresses the battle one tick
*/
func (b *Battle) Tick() {
	b.board.Tick()
}

// #############################################################################
// #							UPDATE
// #############################################################################

func checkAndTickBattle(tick int) {

	if gameStats.battleState == runningBattleState {
		battle.Tick()
	}

	if tick%updateOnEveryTick == 0 {
		if !battle.IsAlive(factionPlayer) {
			playerLost()
		}
		if !battle.IsAlive(factionLost) &&
			!battle.IsAlive(factionRobots) &&
			!battle.IsAlive(factionPlants) {

			playerWon()
		}
	}
}

func playerLost() {}

func playerWon() {}
