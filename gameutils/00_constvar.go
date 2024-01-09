package gameutils

import (
	"voidex/graphics"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
)

const (
	// Factions
	// --------

	// OwnerNeutral does not interact
	OwnerNeutral = 1 << iota
	// OwnerPlayer is player faction
	OwnerPlayer
	// OwnerLost pirates, lost and random
	OwnerLost
	// OwnerRobots nano robot faction
	OnwerRobots
	// OwnerPlant void growth
	OwnerPlants

	// object types (not from interface but from game logic)
	// ------------
	ObjProjectile = 1 << iota
	ObjRocket
	ObjBeam
	ObjPD
	ObjFighter
	ObjCruiser
	ObjCapital
	ObjBehemoth

	// weapon types
	// ------------

	// WeaponBeam for instant beam attacks
	WeaponBeam = 1 << iota
	// WeaponProjectile projectile weapons
	WeaponProjectile
	// WeaponRocket for rockets
	WeaponRocket
	// WeaponExplosion is a one tick effeckt (followup funtion)

	// worker modes
	// ------------

	// WModeDelete delete flagged objects
	WModeDelete = iota
	// WModeMove update positions
	WModeMove
	// WModeCollide checks collisions
	WModeCollide
	// WModeUpdate trigger objects update methods
	WModeUpdate

	// collision modes
	// ---------------

	// CollModeIsInImageBorder if collision radius is set to this value,
	// image rectangle without rotation is used for collision detection (click)
	CollModeIsInImageBorder = -1.0

	// Weapon default values
	// ---------------------
	// Rocket
	wdRocketMass     = .1
	wdRocketCollSize = 3.0
	wdRocketMaxTicks = 300
	wdRocketSpeed    = 7
	wdRocketAccel    = 1
	//wdRocketScanRange   = 3 // TODO, weapon scan vs fo scan vs proj scan?
	wdRocketTargetTypes  = ObjFighter | ObjCruiser | ObjCapital | ObjBehemoth
	wdRocketGrace        = 15
	wdRocketChargeNeeded = 30
	wdRocketChargeMax    = 90
	wdRocketCoolDown     = 5
	wdRocketCorrF        = 0.9
	wdrocketMaxRange     = 400
	wdRocketDuration     = 200
	wdRocketScatter      = 1.0
)

var (
	// WorkMode list all modes which board worker uses, indices are their values
	WorkMode = []int{
		WModeDelete,
		WModeMove,
		WModeCollide,
		WModeUpdate,
	}

	// ColFric is a variable to describe how bouncy (1 bounce, 0 stick) objects
	// are
	ColFric = float64(1.0)
	// CorrFactor multpiplies the speed of a moving object to slow it down

	// BigUpdateEvery is used for expensive searches like target finding
	BigUpdateEvery = 6

	// to write errors to, better implementation MAYBE?
	logError = func(string) {}

	// FactionGraphics -> TODO move, better graphics library
	FactionGraphics = map[int]map[int]*eb.Image{
		OwnerPlayer: {
			ObjRocket:  graphics.GenRocket(3, 1, graphics.ColorWhite),
			ObjFighter: graphics.GenShip(20, 2, graphics.ColorWhite),
		},
		OwnerLost: {
			ObjRocket:  graphics.GenRocket(3, 1, graphics.ColorRed),
			ObjFighter: graphics.GenShip(20, 2, graphics.ColorRed),
		},
	}

	// ButtonBGColor is the button backgroundcolor
	ButtonBGColor = graphics.ColorDarkGray
	// ButtonTextCol is the text color in buttons
	ButtonTextCol = graphics.ColorWhite

	// TextFontMiddle holds ebiten example font in middle size
	TextFontMiddle font.Face
	// TextFontSmall holds ebiten example font in small size
	TextFontSmall font.Face
)

// MAYBE provide, make font modul or move to graphics?
func init() {

	// initialize font for buttons
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		panic(err)
	}
	TextFontMiddle, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err) // TODO error handling
	}
	TextFontSmall, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err) // TODO error handling
	}

}
