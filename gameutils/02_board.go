package gameutils

import (
	"sync"

	eb "github.com/hajimehoshi/ebiten/v2"
)

const (
	// BoardBorderNorth is used when object y < boards minY
	BoardBorderNorth = 1
	// BoardBorderSouth is used when object y > boards minY
	BoardBorderSouth = 2
	// BoardBorderWest is used when object x < boards minX
	BoardBorderWest = 4
	// BoardBorderEast is used when object x > boards minX
	BoardBorderEast = 8
)

var (
	// these are package/global variables used as initial values for creating
	// a new board. Maybe setter and getter are needed

	// BoardMinX border value
	BoardMinX = float64(-1000000)
	// BoardMinX border value
	BoardMinY = float64(-1000000)
	// BoardMinX border value
	BoardMaxX = float64(1000000)
	// BoardMinX border value
	BoardMaxY = float64(1000000)

	// ChunkSize is the default ChunkSize
	ChunkSize = 100
	// ChunkWorker is the default amount of update routines for board worker
	ChunkWorker = 4
	// BorderViolateFunction is the default function called upon MovingObject
	// if it crosses a board border.
	BorderViolateFunction func(*MovingObject, int)

	// DefaultScale is the initial scale
	DefaultScale = 1.0
)

/*
Board holds all GameObjects and provides method for collision detection, etc.
Holds absolute
*/
type Board struct {

	// Enabled exported flag if board is active. Not mutex protected
	Enabled bool
	// DrawMe is a flag
	DrawMe bool

	// position in game
	GameX, GameY float64

	// chunk handling
	chunkSize float64
	chunks    map[int]map[int]*Chunk
	chunkLock sync.RWMutex

	// coords of view port
	viewX, viewY float64
	scale        float64

	// size of the visible board
	height, width int

	// borders of the allowed space
	minX, maxX float64
	minY, maxY float64

	// border violation
	borderViolateFunc func(*MovingObject, int)

	// set the maximum amount of chunk workers
	// workers      []*boardWorker
	workerAmount    int
	worker          []*boardWorker
	workerObjAmount []int
	workerLock      sync.RWMutex

	workWait sync.WaitGroup

	// if clicked into but no object was clicked, maybe getter/setter?
	ClickHandle func(*Board, float64, float64, eb.MouseButton)
}

/*
ObjNode holds interface for the three interfaces, PositionAble, MoveAble and
UpdateAble. All three could point to the same object that provides all three
interfaces (e.g. through inheritance). The node points to previous and next
to be implemented as a double linked list
*/
type ObjNode struct {
	Pos    PositionAble
	Move   MoveAble
	Coll   CollideAble
	Update UpdateAble

	prev *ObjNode
	next *ObjNode
}

/*
NewBoard initialise a new gaming Board and returns it. Provide width and height
for the view window.
*/
func NewBoard(width, height int) *Board {
	ret := new(Board)

	// enable board on default
	ret.Enabled = true
	ret.DrawMe = true

	ret.chunkSize = float64(ChunkSize)
	ret.chunks = make(map[int]map[int]*Chunk)

	ret.width = width
	ret.height = height

	ret.scale = DefaultScale

	ret.borderViolateFunc = BorderViolateFunction
	ret.maxX = BoardMaxX
	ret.maxY = BoardMaxY
	ret.minX = BoardMinX
	ret.minY = BoardMinY

	ret.workerAmount = ChunkWorker
	ret.worker = make([]*boardWorker, ChunkWorker)
	ret.workerObjAmount = make([]int, ChunkWorker)
	for i := 0; i < ChunkWorker; i++ {
		newWorker := new(boardWorker)
		newWorker.workerID = i
		newWorker.myBoard = ret
		newWorker.activate = make(chan int, 1)
		newWorker.work()
		ret.worker[i] = newWorker
	}

	return ret
}

// #############################################################################
// # 							Manage Objects
// #############################################################################

/*
RegisterObj adds an game object. Parameters could be the same refs if it's
drawable and movable for example. Will try to set chunk, board and node to
object via Drawavle interface
*/
func (b *Board) RegisterObj(
	p PositionAble,
	m MoveAble,
	c CollideAble,
	u UpdateAble,
	cl ClickAble) *ObjNode {

	if p == nil {
		logError("try to add unpositional node added to board (skipped)!")
		return nil
	}
	newNode := new(ObjNode)
	newNode.Pos = p
	newNode.Move = m
	newNode.Update = u
	newNode.Coll = c

	// find the best worker
	wNr := b.getminObjWorkerInd()
	b.worker[wNr].AddObj(newNode)
	b.wAmountInc(wNr)

	chunkNode := new(ChunkObjNode)
	chunkNode.Pos = p
	chunkNode.Coll = c
	chunkNode.Click = cl

	x, y := p.GetCoords()
	_, chunk := b.AddObjToChunk(x, y, chunkNode)

	p.SetBoardInfo(b, chunk, chunkNode)

	return newNode
}

