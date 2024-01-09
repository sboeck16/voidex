package main

var (
	// holds the cost for first buy building->resource->mult
	costStart = map[int]map[int]float64{
		matter + collectorBuilding: {
			fabric: 5,
		},
	}

	// holds the multiplicator by level building->resource->mult
	costMult = map[int]map[int]float64{
		matter + collectorBuilding: {
			fabric: 1.5,
		},
	}

	// createCost holds the manual transformation costs
	createCost = map[int]map[int]float64{
		matter: {},
		fabric: {
			matter: 5,
		},
		rareMatter: {
			fabric: 5,
			matter: 20,
		},
		components: {
			rareMatter: 5,
			fabric:     20,
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
		// science
		science: {},
	}

	// inventory cost
	inventoryCost = map[int]map[int]float64{
		ftrEngines: {
			components: 1,
			fabric:     5,
			matter:     20,
		},
		ftrArmor: {
			matter: 100,
			fabric: 20,
		},
		blaster: {
			components: 1,
			fabric:     10,
		},
		missiles: {
			components:   3,
			exoticMatter: 1,
			matter:       10,
		},
		torpedos: {
			components:   2,
			exoticMatter: 5,
			fabric:       20,
		},
		CPU: {
			moduls: 5,
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
