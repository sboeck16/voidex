package gameutils

import (
	"sync"
	"time"
)

/*
boardWorker holds data and go routine
*/
type boardWorker struct {
	myBoard  *Board
	activate chan int
	workerID int

	// game objects
	objHead *ObjNode
	objTail *ObjNode
	objLock sync.RWMutex

	// is set by board
	objAmount int
}

/*
startup method
*/
func (bw *boardWorker) work() {
	go bw.doWork()
}

func (bw *boardWorker) doWork() {
	for {
		// wait for activation and lock our list
		mode := <-bw.activate

		// iterate through
		node := bw.objHead
		ind := 0 // if another object is added during run we wont reach it.
		for node != nil && ind < bw.objAmount {
			switch mode {
			case WModeDelete:
				if node.Pos.Remove() {
					bw.DeleteObj(node)
					bw.myBoard.wAmountDec(bw.workerID)
				}
			case WModeMove:
				if node.Move != nil {
					node.Move.UpdatePosition()
				}
			case WModeCollide:
				if node.Coll != nil {
					bw.myBoard.CheckPossibleCollisions(node.Coll)
				}
			case WModeUpdate:
				if node.Update != nil {
					node.Update.Update()
				}
			}
			node = node.next
			ind++
		}

		// signal we are done
		bw.myBoard.workWait.Done()
	}
}

/*
AddObj adds an object node to board
*/
func (bw *boardWorker) AddObj(newNode *ObjNode) {

	bw.objLock.Lock()
	defer bw.objLock.Unlock()

	// add to own list
	if bw.objHead == nil {
		bw.objHead = newNode
		bw.objTail = newNode
	} else {
		bw.objTail.next = newNode
		newNode.prev = bw.objTail
		bw.objTail = newNode
	}
}

/*
DeleteObj deletes an obj node from board
*/
func (bw *boardWorker) DeleteObj(node *ObjNode) {

	bw.objLock.Lock()
	defer bw.objLock.Unlock()

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		// head node deleted
		bw.objHead = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		// tail node deleted
		bw.objTail = node.prev
	}

	if node.Pos == nil {
		logError("undeletable node detected, skipping delete!")
	}

	chunk, chunkObj := node.Pos.GetNodeInfo()
	chunk.DeleteObj(chunkObj)
}

// #############################################################################
// #							UPDATE
// #############################################################################

/*
Tick makes one tick on the board. Returns duration of update
*/
func (b *Board) Tick() time.Duration {
	start := time.Now()

	for _, mode := range WorkMode {

		for i := 0; i < b.workerAmount; i++ {
			b.workWait.Add(1)
			b.worker[i].objAmount = b.workerObjAmount[i]
			b.worker[i].activate <- mode
		}

		b.workWait.Wait()
	}

	return time.Since(start)
}
