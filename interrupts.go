package gbaLib

import (
	"machine"
	"runtime/interrupt"
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

func SetKeypadInterrupt(f func()) {
	// enable the interrupt
	Register.Key.KeyCnt.SetBits(1 << 0xE)
	// enable all keys
	Register.Key.KeyCnt.SetBits(0b1111111111)
	// create a new Interrupt and store the function
	i := interrupt.New(machine.IRQ_KEYPAD, isr)
	setupInterrupt(i, f)
}

func SetVBlankInterrupt(f func()) {
	// enable the interrupt
	Register.Video.DispStat.SetBits(1 << 3)
	// create a new Interrupt and store the function
	i := interrupt.New(machine.IRQ_VBLANK, isr)
	setupInterrupt(i, f)
}

func Stop() {
	// disable interrupts
	Register.IE.Set(0)
	// keep running
	for {
	}
}
