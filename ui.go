package main

import (
	"strconv"

	"voidex/gameutils"

	eb "github.com/hajimehoshi/ebiten/v2"
)

var (
	firstAddBtn = true
	// boards are located in main.go

	// holds the stat display buttons
	statDisplayButtons = map[int]*gameutils.Button{}
	activeStatDisplay  = map[int]bool{}

	// action buttons
	actionButtons        = map[int]*gameutils.Button{}
	displayActionButtons = map[int]bool{}
	reachedButtonGoal    = -1
)

func init() {
	// prepare all buttons
	nextIMult := statDisplayWidth + statDisplaySpace
	nextY := statDisplayHeight + statDisplayNextRowSpace
	for i := 0; i <= structures; i++ {
		statDisplayButtons[i] = newDisplayButton(float64(nextIMult*i), 0)
	}
	for i := science; i < maxResources; i++ {
		statDisplayButtons[i] = newDisplayButton(
			float64(nextIMult*(i-science)), float64(nextY))
	}

	// prepare all action buttons
	// all resources
	bttnX := actionButtonSpace
	bttnY := actionButtonNextRowSpace
	for _, buildID := range allBuildings {
		if buildID&collectorBuilding > 0 {
			// add clicker
			actionButtons[buildID-collectorBuilding] = newActionButton(
				float64(bttnX+actionButtonAddSubWidth), float64(bttnY),
				actionButtonWidth, actionButtonHeight,
				buy, buildID-collectorBuilding)
			bttnY += actionButtonNextRowSpace + actionButtonHeight
			// sub
			actionButtons[buildID+collectorSub] = newActionButton(
				float64(bttnX), float64(bttnY),
				actionButtonAddSubWidth, actionButtonHeight,
				subActive, buildID+collectorSub)
			// building
			actionButtons[buildID] = newActionButton(
				float64(bttnX+actionButtonAddSubWidth), float64(bttnY),
				actionButtonWidth, actionButtonHeight,
				buy, buildID)
			// add
			actionButtons[buildID+collectorAdd] = newActionButton(
				float64(bttnX+actionButtonAddSubWidth+actionButtonWidth),
				float64(bttnY),
				actionButtonAddSubWidth, actionButtonHeight,
				addActive, buildID+collectorAdd)
			bttnY += actionButtonNextRowSpace + actionButtonHeight
		}
	}

}

/*
updates the dispay of the buttons
*/
func updateButtons() {
	// check if a goal has been reached
	if reachedButtonGoal+1 < len(goalsForButtons) {
		if checkGoal(reachedButtonGoal + 1) {
			bttnIDs := goalsForButtons[reachedButtonGoal+1].unlocksBntID
			for _, bttnID := range bttnIDs {
				displayActionButtons[bttnID] = true
				bttn := actionButtons[bttnID]
				boardButtons.RegisterObj(bttn, nil, nil, nil, bttn)
			}
			reachedButtonGoal++
			if gameStats.buttonLevel < reachedButtonGoal {
				gameStats.buttonLevel = reachedButtonGoal
			}
		}
	}

	// update display
	for buildID, btn := range actionButtons {
		msg := ""
		if buildID < collectorBuilding {
			msg = "collect " + toStringRessource[buildID]
		} else if buildID&collectorAdd > 0 {
			msg = "+"
		} else if buildID&collectorSub > 0 {
			msg = "-"
		} else {
			msg = toStringBuildings[buildID]
		}
		btn.SetText(5, 15, msg)
	}
}

/*
updates the numerical display of ressources
*/
func updateDisplay() {

	for resID := 0; resID < maxResources; resID++ {
		if activeStatDisplay[resID] {
			msg := statDisplayText[resID] + statDisplayDivide
			msg += strconv.FormatFloat(gameStats.Ressources[resID],
				statDisplayFormat, statDisplayPrec, 64)
			statDisplayButtons[resID].SetText(5, 15, msg)
			continue
		}
		if val, ok := gameStats.Ressources[resID]; ok && val > 0 {
			activeStatDisplay[resID] = true
			nb := statDisplayButtons[resID]
			boardDisplayStat.RegisterObj(nb, nil, nil, nil, nil)
		}
	}
}

// #############################################################################
// #							Util
// #############################################################################

func newDisplayButton(x, y float64) *gameutils.Button {
	ret := gameutils.NewButton(x, y, statDisplayWidth, statDisplayHeight, nil)
	ret.SetColorAndFont(statDisplayTextCol, statDisplayBGCol, statDisplayFont)
	return ret
}

func newActionButton(
	x, y float64, w, h, action, with int) *gameutils.Button {
	ret := gameutils.NewButton(x, y, w, h, actionButtonHandle(action, with))
	ret.SetColorAndFont(statDisplayTextCol, statDisplayBGCol, statDisplayFont)

	return ret

}

/*
Returns buttons that can be displayed if goal is reached. utility function,
btnlvl is not checked for out of bound in goals array!
*/
func checkGoal(btnlvl int) bool {
	if btnlvl < gameStats.buttonLevel {
		return true
	}
	for res, valNeed := range goalsForButtons[btnlvl].ressourcesNeeded {
		if valHave, ok := gameStats.Ressources[res]; !ok || valNeed > valHave {
			return false
		}
	}
	return true
}

// #############################################################################
// #							Action
// #############################################################################

func actionButtonHandle(action, with int) func(*gameutils.ClickObject, eb.MouseButton) {
	switch action {
	case buy:
		return func(_ *gameutils.ClickObject, _ eb.MouseButton) { actionBuy(with) }
	case addActive:
		return func(_ *gameutils.ClickObject, _ eb.MouseButton) { actionActivateBuilding(with) }
	case subActive:
		return func(_ *gameutils.ClickObject, _ eb.MouseButton) { actionActivateBuilding(with) }
	}
	return nil
}

func actionBuy(building int) {

	deb("clicked B", building)
	updateButtons()
}

func actionActivateBuilding(building int) {

	deb("clicked +", building)
	updateButtons()
}

func deactivateBuilding(building int) {

	deb("clicked -", building)
	updateButtons()
}
