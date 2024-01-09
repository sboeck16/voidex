package gameutils

import (
	"voidex/graphics"

	eb "github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten/v2/text"
)

/*
Button holds the clickable UI element.
*/
type Button struct {
	ClickObject
	bgColor   *graphics.Col
	textColor *graphics.Col
	textFont  font.Face
}

/*
NewButton is utility method to implement a clickable as a button ui element
on a board.
*/
func NewButton(x, y float64, width, height int,
	handle func(*ClickObject, eb.MouseButton)) *Button {

	ret := new(Button)
	ret.clickHandle = handle
	ret.x = x
	ret.y = y
	ret.collSize = CollModeIsInImageBorder

	// method, function for reseting?
	ret.img = eb.NewImage(width, height)

	// set fallback
	ret.bgColor = ButtonBGColor
	ret.textColor = ButtonTextCol

	return ret
}

/*
NewButton is utility method to implement a clickable as a button ui element
on a board.
*/
func (b *Board) NewButton(x, y float64, width, height int,
	handle func(*ClickObject, eb.MouseButton)) *Button {

	nb := NewButton(x, y, width, height, handle)
	b.RegisterObj(nb, nil, nil, nil, nb)
	return nb
}

/*
SetText sets the button text and renders the button image.
*/
func (bu *Button) SetText(sx, sy int, str string) {

	bu.img.Fill(bu.bgColor)
	text.Draw(bu.img, str, bu.textFont, sx, sy, bu.textColor)
}

/*
SetColorAndFont sets the colors of a given button. As well as the font.
*/
func (bu *Button) SetColorAndFont(textCol, bgCol *graphics.Col, f font.Face) {
	if textCol != nil {
		bu.textColor = textCol
	}
	if bgCol != nil {
		bu.bgColor = bgCol
	}
	if f != nil {
		bu.textFont = f
	}
}
