package main

import (
	"voidex/gameutils"
	"voidex/graphics"
)

// #############################################################################
// #							Global Game
// #############################################################################

const (
	ticksPerSec = 60.0 // defined by ebitengine

	// UpdateOnEveryTick counter when game updates 6 -> every 0.1s
	updateOnEveryTick = 6.0
	// holds updates per second
	updatesPerSecond = ticksPerSec / updateOnEveryTick
)

// #############################################################################
// #							Ressources
// #############################################################################

const (
	// Factions
	// --------
	// OwnerNeutral does not interact
	factionNeutral = gameutils.OwnerNeutral
	// OwnerPlayer is player faction
	factionPlayer = gameutils.OwnerPlayer
	// OwnerLost pirates, lost and random
	factionLost = gameutils.OwnerLost
	// OwnerRobots nano robot faction
	factionRobots = gameutils.OnwerRobots
	// OwnerPlant void growth
	factionPlants = gameutils.OwnerPlants
)
const (
	// ressources
	// ----------
	// matter, ...
	matter = iota
	fabric
	rareMatter
	components
	moduls
	exoticMatter
	structures
	// knowledge
	science
	dimensionalExp
	redExp
	greenExp
	blueExp
	// saved ships
	fighters // always the first ship!
	capitals
	stations
	maxResources

	// for things that need longer to produce
	buildTime = 1 << 9

	// buildings flag
	collectorBuilding       = 1 << 10
	collectorBuildingActive = 1 << 11
	shipBuildingMax         = 1 << 12
	collectorAdd            = 1 << 13
	collectorSub            = 1 << 14
)
const (
	// ship inventory
	// --------------
	// Fighter
	ftrEngines = iota
	ftrArmor
	blaster
	missiles
	torpedos
	CPU
	agilityBoost
	targetBoost
	// Capital
	capEngines
	capArmor
	hangar
	factory
	cannon
	// Station
	staEngines
	staArmor
	docking
	assembly
	beams
)
const (
	// battle states
	noBattleState      = 0
	runningBattleState = -1

	// displayState
	displayMap = iota
	displayScience
	displayExp
	displayRed
	displayGreen
	displayBlue
	displayInfo
	displayBattle
	maxDisplayBoards
)
const (
	// numerical boni
	addToBase = iota
	increaseBasePerc
	multiply

	// actions
	buy = iota
	addActive
	subActive
)

var (
	toStringRessource = map[int]string{
		matter:         "matter",
		fabric:         "fabric",
		rareMatter:     "rareMatter",
		components:     "components",
		moduls:         "moduls",
		exoticMatter:   "exoticMatter",
		structures:     "structures",
		science:        "science",
		dimensionalExp: "dimensionalExp",
		redExp:         "redExp",
		greenExp:       "greenExp",
		blueExp:        "blueExp",
		fighters:       "fighters",
		capitals:       "capitals",
		stations:       "stations",
	}
	toStringBuildings = map[int]string{
		matter + collectorBuilding:                                 "matter collector",
		fabric + collectorBuilding:                                 "fabricator",
		rareMatter + collectorBuilding:                             "rare matter gatherer",
		components + collectorBuilding:                             "components factory",
		moduls + collectorBuilding:                                 "modul constructor",
		exoticMatter + collectorBuilding:                           "exotic matter finder",
		structures + collectorBuilding:                             "structure builder",
		matter + collectorBuilding + collectorBuildingActive:       "matter collector A",
		fabric + collectorBuilding + collectorBuildingActive:       "fabricator A",
		rareMatter + collectorBuilding + collectorBuildingActive:   "rare matter gatherer A",
		components + collectorBuilding + collectorBuildingActive:   "components factory A",
		moduls + collectorBuilding + collectorBuildingActive:       "modul constructor A",
		exoticMatter + collectorBuilding + collectorBuildingActive: "exotic matter finder A",
		structures + collectorBuilding + collectorBuildingActive:   "structure builder A",

		fighters + collectorBuilding:                           "fighter hangars",
		capitals + collectorBuilding:                           "shipyards",
		stations + collectorBuilding:                           "station constructor",
		fighters + collectorBuilding + collectorBuildingActive: "fighter hangars A",
		capitals + collectorBuilding + collectorBuildingActive: "shipyards A",
		stations + collectorBuilding + collectorBuildingActive: "station constructor A",
		fighters + shipBuildingMax:                             "fighter bays",
		capitals + shipBuildingMax:                             "landing fields",
		stations + shipBuildingMax:                             "fleet commands",
	}

	// first ressources, second ship, third ship max
	allBuildingsCollectable = []int{
		matter + collectorBuilding,
		fabric + collectorBuilding,
		rareMatter + collectorBuilding,
		components + collectorBuilding,
		moduls + collectorBuilding,
		exoticMatter + collectorBuilding,
		structures + collectorBuilding,
	}
	allBuildingsShips = []int{
		fighters + collectorBuilding,
		capitals + collectorBuilding,
		stations + collectorBuilding,
		fighters + shipBuildingMax,
		capitals + shipBuildingMax,
		stations + shipBuildingMax,
	}
	allBuildings = append(allBuildingsCollectable, allBuildingsShips...)

	resWithMax = map[int]int{
		fighters: 0,
		capitals: 0,
		stations: 0,
	}
)