/*
AddObject takes object coordinates and its objNode to place it on the
Board. If chunk is not initialized on target coordinates it will be created if
possible. If Board borders are reached an int greater than 0 is returned. Used
chunk is returned.
*/
func (b *Board) AddObjToChunk(x, y float64, node *ChunkObjNode) (int, *Chunk) {

	// border check
	borderViolation := 0
	if x < b.minX {
		borderViolation = borderViolation | BoardBorderWest
		x = b.minX
	}
	if x > b.maxX {
		borderViolation = borderViolation | BoardBorderEast
		x = b.maxX
	}
	if y < b.minY {
		borderViolation = borderViolation | BoardBorderNorth
		y = b.minY
	}
	if y > b.maxY {
		borderViolation = borderViolation | BoardBorderSouth
		y = b.maxY
	}

	// get chunk or generate it
	chunk := b.GetChunk(x, y)

	// add objNode
	chunk.AddObj(node)

	return borderViolation, chunk
}

/*
GetChunk returns chunk identified by coords in it. Chunk will be created if
necessary. Border checks should be done first.
*/
func (b *Board) GetChunk(x, y float64) *Chunk {

	// syncing access
	b.chunkLock.Lock()
	defer b.chunkLock.Unlock()

	chunkX := int(x / b.chunkSize)
	chunkY := int(y / b.chunkSize)
	if _, ok := b.chunks[chunkX]; !ok {
		b.chunks[chunkX] = make(map[int]*Chunk)
	}
	if _, ok := b.chunks[chunkX][chunkY]; !ok {
		newChunk := new(Chunk)
		newChunk.myBoardX = chunkX
		newChunk.myBoardY = chunkY
		newChunk.size = b.chunkSize
		newChunk.x = float64(chunkX) * b.chunkSize
		newChunk.y = float64(chunkY) * b.chunkSize
		newChunk.parent = b

		b.chunks[chunkX][chunkY] = newChunk
	}

	return b.chunks[chunkX][chunkY]
}

/*
LeftChunk is called when a MoveAble left its chunk. The new x,y game coords
must be provided as well as the chunk that has been left. The item that left
is identified by the objNode that is initial provided upon registering.

m ist still locked and can be accessed directly.

Maybe rework with more parameters, etc. but would it be better or faster? Maybe
it needs to be done if m can't be used here
*/
func (b *Board) LeftChunk(m *MovingObject) {
	// delete
	m.myChunk.DeleteObj(m.myChunkObj)
	borderViolation, chunk := b.AddObjToChunk(m.x, m.y, m.myChunkObj)

	// set chunk and borders
	m.myChunk = chunk
	m.chunkMinX = chunk.x
	m.chunkMaxX = chunk.x + chunk.size
	m.chunkMinY = chunk.y
	m.chunkMaxY = chunk.y + chunk.size

	// border violation function (reflect, teleport, etc)
	if borderViolation > 0 && b.borderViolateFunc != nil {
		b.borderViolateFunc(m, borderViolation)
	}
}

// #############################################################################
// # 							Manage Worker
// #############################################################################

/*
utility function that scans for the worker with the least amount of objects, for
a pseudo load balancing
*/
func (b *Board) getminObjWorkerInd() int {
	b.workerLock.Lock()
	defer b.workerLock.Unlock()
	minInd := 0
	minAmount := 0
	for ind, amountOfObjs := range b.workerObjAmount {
		if minAmount > amountOfObjs {
			minInd = ind
			minAmount = amountOfObjs
		}
	}
	return minInd
}

/*
increases ovject counter for a given worker
*/
func (b *Board) wAmountInc(wNr int) {
	b.workerLock.Lock()
	defer b.workerLock.Unlock()

	b.workerObjAmount[wNr]++
}

/*
decreases object counter for a given worker
*/

func (b *Board) wAmountDec(wNr int) {
	b.workerLock.Lock()
	defer b.workerLock.Unlock()

	b.workerObjAmount[wNr]--
}
