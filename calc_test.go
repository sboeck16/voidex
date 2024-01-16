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

}
