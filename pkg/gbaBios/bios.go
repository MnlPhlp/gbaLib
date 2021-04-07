package gbaBios

import (
	"device/arm"

	"github.com/MnlPhlp/gbaLib/pkg/interrupts"
	"github.com/MnlPhlp/gbaLib/pkg/registers"
)

type biosCallId uint8

var (
	vBlankIntrEnabled = false
)

// enables all necessary interrupts for bios call 'VBlankIntrWait'
func VBlankIntrWait_Enable() {
	interrupts.SetVBlankInterrupt(func() {
		registers.Interrupt.IFBios.SetBits(1)
	})
}

func VBlankIntrWait() {
	arm.Asm(Instr_VBlankIntrWait)
}

func Halt() {
	arm.Asm(Instr_Halt)
}

func Stop() {
	arm.Asm(Instr_Stop)
}
