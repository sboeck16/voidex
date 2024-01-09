package gameutils

import (
	eb "github.com/hajimehoshi/ebiten/v2"
)

/*
PosObject holds any game object that is drawable and has a position on the
Board.
Holds some information about the board it belongs to and some information about
type and faction.
It should be inherited by almost any other game objects.
*/
type PosObject struct {

	// faction, bitmasked?
	faction int
	// type, bitmasked?
	objType int

	// remove flage
	removeMe bool

	// where are we placed
	myChunk    *Chunk
	myBoard    *Board
	myChunkObj *ChunkObjNode

	// postion
	x        float64
	y        float64
	rotation float64

	// base image
	img *eb.Image
	// draw option
	op           eb.DrawImageOptions
	drawx, drawy float64
}

/*
Remove returns whether this obect should be removed.
*/
func (p *PosObject) Remove() bool {
	return p.removeMe
}

/*
GetCoords returns the x, y coordinates.
*/
func (p *PosObject) GetCoords() (float64, float64) {
	return p.x, p.y
}

/*
GetNodeInfo returns chunk where this object is placed as well as it chunk object
ndoe
*/
func (p *PosObject) GetNodeInfo() (*Chunk, *ChunkObjNode) {
	return p.myChunk, p.myChunkObj
}

/*
GetFactionAndType returns faction and type of gameobject
*/
func (p *PosObject) GetFactionAndType() (int, int) {
	return p.faction, p.objType
}

/*
Draw brings the object on screen
*/
func (p *PosObject) Draw(scale, offx, offy float64, screen *eb.Image) {
	if p.img == nil {
		return
	}

	p.op.GeoM.Reset()
	p.op.GeoM.Translate(p.drawx, p.drawy)
	p.op.GeoM.Rotate(p.rotation)
	p.op.GeoM.Translate(p.x+offx, p.y+offy)

	screen.DrawImage(p.img, &p.op)
}

/*
SetImg sets an iage for game object and sets its offset to image center
*/
func (p *PosObject) SetImg(img *eb.Image) {
	p.drawx = 0 - float64(img.Bounds().Dx())/2
	p.drawy = 0 - float64(img.Bounds().Dy())/2
	p.img = img
}

/*
SetBoardInfo sets where game object belongs to.
*/
func (p *PosObject) SetBoardInfo(
	b *Board, c *Chunk, node *ChunkObjNode) {
	p.myBoard = b
	p.myChunk = c
	p.myChunkObj = node
}
