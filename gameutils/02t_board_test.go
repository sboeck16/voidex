package gameutils

import (
	"math/rand"
	"testing"
	"time"
)

var (
	stressAmount  = 10000
	stressUpdates = 600
)

type testMov struct {
	MovingObject
}

func TestChunkAcquire(t *testing.T) {
	deb("RUNNING chunk acquire")

	mov := new(testMov)
	b := NewBoard(1000, 1000)

	b.RegisterObj(mov, mov, nil, nil, nil)

	mov.dx = 1.5

	for i := 0; i < 100; i++ {
		b.Tick()
	}

	// panic occur if object didn't move correct
	if b.chunks[0][0].objHead != nil {
		t.Error("chunk deletion didnt work")
	}
	if b.chunks[1][0].objHead == nil {
		t.Error("chunk obj add didt work")
	}
	if mov.x < 100 {
		t.Error("update position for board added obj is wrong")
	}
}

func TestBoardObjStress(t *testing.T) {

	deb("RUNNING: stress test, #obj #updates", stressAmount, stressUpdates)

	ChunkWorker = 10
	b := NewBoard(1000, 1000)
	// time
	st := time.Now()
	for i := 0; i < stressAmount; i++ {
		mov := new(testMov)
		mov.dx = 1 - rand.Float64()*2
		mov.dy = 1 - rand.Float64()*2

		mov.x = rand.Float64() * 10000
		mov.y = rand.Float64() * 10000

		b.RegisterObj(mov, mov, nil, nil, nil)
	}
	addDur := time.Since(st)

	st = time.Now()
	for i := 0; i < stressUpdates; i++ {
		b.Tick()
	}
	updDur := time.Since(st)
	deb(" -- overall #objAdd, #time, #time per tick",
		addDur, updDur, updDur/time.Duration(stressUpdates))
}

func TestDeletion(t *testing.T) {
	deb("RUNNING: deletion test")

	b := NewBoard(1000, 1000)
	mov1 := new(testMov)
	mov2 := new(testMov)
	mov2.removeMe = true
	mov3 := new(testMov)
	mov3.dx = 1.0
	b.RegisterObj(mov1, mov1, nil, nil, nil)
	b.RegisterObj(mov2, mov2, nil, nil, nil)
	b.RegisterObj(mov3, mov3, nil, nil, nil)
	b.Tick()

	if b.worker[1].objHead != nil {
		t.Error("Removing of Pos obj didnt work")
	}

	if b.chunks[0][0].objHead.next.next != nil {
		t.Error("chunk list remove didnt work")
	}

	mov1.removeMe = true
	b.Tick()

	if b.chunks[0][0].objHead.next != nil {
		t.Error("chunk list remove didnt work")
	}
}
