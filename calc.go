package main

import (
	"math"
)

var ()

func init() {

}

/*
holds info what to add to game stats and how
*/
type inc struct {
	add  map[int]float64
	cost []*cost
}

/*
creates new inc object and initializes maps.
*/
func newIncStruct() *inc {
	ret := new(inc)
	ret.add = make(map[int]float64)
	ret.cost = []*cost{}
	return ret
}

/*
holds the cost for a single item
*/
type cost struct {
	pay  map[int]float64
	get  float64
	item int
}

func updateStats() {
	// first add all without a cost
	for addRes, addVal := range incrDecrOnBigUpd.add {
		gameStats.Ressources[addRes] += addVal
	}

	// now add the ressource we can afford
	for _, cst := range incrDecrOnBigUpd.cost {
		// check
		canAfford := true
		for payRes, payAmount := range cst.pay {
			if gameStats.Ressources[payRes] < payAmount {
				canAfford = false
				break
			}
		}
		if !canAfford {
			continue
		}
		for payRes, payAmount := range cst.pay {
			gameStats.Ressources[payRes] -= payAmount
			gameStats.Ressources[cst.item] += cst.get
		}
	}
}

// #############################################################################
// #							CALC FUNCS
// #############################################################################

/*
calcCost calculates cost and returns a map with materials -> cost
*/
func calcCost(building, targetlevel int) map[int]float64 {
	startC, ok1 := costStart[building]
	multC, ok2 := costMult[building]
	if targetlevel < 1 || !ok1 || !ok2 {
		// MAYBE log error?
		return nil
	}

	ret := map[int]float64{}
	// copy start values
	for mat, cost := range startC {
		ret[mat] = cost
	}

	// multiply -> MAYBE more checks needed? -> performance?
	for mat, mult := range multC {
		ret[mat] = math.Ceil(math.Pow(mult, float64(targetlevel-1)) * ret[mat])
	}

	return ret
}

/*
calcUpdate recalculates the global update struct that holds all increasing or
decreasing values.

* base production, starts with 1, can be raised
* add %, all added and aplied to base production
* multiply, multiply through
*/
func calcUpdate() {

}
