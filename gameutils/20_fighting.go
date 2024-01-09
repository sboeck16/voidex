package gameutils

import (
	"math"
	"math/rand"
)

/*
FightingObject is a game object that takes part in the battle simulation.
*/
type FightingObject struct {
	ClickObject // inherits, move, collide, game object

	// keep track of tick to make expensive update
	tick int

	// keep track of damage health
	maxhealth, health float64

	// moving
	accel    float64
	speed    float64
	corr     float64
	avoidDir float64
	minDist  float64

	// defenses
	agility float64 // avoid

	// Weapons, attack mode of the ship
	Weapons []*Weapon

	// selected enemy
	enemyTarget   PositionAble
	scanRange     int
	targetFaction int
	targetType    int
	targetDist    float64

	// keep track of hits
	hitAvoidCounter int

	// idleTarget what to follow if no enemy is here
	idleTarget PositionAble
}

/*
Update method that searches for a target and fires the weapons. Resolves
impacts.
*/
func (fo *FightingObject) Update() {
	fo.ResolveImpacts(ObjFighter)
	fo.collideEvents = []*Collision{}

	if fo.tick%BigUpdateEvery == 0 {
		// MAYBE research target if target is to far away -> very big update?
		if fo.enemyTarget == nil || fo.enemyTarget.Remove() {
			fo.enemyTarget = fo.myBoard.GetNearestObject(
				fo.x, fo.y, fo.targetFaction, fo.targetType, fo.scanRange)
		}
		if fo.enemyTarget != nil {
			fo.FighterMove()
		} else {
			fo.targetDist = 0
			// TODO move towards carrier -> needs formation move
		}
	}
	fo.tick++

	// fire weapons
	for _, weapon := range fo.Weapons {
		// charge them with ticks, for now...
		weapon.Charge(1.0)
		weapon.fire()
	}
}

/*
TakeDamage resolves a collision applying damage if hit and returns whether a
hit occurred or not.
*/
func (fo *FightingObject) TakeDamage(damage, precision float64) bool {
	// misses
	if rand.Float64() > precision/(precision+fo.agility) {
		fo.hitAvoidCounter++
		return false
	}
	// reduce by armor or resistance TODO

	fo.health -= damage
	fo.removeMe = fo.health <= 0
	return true
}

/*
AddWeapon adds a weapon to fighting object.
*/
func (fo *FightingObject) AddWeapon(weapon *Weapon) {
	fo.Weapons = append(fo.Weapons, weapon)
	weapon.myFightObj = fo
}

// #############################################################################
// #							MOVE -> new class or function pointer for move?
// #############################################################################

/*
RotMove provides an update function that uses target coordinates an acceleration
a maximum speed and a correction factor to recalculate the dx and dy values.
*/
func (fo *FightingObject) FighterMove() {
	// reduce drift
	fo.dx = fo.dx * fo.corr
	fo.dy = fo.dy * fo.corr

	// diff vektor
	tx, ty := fo.enemyTarget.GetCoords()
	cx := tx - fo.x
	cy := ty - fo.y

	clen := math.Sqrt(cx*cx + cy*cy)
	normalx := cx / clen
	normaly := cy / clen

	if clen < fo.minDist {
		fo.avoidDir += (0.5 - rand.Float64()) * (1 / (1 + fo.avoidDir))

		fo.dx -= (normalx + math.Sin(fo.avoidDir)) * fo.accel
		fo.dy -= (normaly + math.Cos(fo.avoidDir)) * fo.accel
	} else {
		fo.dx += normalx * fo.accel
		fo.dy += normaly * fo.accel
	}

	speed := math.Sqrt(fo.dx*fo.dx + fo.dy*fo.dy)
	if fo.speed > 0 && speed > fo.speed {
		fo.dx = fo.speed * normalx
		fo.dy = fo.speed * normaly
	}

	fo.rotation = math.Pi - math.Atan2(cx, cy)
}
