package gbaDraw

import (
	"image/color"
	"runtime/volatile"
	"unsafe"

	"github.com/MnlPhlp/gbaLib/pkg/gbaBios"
	"github.com/MnlPhlp/gbaLib/pkg/registers"
)

type GbaDisplay struct {
	vRam [2](*[160][120]volatile.Register16)
}
type ColorIndex uint8
type colorPalette *[256]volatile.Register16

var Display = GbaDisplay{
	vRam: [2](*[160][120]volatile.Register16){
		(*[160][120]volatile.Register16)(unsafe.Pointer(uintptr(0x06000000))), // Page 1
		(*[160][120]volatile.Register16)(unsafe.Pointer(uintptr(0x0600A000))), // Page 2
	},
}
var (
	palette   = (*[256]volatile.Register16)(unsafe.Pointer(uintptr(0x05000000)))
	lastIndex = ColorIndex(0)
	drawPage  = 1
)

func toRGB15(c color.RGBA) uint16 {
	return uint16(c.R)&0x1f | uint16(c.G)&0x1f<<5 | uint16(c.B)&0x1f<<10
}

func ToColorIndex(c color.RGBA) ColorIndex {
	lastIndex++
	palette[lastIndex].Set(toRGB15(c))
	return lastIndex
}

//go:inline
func (dsp GbaDisplay) SetPixelPallette(x, y int16, c ColorIndex) {
	dsp.vRam[drawPage][y][x>>1].ReplaceBits(uint16(c), 0xff, uint8(x&1)<<3)
}

func (dsp GbaDisplay) Configure() {
	value := uint16(0x04)              //Set Display Mode
	value |= 1 << 0xA                  //Enable BG2
	registers.Video.DispCnt.Set(value) //set Value
	// enable Function for VSync
	gbaBios.VBlankIntrWait_Enable()
	// set some colors
	palette[Black].Set(0)
	palette[White].Set(0xffff)
	palette[Red].Set(toRGB15(color.RGBA{R: 255}))
	palette[Green].Set(toRGB15(color.RGBA{G: 255}))
	palette[Blue].Set(toRGB15(color.RGBA{B: 255}))
	lastIndex = Blue
}

func (dsp GbaDisplay) Display() error {
	old := registers.Video.DispCnt.Get()
	registers.Video.DispCnt.Set(old ^ (uint16(drawPage) << 4)) // flip display
	drawPage ^= 1                                              // switch drawPage
	return nil
}

// Be compatible to tinydraw/tinyfont
func (dsp GbaDisplay) SetPixel(x, y int16, c color.RGBA) {
	dsp.SetPixelPallette(x, y, ToColorIndex(c))
}

func (dsp GbaDisplay) Size() (int16, int16) {
	return 240, 160
}
