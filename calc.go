package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
holds the cost for a single item
*/
type cost struct {
	ressAdd map[int]float64 // use negative for cost
	item    int
}

func updateStats() {

	// now add the ressource we can afford
	for _, cst := range incrDecrOnBigUpd {
		// check
		canAfford := true
		for payRes, payAmount := range cst.ressAdd {
			if gameStats.Ressources[payRes]+payAmount < 0 {
				canAfford = false
				break
			}
			if max, ok := resWithMax[payRes]; ok {
				if gameStats.Ressources[payRes] >= float64(max) {
					gameStats.Ressources[payRes] = float64(max)
					canAfford = false
				}
			}
		}
		if !canAfford {
			continue
		}
		for payRes, payAmount := range cst.ressAdd {
			gameStats.Ressources[payRes] += payAmount
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
	startC, ok1 := baseCost[building]
	multC, ok2 := costMult[building]
	if targetlevel < 1 || !ok1 {
		// MAYBE log error?
		return nil
	}

	ret := map[int]float64{}
	// copy start values
	for mat, cost := range startC {
		ret[mat] = cost
	}

	// no multipliers
	if !ok2 {
		return ret
	}

	// multiply -> MAYBE more checks needed? -> performance?
	for mat, mult := range multC {
		ret[mat] = math.Ceil(math.Pow(mult, float64(targetlevel-1)) * ret[mat])
	}

	return ret
}

/*
utility function to check and buy if possible
*/
func checkAndBuy(id int) {

	tLevel := 0
	if id >= collectorBuilding {
		tLevel, _ = gameStats.Buildings[id]
	}
	// next level should be bought or "level 1 ressource" clicked
	tLevel++

	cost := calcCost(id, tLevel)
	if cost == nil {
		checkError(fmt.Errorf("no cost for %+v", id))
		return
	}

	canAfford := true
	for res, amount := range cost {
		if amount > gameStats.Ressources[res] {
			canAfford = false
			break
		}
	}

	// cant do anything
	if !canAfford {
		return
	}

	// pay
	for res, amount := range cost {
		gameStats.Ressources[res] -= amount
	}

	// get
	if id >= collectorBuilding {
		gameStats.Buildings[id]++
		gameStats.BuildingsActive[id]++
		updateTickIncrease()
		updateButtons()
	} else {
		gameStats.Ressources[id]++
	}
	updateStats()
}

/*
sorts and stringify calculated costs
*/
func calcCostToString(building, targetlevel int) string {

	costs := calcCost(building, targetlevel)
	if costs == nil {
		return strInvalidCosts
	}

	ret := []string{}

	for cost := matter; cost <= structures; cost++ {
		symbol := costSymbol[cost]
		if val, ok := costs[cost]; ok {
			ret = append(ret, symbol+costSymbolAmountDivide+
				strconv.FormatFloat(val, costDisplayFormat, costDisplayPrec, 64))
		}
	}

	return strings.Join(ret, costStringsJoin)
}

// #############################################################################
// #							Gains on UPDATE
// #############################################################################

/*
calcUpdate recalculates the global update struct that holds all increasing or
decreasing values.

* base production, starts with 1, can be raised
* add %, all added and aplied to base production
* multiply, multiply through
*/
func updateTickIncrease() {
	// reset all
	incrDecrOnBigUpd = []*cost{}

	for buildID, amountActive := range gameStats.BuildingsActive {
		// TODO ship ressources will not work here
		if amountActive == 0 || buildID&shipBuildingMax > 0 {
			continue
		}

		ressBase := map[int]float64{}
		ressInc := map[int]float64{}
		ressMult := map[int]float64{}

		// which ressource
		ress := buildID - collectorBuilding
		ressBase[ress] = 1.0
		ressInc[ress] = 1.0
		ressMult[ress] = 1.0

		// cost
		if _, ok := baseCost[ress]; !ok {
			checkError(fmt.Errorf("no cost for ressource id: %+v", ress))
			continue
		}

		buildDur := 1.0

		for pay, amount := range baseCost[ress] {
			if pay == buildTime {
				buildDur = amount
				continue
			}
			ressBase[pay] = amount * -1
			ressInc[pay] = 1.0
			ressMult[pay] = 1.0
		}

		/*
			for _, sci := range gameStats.UnlockedScience{
			}
		*/
		newC := new(cost)
		newC.item = ress
		newC.ressAdd = make(map[int]float64)
		for r, base := range ressBase {
			base *= float64(amountActive) / (updatesPerSecond * buildDur)
			newC.ressAdd[r] = base * ressInc[r] * ressMult[r]
		}
		incrDecrOnBigUpd = append(incrDecrOnBigUpd, newC)
	}

	// raise limits of limiteded ressources
	for res := range resWithMax {
		if val, ok := gameStats.Buildings[res+shipBuildingMax]; ok {
			resWithMax[res] = val
		}
	}
}
