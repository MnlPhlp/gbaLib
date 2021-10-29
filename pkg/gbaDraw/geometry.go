package gbaDraw

import "math"

func (dsp GbaDisplay) Filled2PointRect(x1, y1, x2, y2 int16, c ColorIndex) {
	if y1 > y2 {
		tmp := y1
		y1 = y2
		y2 = tmp
	}
	for y := y1; y <= y2; y++ {
		dsp.HLine(x1, x2, y, c)
	}
}

func (dsp GbaDisplay) HLine(x1, x2, y int16, c ColorIndex) {
	if x1 > x2 {
		tmp := x1
		x1 = x2
		x2 = tmp
	}
	//check if starting with odd pixel
	if x1&1 == 1 {
		dsp.SetPixelPallette(x1, y, c)
		x1++
	}
	// set pixel in blocks of two
	c2 := uint16(c) | (uint16(c) << 8)
	for x := (x1 >> 1); x < (x2 >> 1); x++ {
		dsp.vRam[drawPage][y][x].Set(c2)
	}
	// check if ending with even pixel
	if x2&1 == 0 {
		dsp.SetPixelPallette(x2, y, c)
	} else {
		dsp.vRam[drawPage][y][x2>>1].Set(c2)
	}
}

func (dsp GbaDisplay) FilledDiamond(x, y, r int16, c ColorIndex) {
	dsp.SetPixelPallette(x, y-r, c)
	dsp.SetPixelPallette(x, y+r, c)
	dsp.HLine(x-r, x+r, y, c)
	for i := int16(1); i < r; i++ {
		width := (r / (r >> 1)) * (r - i)
		width >>= 1
		dsp.HLine(x-width, x+width, y+i, c)
		dsp.HLine(x-width, x+width, y-i, c)
	}
}

func (dsp GbaDisplay) CircleSqrt(x, y, r int16, c ColorIndex) {
	// This is a lot slower than Circle since it uses floating point math
	y1 := r
	rsq := r * r
	for x1 := int16(0); x1 <= y1; x1++ {
		// Just calculate y1 = sqrt(r^2 - x1^2)
		y1 := int16(math.Round(math.Sqrt(float64(rsq - x1*x1))))
		dsp.SetPixelPallette(x1+x, y1+y, c)
		dsp.SetPixelPallette(-x1+x, y1+y, c)
		dsp.SetPixelPallette(x1+x, -y1+y, c)
		dsp.SetPixelPallette(-x1+x, -y1+y, c)
		dsp.SetPixelPallette(y1+x, x1+y, c)
		dsp.SetPixelPallette(-y1+x, x1+y, c)
		dsp.SetPixelPallette(y1+x, -x1+y, c)
		dsp.SetPixelPallette(-y1+x, -x1+y, c)
	}
}

func (dsp GbaDisplay) Circle(x, y, r int16, c ColorIndex) {
	/// Use the Bresenham algorithm
	x1 := int16(0)
	y1 := r
	xsq := int16(0)
	rsq := r * r
	ysq := rsq
	// Loop x from 0 to the line x==y. Start y at r and each time
	// around the loop either keep it the same or decrement it.
	for x1 <= y1 {
		dsp.SetPixelPallette(x1+x, y1+y, c)
		dsp.SetPixelPallette(-x1+x, y1+y, c)
		dsp.SetPixelPallette(x1+x, -y1+y, c)
		dsp.SetPixelPallette(-x1+x, -y1+y, c)
		dsp.SetPixelPallette(y1+x, x1+y, c)
		dsp.SetPixelPallette(-y1+x, x1+y, c)
		dsp.SetPixelPallette(y1+x, -x1+y, c)
		dsp.SetPixelPallette(-y1+x, -x1+y, c)

		// New x^2 = (x+1)^2 = x^2 + 2x + 1
		xsq = xsq + 2*x1 + 1
		x1++
		// Potential new y^2 = (y-1)^2 = y^2 - 2y + 1
		ysq1 := ysq - 2*y1 + 1
		// Choose y or y-1, whichever gives smallest error
		a := xsq + ysq
		b := xsq + ysq1
		if a-rsq >= rsq-b {
			y1--
			ysq = ysq1
		}
	}

}

// Adapted from https://github.com/StephaneBunel/bresenham/blob/master/drawline.go

