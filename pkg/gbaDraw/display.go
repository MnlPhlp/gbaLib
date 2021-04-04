package gbaDraw

import (
	"image/color"
	"runtime/volatile"
	"unsafe"

	"github.com/MnlPhlp/gbaLib/pkg/registers"
)

type GbaDisplay struct {
	vRam     [2](*[160][120]volatile.Register16)
	drawPage uint8
}
type colorIndex uint8
type colorPalette *[256]volatile.Register16

var Display = GbaDisplay{
	vRam: [2](*[160][120]volatile.Register16){
		(*[160][120]volatile.Register16)(unsafe.Pointer(uintptr(0x06000000))), // Page 1
		(*[160][120]volatile.Register16)(unsafe.Pointer(uintptr(0x0600A000))), // Page 2
	},
	drawPage: 1,
}
var palette = (*[256]volatile.Register16)(unsafe.Pointer(uintptr(0x05000000)))
var lastIndex = colorIndex(0)

func toRGB15(c color.RGBA) uint16 {
	return uint16(c.R)&0x1f | uint16(c.G)&0x1f<<5 | uint16(c.B)&0x1f<<10
}

func ToColorIndex(c color.RGBA) colorIndex {
	lastIndex++
	palette[lastIndex].Set(toRGB15(c))
	return lastIndex
}

//go:inline
func (dsp *GbaDisplay) SetPixel(x, y int16, c colorIndex) {
	// if x&1 == 0 { // for even registers shift value left
	// 	dsp.vRam[dsp.drawPage][y][x>>1].ReplaceBits(uint16(c), 0xff, 8)
	// } else {
	// 	dsp.vRam[dsp.drawPage][y][x>>1].ReplaceBits(uint16(c), 0xff, 0)
	// }
	dsp.vRam[dsp.drawPage][y][x>>1].ReplaceBits(uint16(c), 0xff, uint8(x&1)<<3)
}

func (dsp GbaDisplay) Configure() {
	value := uint16(0x04)              //Set Display Mode
	value |= 1 << 0xA                  //Enable BG2
	registers.Video.DispCnt.Set(value) //set Value
	// set some colors
	palette[Black].Set(0)
	palette[White].Set(0xffff)
	palette[Red].Set(toRGB15(color.RGBA{R: 255}))
	palette[Green].Set(toRGB15(color.RGBA{G: 255}))
	palette[Blue].Set(toRGB15(color.RGBA{B: 255}))
	lastIndex = Blue
}

//go:inline
func (dsp *GbaDisplay) Display() error {
	old := registers.Video.DispCnt.Get()
	registers.Video.DispCnt.Set(old ^ (uint16(dsp.drawPage) << 4)) // flip display
	dsp.drawPage ^= 1                                              // switch drawPage
	return nil
}
