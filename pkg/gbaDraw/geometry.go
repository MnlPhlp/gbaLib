package gbaDraw

import (
	"image/color"
)

func ToRgba15(c color.RGBA) uint16 {
	return uint16(c.R)&0x1f | uint16(c.G)&0x1f<<5 | uint16(c.B)&0x1f<<10
}

func Filled2PointRect(x1, y1, x2, y2 int16, c uint16) {
	xStep := int16(1)
	if x2 < x1 {
		xStep = -1
	}
	yStep := int16(1)
	if y2 < y1 {
		yStep = -1
	}
	for x := x1; x != x2; x += xStep {
		for y := y1; y != y2; y += yStep {
			Display.SetPixel(x, y, c)
		}
		Display.SetPixel(x, y2, c)
	}
	for y := y1; y != y2; y += yStep {
		Display.SetPixel(x2, y, c)
	}
	Display.SetPixel(x2, y2, c)
}
