package main

/*
GameStats holds all values for a voidex idle game. Used to import and export
game states.
*/
type GameStats struct {
	// holds all resources like matter and science
	Ressources map[int]float64

	// buildings
	Buildings       map[int]int
	BuildingsActive map[int]int

	// unlocked science nodes
	UnlockedScience []int

	// universum seed and explored nodes
	UniversumSeed int64
	ExploredNodes []int

	// ship building
	FighterLoadOut *FighterLoadout
	CapitalLoadOut *CapitalLoadOut
	StationLoadout *StationLoadout

	// not safed but stored here

	// states
	battleState  int
	displayState int
	// used to enable buttons
	buttonLevel int

	// normal, science,
	displayMode int

	// max Ressources
	maxResources map[int]float64
}

/*
NewGameStats returns an empty GameStats but will preinitialze maps.
*/
func NewGameStats() *GameStats {
	ret := new(GameStats)

	ret.Ressources = make(map[int]float64)
	ret.Buildings = make(map[int]int)
	ret.BuildingsActive = make(map[int]int)
	ret.maxResources = make(map[int]float64)

	ret.FighterLoadOut = GetDefaultFtrLoadout()
	ret.CapitalLoadOut = GetDefaultCapLoadout()
	ret.StationLoadout = GetDefaultStaLoadout()

	return ret
}

/*
FighterLoadout holds the constructed fighter.
*/
type FighterLoadout struct {
	Inventory map[int]int
}

/*
CapitalLoadout holds the carrier/battleship class.
*/
type CapitalLoadOut struct {
	Inventory map[int]int
}

/*
StationLoadout holds the mobile station class.
*/
type StationLoadout struct {
	Inventory map[int]int
}
