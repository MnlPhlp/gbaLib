package interrupts

import (
	"machine"
	"runtime/interrupt"

	"github.com/MnlPhlp/gbaLib/pkg/registers"
)

type interruptHandler func()

var interrupts = make(map[interrupt.Interrupt]interruptHandler)

func isr(i interrupt.Interrupt) {
	if handler, ok := interrupts[i]; ok {
		handler()
	}
}

func setupInterrupt(i interrupt.Interrupt, f func()) {
	interrupts[i] = interruptHandler(f)
	i.Enable()
}

func SetVBlankInterrupt(f func()) {
	// enable the interrupt
	registers.Video.DispStat.SetBits(1 << 3)
	// create a new Interrupt and store the function
	i := interrupt.New(machine.IRQ_VBLANK, isr)
	setupInterrupt(i, f)
}

func Disable() {
	// disable interrupts
	interrupt.Disable()
}
