package main

// #############################################################################
// #							Base cost
// #############################################################################
var (

	// base cost for everything
	baseCost = map[int]map[int]float64{
		matter: {},
		fabric: {
			matter: 5,
		},
		rareMatter: {
			fabric: 5,
			matter: 10,
		},
		components: {
			rareMatter: 5,
			fabric:     10,
			matter:     100,
		},
		moduls: {
			components: 20,
			rareMatter: 20,
		},
		exoticMatter: {
			fabric:     20,
			rareMatter: 200,
		},
		structures: {
			fabric:     2000,
			components: 20,
			moduls:     10,
		},

		// buildings
		matter + collectorBuilding: {
			fabric: 3,
		},
		fabric + collectorBuilding: {
			rareMatter: 3,
			fabric:     5,
		},
		rareMatter + collectorBuilding: {
			components: 3,
			rareMatter: 5,
			fabric:     50,
		},
		components + collectorBuilding: {
			moduls:     3,
			components: 10,
			matter:     300,
		},
		moduls + collectorBuilding: {
			exoticMatter: 3,
			moduls:       10,
		},
		exoticMatter + collectorBuilding: {
			structures:   3,
			exoticMatter: 10,
		},
		structures + collectorBuilding: {
			structures: 10,
			moduls:     20,
			components: 30,
		},
		// ships
		fighters + shipBuildingMax: {
			components: 10,
			matter:     100,
		},
		fighters + collectorBuilding: {
			moduls:     0.8,
			components: 10,
			fabric:     100,
		},
	}
)

// #############################################################################
// #							Multiply
// #############################################################################
var (
	// holds the multiplicator by level building->resource->mult
	costMult = map[int]map[int]float64{
		matter + collectorBuilding: {
			fabric: 1.1,
		},
		// buildings
		fabric + collectorBuilding: {
			rareMatter: 1.1,
			fabric:     1.1,
		},
		rareMatter + collectorBuilding: {
			components: 1.1,
			rareMatter: 1.1,
			fabric:     1.1,
		},
		components + collectorBuilding: {
			moduls:     1.1,
			components: 1.1,
			matter:     1.1,
		},
		moduls + collectorBuilding: {
			exoticMatter: 1.5,
			moduls:       1.5,
		},
		exoticMatter + collectorBuilding: {
			structures:   1.5,
			exoticMatter: 1.5,
		},
		structures + collectorBuilding: {
			structures: 2,
			moduls:     2,
			components: 2,
		},
		// ships
		fighters + shipBuildingMax: {
			components: 5,
			matter:     5,
		},
		fighters + collectorBuilding: {
			moduls:     3,
			components: 3,
			fabric:     3,
		},
	}
)

// #############################################################################
// #							Inventory
// #############################################################################
var (
	// inventory cost
	inventoryCost = map[int]map[int]float64{
		ftrEngines: {
			components: 1,
			fabric:     5,
			matter:     20,
			buildTime:  3,
		},
		ftrArmor: {
			matter:    100,
			fabric:    20,
			buildTime: 1,
		},
		blaster: {
			components: 1,
			fabric:     10,
			buildTime:  2,
		},
		missiles: {
			components:   3,
			exoticMatter: 1,
			matter:       10,
			buildTime:    3,
		},
		torpedos: {
			components:   2,
			exoticMatter: 5,
			fabric:       20,
			buildTime:    5,
		},
		CPU: {
			moduls:    5,
			buildTime: 5,
		},
		agilityBoost: {},
		targetBoost:  {},
	}
)

// #############################################################################
// #							Loadout
// #############################################################################

/*
GetDefaultFtrLoadout returns default fighter loadout.
*/
func GetDefaultFtrLoadout() *FighterLoadout {
	ret := new(FighterLoadout)
	ret.Inventory = map[int]int{
		ftrEngines: 3,
		ftrArmor:   3,
		blaster:    2,
	}
	return ret
}

/*
GetDefaultCapLoadout return default capital loadout.
*/
func GetDefaultCapLoadout() *CapitalLoadOut {
	ret := new(CapitalLoadOut)
	ret.Inventory = map[int]int{
		capEngines: 1,
		capArmor:   3,
		hangar:     3,
		factory:    1,
		cannon:     1,
		missiles:   4,
	}
	return ret
}

/*
GetDefaultStaLoadout() return default station loadout.
*/
func GetDefaultStaLoadout() *StationLoadout {
	ret := new(StationLoadout)
	ret.Inventory = map[int]int{
		staEngines: 1,
		staArmor:   3,
		docking:    3,
		assembly:   1,
		beams:      1,
	}
	return ret
}
