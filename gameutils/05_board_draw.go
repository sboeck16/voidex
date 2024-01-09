package gameutils

import (
	"voidex/graphics"

	eb "github.com/hajimehoshi/ebiten/v2"
)

/*
Draw draws all elements
*/
func (b *Board) Draw() *eb.Image {
	ret := eb.NewImage(b.width, b.height)
	ret.Fill(graphics.ColorGreen)
	b.chunkLock.RLock()
	defer b.chunkLock.RUnlock()

	// TODO draw only needed elements (and test if we have to)
	for _, chunkSubMap := range b.chunks {
		for _, chunk := range chunkSubMap {
			node := chunk.objHead
			for node != nil {
				node.Pos.Draw(b.viewX, b.viewY, b.scale, ret)
				node = node.next
			}
		}
	}
	return ret
}
