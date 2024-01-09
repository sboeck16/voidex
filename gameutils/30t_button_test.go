package gameutils

import (
	"testing"
	"voidex/graphics"

	eb "github.com/hajimehoshi/ebiten/v2"
)

var (
	clicked bool
)

func TestButton(t *testing.T) {
	deb("RUNNING: button test")

	// create game
	board := NewBoard(testWidth, testHeight)
	game := NewGame(testWidth, testHeight)
	game.AddBoard(board)

	bu := board.NewButton(100, 100, 200, 50, testClickHandle)
	bu.SetColorAndFont(graphics.ColorDarkGray, graphics.ColorWhite, TextFontMiddle)
	bu.SetText(20, 30, "click me")

	if displayGame {
		game.StartGame("TEST button")
	} else {
		board.BoardClicked(110, 110, eb.MouseButtonLeft)
		if !clicked {
			t.Error("click button didn't work")
		}
	}

}

func testClickHandle(_ *ClickObject, _ eb.MouseButton) {
	if displayGame {
		deb("button clicked")
	} else {
		clicked = true
	}
}
