package gbaDraw

func (dsp GbaDisplay) Filled2PointRect(x1, y1, x2, y2 int16, c colorIndex) {
	if y1 > y2 {
		tmp := y1
		y1 = y2
		y2 = tmp
	}
	for y := y1; y <= y2; y++ {
		dsp.HLine(x1, x2, y, c)
	}
}

func (dsp GbaDisplay) HLine(x1, x2, y int16, c colorIndex) {
	if x1 > x2 {
		tmp := x1
		x1 = x2
		x2 = tmp
	}
	//check if starting with odd pixel
	if x1&1 == 1 {
		dsp.SetPixel(x1, y, c)
		x1++
	}
	// set pixel in blocks of two
	c2 := uint16(c) | (uint16(c) << 8)
	for x := (x1 >> 1); x <= (x2 >> 1); x++ {
		dsp.vRam[dsp.drawPage][y][x].Set(c2)
	}
	// check if ending with even pixel
	if x2&1 == 0 {
		dsp.SetPixel(x2, y, c)
	}
}

func (dsp GbaDisplay) FilledCircle(x, y, r int16, c colorIndex) {
	// dsp.SetPixel(x, y-r, c)
	// dsp.SetPixel(x, y+r, c)
	// dsp.HLine(x-r, x+r, y, c)
	dsp.Filled2PointRect(x-r, y-r, x+r, y+r, c)
}
