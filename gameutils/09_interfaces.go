package gameutils

import (
	eb "github.com/hajimehoshi/ebiten/v2"
)

/*
PositionAble is the corresponding interface to access PosObjects. Provides:
* Remove, if chunk is nil do not remove this tick. else remove from Chunk+Board
* Draw, provide a scale factor and a scrren do draw on
* SetBoardChunkAndNode, for communication with board
*/
type PositionAble interface {
	Remove() bool
	GetCoords() (float64, float64)
	GetNodeInfo() (*Chunk, *ChunkObjNode)
	Draw(float64, float64, float64, *eb.Image)
	SetBoardInfo(*Board, *Chunk, *ChunkObjNode)
	GetFactionAndType() (int, int)
}

/*
MoveAble provides the interface to all moving and colliding game objects
*/
type MoveAble interface {
	UpdatePosition()
	GetNodeInfo() (*Chunk, *ChunkObjNode)
}

/*
CollideAble provides the interface to check for collision and assign
Collision events by accessing CollidingObject
*/
type CollideAble interface {
	CheckCollide(*CollidingObject) bool
	GetCollidingObject() *CollidingObject
	GetNodeInfo() (*Chunk, *ChunkObjNode)
}

/*
ClickAble is the interface for game objects that can be clicked
*/
type ClickAble interface {
	CheckClick(*CollidingObject, eb.MouseButton) bool
}

// #############################################################################

/*
UpdateAble provides interface to anything that needs actual logic besides moving
and collision detection.
*/
type UpdateAble interface {
	Update()
}
