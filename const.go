package main

import (
	"voidex/gameutils"
	"voidex/graphics"
)

const (
	// ressources
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
	fighters
	capitals
	stations
	maxResources

	// buildings flag
	collectorBuilding       = 1 << 10
	collectorBuildingActive = 1 << 11
	shipBuilding            = 1 << 12
	shipBuildingMax         = 1 << 13
	collectorAdd            = 1 << 14
	collectorSub            = 1 << 15

	/*
		// buildings
		matterCollector = iota
		rareMatterCollector
		exoticMatterCollector
		materialFabricator
		materialFabricatorActive
		componentFactory
		componentFactoryActive
		modulCreator
		moduleCreatorActive
		structureBuilder
		structureBuilderActive
		fighterHangars
		fighterHangarsActive
		shipyards
		shipyardsActive
		stationConstructor
		stationConstructorActive
		fighterBays
		landingfields
		fleetCommands
		maxBuildings
		// for clicking
		clickMatterPseudo
		clickRarePseudo
		clickExoticPseudo
		clickComponentsPseudo
		clickModulsPseudo
		clickExoticMatterPseudo
		clickStructuresPseudo
	*/

	// ship inventory
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

	allBuildings = []int{
		matter + collectorBuilding,
		fabric + collectorBuilding,
		rareMatter + collectorBuilding,
		components + collectorBuilding,
		moduls + collectorBuilding,
		exoticMatter + collectorBuilding,
		structures + collectorBuilding,
		fighters + collectorBuilding,
		capitals + collectorBuilding,
		stations + collectorBuilding,
		fighters + shipBuildingMax,
		capitals + shipBuildingMax,
		stations + shipBuildingMax,
	}
)

// #############################################################################
// #							Display
// #############################################################################

const (
	statDisplayDivide = ": "
	statDisplayFormat = 'f'
	statDisplayPrec   = 1

	statDisplayWidth        = 160
	statDisplaySpace        = 10
	statDisplayHeight       = 18
	statDisplayNextRowSpace = 3

	actionButtonWidth        = 160
	actionButtonAddSubWidth  = 20
	actionButtonHeight       = 18
	actionButtonNextRowSpace = 5
	actionButtonSpace        = 10
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
)