// #############################################################################
// #							Display
// #############################################################################

const (
	// generic string constants
	sp = " "
	nl = "\n"

	// strings
	strInvalidCosts = "-invalid-"

	costSymbolAmountDivide = ":"
	costStringsJoin        = ","
	costDisplayFormat      = 'f'
	costDisplayPrec        = 1

	statDisplayDivide       = ": "
	statDisplayFormat       = 'f'
	statDisplayPrec         = 1
	statDisplayMaxResDivide = "|"

	statDisplayWidth        = 160
	statDisplaySpace        = 10
	statDisplayHeight       = 18
	statDisplayNextRowSpace = 3

	actionButtonActive1 = "("
	actionButtonActive2 = "|"
	actionButtonActive3 = ")"

	actionStrCollect = "collect"
	actionStrAdd     = "+"
	actionStrSub     = "-"

	actionButtonWidth        = 160
	actionButtonAddSubWidth  = 20
	actionButtonHeight       = 34
	actionButtonNextRowSpace = 5
	actionButtonSpace        = 10
	actionButtonColumnSpace  = 70
)

var (
	statDisplayText = map[int]string{
		matter:         "Matter",
		fabric:         "Fabric",
		rareMatter:     "Rare",
		components:     "Components",
		moduls:         "Moduls",
		exoticMatter:   "Exotic",
		structures:     "Structures",
		science:        "Science",
		dimensionalExp: "Dim",
		redExp:         "Red",
		greenExp:       "Green",
		blueExp:        "Blue",
		fighters:       "Fighters",
		capitals:       "Cap Ships",
		stations:       "Stations",
	}
	costSymbol = map[int]string{
		matter:         "Ma",
		fabric:         "Fa",
		rareMatter:     "Ra",
		components:     "Co",
		moduls:         "Mo",
		exoticMatter:   "Ex",
		structures:     "St",
		science:        "Sc",
		dimensionalExp: "Dim",
		redExp:         "Red",
		greenExp:       "Green",
		blueExp:        "Blue",
		fighters:       "Ftr",
		capitals:       "Cap",
		stations:       "Stn",
	}

	statDisplayFont    = gameutils.TextFontSmall
	statDisplayBGCol   = graphics.ColorBlack
	statDisplayTextCol = graphics.ColorWhite

	actionDisplayFont    = gameutils.CreateFontMust(13, 72)
	actionDisplayBGCol   = graphics.ColorBlack
	actionDisplayTextCol = graphics.ColorWhite
)
