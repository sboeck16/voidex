package gameutils

import (
	"sync"

	eb "github.com/hajimehoshi/ebiten/v2"
)

var (
	// to register clicks we add a colliding object to board and make a
	// collision check
	click     = getClickObj()
	clickLock sync.RWMutex
)

// #############################################################################
// #							CollisionDetection
// #############################################################################

/*
ChecksPossibleCollisions for given object. Uses neigbouring chunks
*/
func (b *Board) CheckPossibleCollisions(co CollideAble) {

	// get chunk and coords
	coChunk, _ := co.GetNodeInfo()
	coChunk.lock.RLock()
	cx := coChunk.myBoardX
	cy := coChunk.myBoardY
	coChunk.lock.RUnlock()
	// read lock for board chunks
	b.chunkLock.RLock()
	defer b.chunkLock.RUnlock()
	// now iterate through all possible chunks and their nodes
	for cxi := -1; cxi < 2; cxi++ {
		for cyi := -1; cyi < 2; cyi++ {
			if subMap, ok1 := b.chunks[cxi+cx]; ok1 {
				if chunk, ok2 := subMap[cyi+cy]; ok2 {
					chunk.lock.RLock()
					node := chunk.objHead
					for node != nil {
						if node.Coll != nil {
							co.CheckCollide(node.Coll.GetCollidingObject())
						}
						node = node.next
					}
					chunk.lock.RUnlock()
				}
			}
		}
	}
}

/*
utility method to initialize click object
*/
func getClickObj() *CollidingObject {
	ret := new(CollidingObject)
	return ret
}

// #############################################################################
// #							Board searches
// #############################################################################

/*
GetNearesObject searches nearest object that matches type and faction from
starting coords x, y. Returns object. provide an integer how far in chunks
a target is searched.
*/
func (b *Board) GetNearestObject(x, y float64, faction, objType int,
	searchRange int) PositionAble {

	centerChunk := b.GetChunk(x, y)
	if ret := lookupnearestObj(
		x, y, faction, objType, []*Chunk{centerChunk}); ret != nil {
		return ret
	}
	dx := centerChunk.myBoardX
	dy := centerChunk.myBoardY
	for dist := 1; dist <= searchRange; dist++ {
		chunks := []*Chunk{}

		b.chunkLock.Lock()
		// horizontal rows and edges
		for _, sy := range []int{dy - dist, dy + dist} {
			for sx := dx - dist; sx <= dx+dist; sx++ {
				if cmap, ok := b.chunks[sx]; ok {
					if chu, ok := cmap[sy]; ok {
						chunks = append(chunks, chu)
					}
				}
			}
		}
		// vertical rows
		for _, sx := range []int{dx - dist, dx + dist} {
			for sy := dy - dist + 1; sy < dy+dist; sy++ {
				if cmap, ok := b.chunks[sx]; ok {
					if chu, ok := cmap[sy]; ok {
						chunks = append(chunks, chu)
					}
				}
			}
		}
		b.chunkLock.Unlock()

		ret := lookupnearestObj(x, y, faction, objType, chunks)
		if ret != nil {
			return ret
		}
	}

	return nil
}

/*
Utility method takes a slice of chunks and returns nearest positionable.
*/
func lookupnearestObj(x, y float64, faction, objType int,
	chunks []*Chunk) PositionAble {

	var ret PositionAble
	sqDist := 0.0
	for _, chunk := range chunks {
		chunk.lock.RLock()
		node := chunk.objHead
		for node != nil {
			if node.Pos != nil {
				fac, typ := node.Pos.GetFactionAndType()
				if (fac&faction > 0) && (typ&objType > 0) {
					nx, ny := node.Pos.GetCoords()
					nx -= x
					ny -= y
					dist := nx*nx + ny*ny
					if sqDist == 0 || dist < sqDist {
						ret = node.Pos
						sqDist = dist
					}
				}
			}
			node = node.next
		}
		chunk.lock.RUnlock()
	}
	return ret
}

// #############################################################################
// # 							Click
// #############################################################################

/*
BoardClicked is used when the board is clicked.
*/
func (b *Board) BoardClicked(x, y float64, button eb.MouseButton) {

	// clicklock allows us to use and place click object
	clickLock.Lock()
	defer clickLock.Unlock()

	// incoming x is game coords - board position in game -> adapt to view port
	x += b.viewX
	y += b.viewY
	// set click object
	click.x = x
	click.y = y
	// get chunk
	chunk := b.GetChunk(float64(x), float64(y))
	cx := chunk.myBoardX
	cy := chunk.myBoardY

	// iterate through chunk and 8 neighbours
	for cxi := -1; cxi < 2; cxi++ {
		for cyi := -1; cyi < 2; cyi++ {
			if subMap, ok1 := b.chunks[cxi+cx]; ok1 {
				if chunk, ok2 := subMap[cyi+cy]; ok2 {
					chunk.lock.RLock()
					node := chunk.objHead
					for node != nil {
						if node.Click != nil {
							if node.Click.CheckClick(click, button) {
								chunk.lock.RUnlock()
								return
							}
						}
						node = node.next
					}
					chunk.lock.RUnlock()
				}
			}
		}
	}

	// if we reached here no object was clicked so we run board click handle
	if b.ClickHandle != nil {
		b.ClickHandle(b, x, y, button)
	}
}
