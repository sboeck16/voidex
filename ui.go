package main

import (
	"strconv"

	"voidex/gameutils"

	eb "github.com/hajimehoshi/ebiten/v2"
)

var (
	firstAddBtn = 0
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
	for _, buildID := range allBuildingsCollectable {
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
			subActive, buildID)
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
			addActive, buildID)
		bttnY += actionButtonNextRowSpace + actionButtonHeight
	}
	bttnX = actionButtonSpace + actionButtonWidth + actionButtonAddSubWidth +
		actionButtonAddSubWidth + actionButtonColumnSpace
	bttnY = actionButtonNextRowSpace
	for _, buildID := range allBuildingsShips {
		// sub
		actionButtons[buildID+collectorSub] = newActionButton(
			float64(bttnX), float64(bttnY),
			actionButtonAddSubWidth, actionButtonHeight,
			subActive, buildID)
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
			addActive, buildID)
		bttnY += actionButtonNextRowSpace + actionButtonHeight
	}

	// action buttons like battle, story or science
	actionButtons[buttonBattle] = newMainButton(battleButtonX, battleButtonY,
		battleButtonW, battleButtonH, battleButtonText, initiateBattle)

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
		if buildID&modeButton > 0 {
			continue
		}
		if buildID < collectorBuilding {
			msg = actionStrCollect + sp + toStringRessource[buildID]
			msg += nl + calcCostToString(buildID, 1)
		} else if buildID&collectorAdd > 0 {
			msg = actionStrAdd
		} else if buildID&collectorSub > 0 {
			msg = actionStrSub
		} else {
			lvl, _ := gameStats.Buildings[buildID]
			// MAYBE refactor, string concat or buffer write?
			msg = toStringBuildings[buildID]
			msg += actionButtonActive1
			msg += strconv.Itoa(gameStats.BuildingsActive[buildID])
			msg += actionButtonActive2
			msg += strconv.Itoa(gameStats.Buildings[buildID])
			msg += actionButtonActive3 + nl
			msg += calcCostToString(buildID, lvl+1)
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
			if val, ok := resWithMax[resID]; ok {
				msg += statDisplayMaxResDivide + strconv.Itoa(val)
			}
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
	ret.SetColorAndFont(
		actionDisplayTextCol, actionDisplayBGCol, actionDisplayFont)

	return ret
}

func newMainButton(x, y float64, w, h int, display string,
	act func(_ *gameutils.ClickObject, _ eb.MouseButton)) *gameutils.Button {

	ret := gameutils.NewButton(x, y, w, h, act)
	ret.SetColorAndFont(
		battleButtonTextCol, battleButtonBGCol, battleButtonFont)
	ret.SetText(5, 15, display)
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
		return func(_ *gameutils.ClickObject, _ eb.MouseButton) { checkAndBuy(with) }
	case addActive:
		return func(_ *gameutils.ClickObject, _ eb.MouseButton) { activateBuilding(with) }
	case subActive:
		return func(_ *gameutils.ClickObject, _ eb.MouseButton) { deactivateBuilding(with) }
	}
	return nil
}

func activateBuilding(building int) {

	if gameStats.BuildingsActive[building] < gameStats.Buildings[building] {
		gameStats.BuildingsActive[building]++
		updateButtons()
		updateTickIncrease()
	}
}

func deactivateBuilding(building int) {

	if gameStats.BuildingsActive[building] > 0 {
		gameStats.BuildingsActive[building]--
		updateButtons()
		updateTickIncrease()
	}
}

func initiateBattle(_ *gameutils.ClickObject, _ eb.MouseButton) {
	deb("BATTLE")
}
