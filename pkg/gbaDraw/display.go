package gbaDraw

import (
	"runtime/volatile"
	"unsafe"

	"github.com/MnlPhlp/gbaLib/pkg/registers"
)

type GbaDisplay struct {
	vRam *[160][240]volatile.Register16
}

var Display = GbaDisplay{(*[160][240]volatile.Register16)(unsafe.Pointer(uintptr(0x06000000)))}

func (dsp GbaDisplay) SetPixel(x, y int16, c uint16) {
	dsp.vRam[y][x].Set(c)
}

func (dsp GbaDisplay) Configure() {
	registers.Video.DispCnt.Set(0x03)         //Set Display Mode
	registers.Video.DispCnt.SetBits(1 << 0xA) //Enable BG2
}

func (dsp GbaDisplay) Blank() {
	registers.Video.DispCnt.SetBits(1 << 7)
}

func (dsp GbaDisplay) Display() {
	registers.Video.DispCnt.ClearBits(1 << 7)
}
