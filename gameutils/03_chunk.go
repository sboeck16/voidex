package gameutils

import (
	"sync"
)

/*
Chunk is a single field of the Board and holds a list of GameObjects. Holds
its coordinates to pass them to GameObjects that moves around
*/
type Chunk struct {
	lock sync.RWMutex

	// coords
	x float64
	y float64

	// chunk info
	size     float64
	parent   *Board
	myBoardX int
	myBoardY int

	// list for objects
	objHead *ChunkObjNode
	objTail *ChunkObjNode
}

/*
ChunkObjNode holds moveable interfaces for chunks for collision and position
or enemy detection
*/
type ChunkObjNode struct {
	Pos   PositionAble
	Coll  CollideAble
	Click ClickAble

	prev *ChunkObjNode
	next *ChunkObjNode
}

/*
AddObj adds an objNode to chunk.
*/
func (c *Chunk) AddObj(node *ChunkObjNode) {

	// syncing
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.objHead == nil {
		c.objHead = node
		c.objTail = node
	} else {
		c.objTail.next = node
		node.prev = c.objTail
		c.objTail = node
	}
}

/*
DeleteObj deletes an obj node
*/
func (c *Chunk) DeleteObj(node *ChunkObjNode) {

	// syncing
	c.lock.Lock()
	defer c.lock.Unlock()

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		// head node deleted
		c.objHead = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		// tail node deleted
		c.objTail = node.prev
	}
	// chunk object should hold no refs to this list so delete them
	// to avoid mixing chunk lists
	node.prev = nil
	node.next = nil
}
