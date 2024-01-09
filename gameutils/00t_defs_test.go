package gameutils

var (
	// will fail if set true and multiple tests are run
	displayGame = false
	// used for non displayed test or to end display
	testTicks = 600
	// game/board size
	testHeight = 400
	testWidth  = 600

	// search for targets
	searchRange = 10
	lastTarget  *PosObject
)
