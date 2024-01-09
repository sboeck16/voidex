package gameutils

import (
	"sync"

	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

/*
Game implements the ebiten Game interface
*/
type Game struct {

	// tick counter
	tick int

	// image per tick
	bgImage  *eb.Image
	drawlock sync.RWMutex

	// to generate overall image
	width, height int
	imageLock     sync.RWMutex
	drawOp        eb.DrawImageOptions

	// game elements are stored in boards
	boards []*Board

	// function pointer for additional logic
	additionalTickLogic func(int)
}

/*
Creates a new game with given width and height.
*/
func NewGame(w, h int) *Game {
	ret := new(Game)
	ret.width = w
	ret.height = h
	return ret
}

/*
AddBoard adds a board to game
*/
func (g *Game) AddBoard(b *Board) {
	g.boards = append(g.boards, b)
}

/*
StartGame starts the game with given title.
*/
func (g *Game) StartGame(title string) {
	eb.SetWindowSize(g.width, g.height)
	eb.SetWindowTitle(title)
	eb.RunGame(g)
}

/*
SetBGImg sets game background image
*/
func (g *Game) SetBGImg(img *eb.Image) {
	g.drawlock.Lock()
	defer g.drawlock.Unlock()
	g.bgImage = img
}

/*
GetBGImg returns the background image
*/
func (g *Game) GetBGImg() *eb.Image {
	g.drawlock.Lock()
	defer g.drawlock.Unlock()
	return g.bgImage
}

/*
SetTickFunction sets a function that is called every tick
*/
func (g *Game) SetTickFunction(f func(int)) {
	g.additionalTickLogic = f
}

// #############################################################################
// #							Ebiten Interface
// #############################################################################

/*
Update will be called by ebitenging trickering a "tick" for game logic.
*/
func (g *Game) Update() error {
	// handle input
	if inpututil.IsMouseButtonJustPressed(eb.MouseButtonLeft) {
		g.handleMouseClick(eb.MouseButtonLeft)
	}
	// maybe we will go different paths here? -> reduce redundant code?
	if inpututil.IsMouseButtonJustPressed(eb.MouseButtonRight) {
		g.handleMouseClick(eb.MouseButtonRight)
	}

	// trigger all boards to update
	for _, board := range g.boards {
		// skip disabled boards
		if !board.Enabled {
			continue
		}
		// for now serialize board updates
		// -> all in one board? or board / tick sync needed?
		board.Tick()
	}

	if g.additionalTickLogic != nil {
		g.additionalTickLogic(g.tick)
	}

	g.tick++
	return nil
}

/*
Layout TODO
*/
func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}

/*
Draw is called by ebitengine and used to display game.
*/
func (g *Game) Draw(screen *eb.Image) {

	// sync accessing draw files
	g.drawlock.Lock()
	defer g.drawlock.Unlock()

	// draw background
	if g.bgImage != nil {
		screen.DrawImage(g.bgImage, nil)
	}

	// trigger all boards to update
	for _, board := range g.boards {
		if !board.DrawMe {
			continue
		}
		// asynchronous draw could be implemented as this (if ever needed):
		//	go func(board *Board) {
		img := board.Draw()
		//ga.imageLock.Lock()
		g.drawOp.GeoM.Reset()
		g.drawOp.GeoM.Translate(board.GameX, board.GameY)
		screen.DrawImage(img, &g.drawOp)
		//ga.imageLock.Unlock()
		//	}(board)
	}
}

/*
search for board that is clicked and redirect click there. Maybe handler?
*/
func (g *Game) handleMouseClick(button eb.MouseButton) {
	xInt, yInt := eb.CursorPosition()
	x := float64(xInt)
	y := float64(yInt)
	// find board that has been clicked
	for _, board := range g.boards {
		// skip board that are not drawn
		if !board.DrawMe {
			continue
		}
		if board.GameX <= x && board.GameX+float64(board.width) >= x &&
			board.GameY <= y && board.GameY+float64(board.height) >= y {

			board.BoardClicked(x-board.GameX, y-board.GameY, button)
		}
	}
}
