package gameutils

import (
	eb "github.com/hajimehoshi/ebiten/v2"
)

/*
ClickObject is a simple clickable game object
*/
type ClickObject struct {
	CollidingObject
	// function that is called on click
	clickHandle func(*ClickObject, eb.MouseButton)
}

/*
CheckClick checks if clickable object is clicked, click should be provided as
a colliding object
*/
func (c *ClickObject) CheckClick(
	click *CollidingObject, button eb.MouseButton) bool {

	if c.clickHandle != nil && c.CheckCollide(click) {
		c.clickHandle(c, button)
		return true
	}

	return false
}
