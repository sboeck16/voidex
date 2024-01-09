package gameutils

/*
Weapon descripes the a weapon
*/
type Weapon struct {
	// to acces owner object and reposition from its center
	myFightObj *FightingObject
	// in rotation coords
	ownerAngle  float64
	ownerRadius float64

	// type, behavior
	weaponType int
	targetType int

	fireFunction func(*Weapon)

	// damage values
	damageFirst    float64
	mass           float64 // mass of projected effect
	precision      float64
	collSize       float64
	hitMinSize     int
	impulseMinSize int

	// how often in fires
	chargesNeeded float64
	maxCharge     float64
	chargeState   float64 // must be increased on update
	cooldown      float64
	cooldownState float64 // can fire on reaching cooldown, resets on firing
	maxRange      float64

	// projectile stats
	projSpeed  float64
	accel      float64 // needed for homing
	corrFactor float64
	duration   int // how long the projectile lifes in ticks
}

/*
Charge charges the weapon.
*/
func (w *Weapon) Charge(charges float64) {
	w.chargeState += charges
	if w.chargeState > w.maxCharge {
		w.chargeState = w.maxCharge
	}
}

/*
checks if weapon is ready to fire and if so fire it
*/
func (w *Weapon) fire() {

	// tick passed increase cooldown
	w.cooldownState++

	// check charge and cooldown and target, check if in range
	if w.chargeState < w.chargesNeeded || w.cooldownState < w.cooldown ||
		w.myFightObj.enemyTarget == nil ||
		(w.myFightObj.targetDist > 0 && w.myFightObj.targetDist > w.maxRange) {
		return
	}

	// we fire now, reset cooldown reduce carge
	w.cooldownState = 0
	w.chargeState -= w.chargesNeeded

	if w.fireFunction != nil {
		w.fireFunction(w)
		return
	}

	// if we reach here we havent fired weapon yet, so we use generic
	// method. type will be interpreted there
	w.CreateDamagingObject(w.myFightObj.myBoard)
}

// #############################################################################
// #							DefaultWeapons
// #############################################################################

/*
GetDefaultRocketLauncher returns a rocket launcher with default values.
*/
func GetDefaultRocketLauncher(damage, precision float64) *Weapon {
	ret := &Weapon{
		weaponType:     WeaponRocket,
		targetType:     wdRocketTargetTypes,
		damageFirst:    damage,
		mass:           wdRocketMass,
		precision:      precision,
		collSize:       wdRocketCollSize,
		hitMinSize:     ObjPD,
		impulseMinSize: ObjCruiser,
		chargesNeeded:  wdRocketChargeNeeded,
		maxCharge:      wdRocketChargeMax,
		chargeState:    wdRocketChargeMax,
		cooldown:       wdRocketCoolDown,
		cooldownState:  wdRocketCoolDown,
		maxRange:       wdrocketMaxRange,
		projSpeed:      wdRocketSpeed,
		accel:          wdRocketAccel,
		corrFactor:     wdRocketCorrF,
		duration:       wdRocketDuration,
	}
	return ret
}
