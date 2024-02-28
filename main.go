package main

import (
	"fmt"
	"voidex/gameutils"
	"voidex/graphics"

	eb "github.com/hajimehoshi/ebiten/v2"
)

var (
	// holds the state the game is in
	gameStats = NewGameStats()
	// holds the increase on big tick
	incrDecrOnBigUpd = []*cost{}
	// hold game global for easy access
	game *gameutils.Game

	// holds boards for easy access
	boardDisplayStat *gameutils.Board
	boardButtons     *gameutils.Board
	boardsMain       = map[int]*gameutils.Board{}
)

func main() {

	// start screen

	// load game

	// prepare game
	game = initGame()

	// set game logic
	game.SetTickFunction(gameUpdate)

	// start game
	game.StartGame("Void Explorer")

}

// print wrapper
func deb(i ...any) {
	fmt.Println(i...)
}

func initGame() *gameutils.Game {
	ret := gameutils.NewGame(width, height)

	ret.SetBGImg(getBGImage())

	// values board
	boardDisplayStat = gameutils.NewBoard(infoWidth, infoHeight)
	boardDisplayStat.GameX = infoOffX
	boardDisplayStat.GameY = infoOffY
	ret.AddBoard(boardDisplayStat)

	// buttons board
	boardButtons = gameutils.NewBoard(buttonsWidth, buttonsHeight)
	boardButtons.GameX = buttonsOffX
	boardButtons.GameY = buttonsOffY
	ret.AddBoard(boardButtons)

	// main boards
	for i := 0; i < maxDisplayBoards; i++ {
		boardsMain[i] = gameutils.NewBoard(mainWidth, mainHeight)
		boardsMain[i].GameX = mainOffX
		boardsMain[i].GameY = mainOffY
		ret.AddBoard(boardsMain[i])
	}

	// set initial display
	setDisplay(displayMap)

	// startup set TODO remove? -> loading?
	setStartUpGame()

	return ret
}

// #############################################################################
// #							TICK
// #############################################################################

func gameUpdate(tick int) {

	if tick%updateOnEveryTick == 0 {
		// update stats
		updateStats()

		// check if another thing needs to be enabled
		updateButtons()

		// update display
		updateDisplay()
	}

	// show battle
	checkAndTickBattle(tick)
}

// #############################################################################
// #							Access
// #############################################################################

func getBGImage() *eb.Image {
	ret := eb.NewImage(width, height)
	ret.Fill(graphics.ColorBlue)
	return ret
}

/*
Sets which main display is to be drawn.
*/
func setDisplay(active int) {
	for i := 0; i < maxDisplayBoards; i++ {
		boardsMain[i].DrawMe = i == active
	}
}

// #############################################################################
// #							MockUp
// #############################################################################

/*
utility function for testing visuals and progressing. creates gamestats TODO
*/
func setStartUpGame() {
	gameStats = NewGameStats()
	for i := 0; i < maxResources; i++ {
		gameStats.Ressources[i] = 100
	}
	for _, b := range allBuildingsShips {
		gameStats.Buildings[b] = 2
	}
}

/*
TODO! error handling and logging
*/
func checkError(errs ...error) bool {
	ret := false
	for _, err := range errs {
		if err != nil {
			ret = true
			deb(err.Error())
		}
	}
	return ret
}
