package gameutils

import (
	"math"
	"math/rand"
)

/*
DamagingObject represent an object that can damage fighting objects.
*/
type DamagingObject struct {
	CollidingObject
	tick           int
	maxTicks       int
	speed          float64
	accel          float64
	corrFactor     float64
	targetFaction  int
	targetType     int
	scanRange      int
	damage         float64
	precision      float64
	target         PositionAble
	minSizeImpulse int
	minSizeCollide int
	graceTicks     int
	savedSize      float64
}

/*
CreateDamagingObj create a game obj that can hit things with health from weapon.
If a board is provided it will be registered there.
*/
func (w *Weapon) CreateDamagingObject(registerOn *Board) *DamagingObject {

	dObj := &DamagingObject{
		maxTicks:       w.duration,
		speed:          w.projSpeed,
		targetFaction:  w.myFightObj.targetFaction,
		targetType:     w.targetType,
		scanRange:      w.myFightObj.scanRange,
		damage:         w.damageFirst,
		precision:      w.precision,
		target:         w.myFightObj.enemyTarget,
		minSizeImpulse: w.impulseMinSize,
		minSizeCollide: w.hitMinSize,
		accel:          w.accel,
		corrFactor:     w.corrFactor,
		savedSize:      w.collSize,
	}
	dObj.faction = w.myFightObj.faction
	dObj.collSize = w.collSize
	dObj.mass = w.mass
	dObj.myCollObj = &Collision{
		mass: dObj.mass,
	}
	dObj.x = w.myFightObj.x +
		math.Sin(w.ownerAngle+w.myFightObj.rotation)*w.ownerRadius
	dObj.y = w.myFightObj.y +
		math.Cos(w.ownerAngle+w.myFightObj.rotation)*w.ownerRadius

	// interpret weapontype
	switch w.weaponType {
	case WeaponRocket:
		dObj.SetImg(FactionGraphics[w.myFightObj.faction][ObjRocket])
		dObj.accel = w.accel
		rot := rand.Float64() * math.Pi * 2
		dObj.dx = w.myFightObj.dx + wdRocketScatter*math.Sin(rot)
		dObj.dy = w.myFightObj.dy + wdRocketScatter*math.Cos(rot)
	case WeaponProjectile:
	case WeaponBeam:
	}

	if registerOn != nil {
		registerOn.RegisterObj(dObj, dObj, dObj, dObj, nil)
	}

	return dObj
}

/*
Update method to resolve hitsand apply damage
*/
func (do *DamagingObject) Update() {
	do.checkHits()

	if do.accel == 0 {
		return
	}
	// accel > 0 homing projectile
	if do.tick%BigUpdateEvery == 0 {
		if do.target == nil || do.target.Remove() {
			do.target = do.myBoard.GetNearestObject(
				do.x, do.y, do.targetFaction, do.targetType, do.scanRange)
		}
		if do.target != nil {
			tx, ty := do.target.GetCoords()
			do.RotMove(tx, ty, do.accel, do.speed, do.corrFactor)
		}
	}
	do.tick++
	// remove after duration
	do.removeMe = do.removeMe || do.tick > do.maxTicks
}

func (do *DamagingObject) checkHits() {

	// grace period after pierce or misses
	if do.graceTicks > 0 {
		do.graceTicks--
		return
	}
	// reenable collision
	do.collSize = do.savedSize

	// check collisions
	for _, coll := range do.collideEvents {
		if coll.objType < do.minSizeCollide {
			continue
		}
		if coll.objType >= do.minSizeImpulse {
			do.ResolveImpulse(coll)
		}
		if coll.faction == do.faction {
			continue
		}
		// we hit something not from our faction
		if coll.foRef == nil {
			continue
		}
		hit := coll.foRef.TakeDamage(do.damage, do.precision)
		do.removeMe = hit
		if !hit && do.accel > 0 {
			do.graceTicks = wdRocketGrace
			do.collSize = 0
		}
	}
	// clear collisions
	do.collideEvents = []*Collision{}
}
