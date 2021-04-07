package gbaDraw

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
		width := (r / (r >> 1)) * (r - i) //math.Sqrt(float64((r*r - i*i))) / 2))
		width >>= 1
		dsp.HLine(x-width, x+width, y+i, c)
		dsp.HLine(x-width, x+width, y-i, c)
	}
}