// dx > dy; x1 < x2; y1 < y2
func BresenhamDxXRYD(dsp GbaDisplay, x1, y1, x2, y2 int16, c ColorIndex) {
	dx, dy := x2-x1, 2*(y2-y1)
	e, slope := dx, 2*dx
	for ; dx != 0; dx-- {
		dsp.SetPixelPallette(x1, y1, c)
		x1++
		e -= dy
		if e < 0 {
			y1++
			e += slope
		}
	}
}

// dy > dx; x1 < x2; y1 < y2
func BresenhamDyXRYD(dsp GbaDisplay, x1, y1, x2, y2 int16, c ColorIndex) {
	dx, dy := 2*(x2-x1), y2-y1
	e, slope := dy, 2*dy
	for ; dy != 0; dy-- {
		dsp.SetPixelPallette(x1, y1, c)
		y1++
		e -= dx
		if e < 0 {
			x1++
			e += slope
		}
	}
}

// dx > dy; x1 < x2; y1 > y2
func BresenhamDxXRYU(dsp GbaDisplay, x1, y1, x2, y2 int16, c ColorIndex) {
	dx, dy := x2-x1, 2*(y1-y2)
	e, slope := dx, 2*dx
	for ; dx != 0; dx-- {
		dsp.SetPixelPallette(x1, y1, c)
		x1++
		e -= dy
		if e < 0 {
			y1--
			e += slope
		}
	}
}

func BresenhamDyXRYU(dsp GbaDisplay, x1, y1, x2, y2 int16, c ColorIndex) {
	dx, dy := 2*(x2-x1), y1-y2
	e, slope := dy, 2*dy
	for ; dy != 0; dy-- {
		dsp.SetPixelPallette(x1, y1, c)
		y1--
		e -= dx
		if e < 0 {
			x1++
			e += slope
		}
	}
}

// Generalized with integer
func (dsp GbaDisplay) Line(x1, y1, x2, y2 int16, c ColorIndex) {
	var dx, dy, e, slope int16

	// Because drawing p1 -> p2 is equivalent to draw p2 -> p1,
	// I sort points in x-axis order to handle only half of possible cases.
	if x1 > x2 {
		x1, y1, x2, y2 = x2, y2, x1, y1
	}

	dx, dy = x2-x1, y2-y1
	// Because point is x-axis ordered, dx cannot be negative
	if dy < 0 {
		dy = -dy
	}

	switch {

	// Is line a point ?
	case x1 == x2 && y1 == y2:
		dsp.SetPixelPallette(x1, y1, c)

	// Is line an horizontal ?
	case y1 == y2:
		for ; dx != 0; dx-- {
			dsp.SetPixelPallette(x1, y1, c)
			x1++
		}
		dsp.SetPixelPallette(x1, y1, c)

	// Is line a vertical ?
	case x1 == x2:
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for ; dy != 0; dy-- {
			dsp.SetPixelPallette(x1, y1, c)
			y1++
		}
		dsp.SetPixelPallette(x1, y1, c)

	// Is line a diagonal ?
	case dx == dy:
		if y1 < y2 {
			for ; dx != 0; dx-- {
				dsp.SetPixelPallette(x1, y1, c)
				x1++
				y1++
			}
		} else {
			for ; dx != 0; dx-- {
				dsp.SetPixelPallette(x1, y1, c)
				x1++
				y1--
			}
		}
		dsp.SetPixelPallette(x1, y1, c)

	// wider than high ?
	case dx > dy:
		if y1 < y2 {
			// BresenhamDxXRYD(img, x1, y1, x2, y2, c)
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				dsp.SetPixelPallette(x1, y1, c)
				x1++
				e -= dy
				if e < 0 {
					y1++
					e += slope
				}
			}
		} else {
			// BresenhamDxXRYU(img, x1, y1, x2, y2, c)
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				dsp.SetPixelPallette(x1, y1, c)
				x1++
				e -= dy
				if e < 0 {
					y1--
					e += slope
				}
			}
		}
		dsp.SetPixelPallette(x2, y2, c)

	// higher than wide.
	default:
		if y1 < y2 {
			// BresenhamDyXRYD(img, x1, y1, x2, y2, c)
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				dsp.SetPixelPallette(x1, y1, c)
				y1++
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		} else {
			// BresenhamDyXRYU(img, x1, y1, x2, y2, c)
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				dsp.SetPixelPallette(x1, y1, c)
				y1--
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		}
		dsp.SetPixelPallette(x2, y2, c)
	}
}
