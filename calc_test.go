package main

import (
	"testing"
)

var (
	csv         = true
	calclvlsFor = []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		15, 20, 25, 30, 40, 50,
		75, 100,
	}
)

func TestCalc(t *testing.T) {
	/*
	   csv := ""
	   for buildID := 0; buildID < maxBuildings; buildID++ {

	   		head := []string{toStringBuildings[buildID]}
	   		cont := map[string][]string{}
	   		for _, lvl := range calclvlsFor {
	   			head = append(head, strconv.Itoa(lvl))
	   			costM := calcCost(buildID, lvl)
	   			for resID := 0; resID < maxResources; resID++ {

	   				if val, ok := costM[resID]; ok {
	   					cont[toStringRessource[resID]] = append(
	   						cont[toStringRessource[resID]],
	   						strconv.FormatFloat(val, 'f', 0, 64))
	   				}
	   			}

	   		}
	   		csv += strings.Join(head, ",") + "\n"
	   		for resID := 0; resID < maxResources; resID++ {
	   			if val, ok := cont[toStringRessource[resID]]; ok {
	   				csv += toStringRessource[resID] + "," + strings.Join(val, ",") + "\n"
	   			}
	   		}
	   	}

	   deb(csv)
	*/
}

func TestInc(t *testing.T) {
	gameStats = NewGameStats()
	for i := 0; i < maxResources; i++ {
		gameStats.Ressources[i] = 10
	}

	c := new(cost)
	c.pay = map[int]float64{matter: .5}
	c.get = 0.5
	c.item = fabric
	incrDecrOnBigUpd.cost = []*cost{c}
	incrDecrOnBigUpd.add[rareMatter] = 1

	for i := 0; i < 30; i++ {
		updateStats()
	}
	if gameStats.Ressources[matter] != 0.0 ||
		gameStats.Ressources[fabric] != 20.0 ||
		gameStats.Ressources[rareMatter] != 40.0 {
		deb(gameStats.Ressources)
		t.Error("update calc is wrong")
	}

}
