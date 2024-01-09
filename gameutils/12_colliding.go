package gameutils

import (
	"math"
	"math/rand"
	"sync"
)

/*
CollidingObject describes an object with a collision radius and the ability
to store Collision events to be processed by update function
*/
type CollidingObject struct {
	MovingObject

	// mass
	mass float64

	// collision
	collSize float64
	// stored collsion type (mass damage etc)
	myCollObj *Collision
	// events from the object that collided with this
	collideEvents []*Collision
	collLock      sync.RWMutex
}

/*
Collision holds fields to resove a collision
*/
type Collision struct {
	// position and vector
	x        float64
	y        float64
	dx       float64
	dy       float64
	mass     float64
	collSize float64
	faction  int
	objType  int
	foRef    *FightingObject // for damaging collision
}

/*
GetCollidingObject is an interface method to acces CollidingObject for collision
detection (direct access makes it easier).
*/
func (co *CollidingObject) GetCollidingObject() *CollidingObject {
	return co
}

/*
CheckCollide is triggerd external. A bool is returned if needed and a collsion
is added to o.
*/
func (co *CollidingObject) CheckCollide(c *CollidingObject) bool {

	// MAYBE implement double hit with a hash set? could be more effective.

	// ignore self collision...
	if co == c {
		return false
	}

	collided := false

	if co.collSize > 0 {
		squareDist := (c.x-co.x)*(c.x-co.x) + (c.y-co.y)*(c.y-co.y)
		collided = squareDist < (c.collSize+co.collSize)*(c.collSize+co.collSize)
	} else if co.collSize == CollModeIsInImageBorder {

		bounds := co.img.Bounds()
		ix := float64(bounds.Max.X)
		iy := float64(bounds.Max.Y)

		collided = co.x+co.drawx <= c.x && co.x+co.drawx+ix >= c.x &&
			co.y+co.drawy <= c.y && co.y+co.drawy+iy >= c.y
	}

	if collided && co.myCollObj != nil {
		// update own coll obj
		co.myCollObj.x = co.x
		co.myCollObj.y = co.y
		co.myCollObj.dx = co.dx
		co.myCollObj.dy = co.dy
		// add now, things like mass or damage are created on object
		// creation
		c.AddCollision(co.myCollObj)
		// return value
	}

	return collided
}

/*
AddCollision adds an collsion object to the object.
*/
func (co *CollidingObject) AddCollision(coll *Collision) {

	co.collLock.Lock()
	co.collideEvents = append(co.collideEvents, coll)
	co.collLock.Unlock()
}

// #############################################################################
// #							Resolve Impacts
// #############################################################################

/*
ResolveImpacts resolvethe impuls or moving part of a collision. Must be called
in update and cleared there!
*/
func (co *CollidingObject) ResolveImpacts(minSize int) {

	// nothing to do if wew have no mass
	if co.mass == 0 {
		return
	}

	// iterate through
	for _, col := range co.collideEvents {
		// mass zero impacts have no result
		if col.mass == 0 || col.objType < minSize {
			continue
		}
		co.ResolveImpulse(col)
	}
}

/*
ResolveImpulse resolves the physical impuls of a Collision event
*/

func (co *CollidingObject) ResolveImpulse(col *Collision) {
	cx := co.x - col.x
	cy := co.y - col.y
	clen := math.Sqrt(cx*cx + cy*cy)

	if clen == 0 {
		clen = 1
		rot := getNextGrad()
		cx += math.Sin(rot)
		cy += math.Cos(rot)
	}

	normalx := cx / clen
	normaly := cy / clen

	ff := normalx*(co.dx-col.dx) + normaly*(co.dy-col.dy)
	f := (1 + ColFric) * ff / (1/co.mass + 1/col.mass)

	co.dx -= f * normalx / co.mass
	co.dy -= f * normaly / co.mass

	collRad := co.collSize + col.collSize
	cr := (collRad - clen) * col.mass / (co.mass + col.mass)

	co.dx += normalx * cr / collRad
	co.dy += normaly * cr / collRad

	// move out of the way
	co.x += normalx * cr
	co.y += normaly * cr
}

// #############################################################################
// #							utility
// #############################################################################

/*
Returns a rotation and increases it for next.
*/
func getNextGrad() float64 {
	return rand.Float64() * math.Pi * 2
}
