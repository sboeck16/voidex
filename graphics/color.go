/*
Package graphics provides method to generate or load pictures, as well as the
picture directly
*/
package graphics

var (
	// ColorRed is exactly that
	ColorRed = &Col{0xffff, 0, 0, 0xffff}
	// ColorGreen is exactly that
	ColorGreen = &Col{0, 0xffff, 0, 0xffff}
	// ColorBlue is exactly that
	ColorBlue = &Col{0, 0, 0xffff, 0xffff}
	// ColorWhite is white
	ColorWhite = &Col{0xffff, 0xffff, 0xffff, 0xffff}
	// ColorGray is gray
	ColorGray = &Col{0x7fff, 0x7fff, 0x7fff, 0xffff}
	// ColorDarkGray is a darker gray
	ColorDarkGray = &Col{0x3fff, 0x3fff, 0x3fff, 0xffff}
	// ColorBlack is a black color
	ColorBlack = &Col{0, 0, 0, 0xffff}
)

/*
Col provides a color.Color compatible struct
*/
type Col struct {
	r uint32
	g uint32
	b uint32
	a uint32
}

func (c *Col) RGBA() (uint32, uint32, uint32, uint32) {
	return c.r, c.g, c.b, c.a
}
