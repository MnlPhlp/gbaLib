package buttons

import (
	"machine"
	"runtime/interrupt"

	"github.com/MnlPhlp/gbaLib/pkg/registers"
)

type Button uint16
type ButtonState uint16

const (
	A = Button(1 << iota)
	B
	Select
	Start
	Right
	Left
	Up
	Down
	R
	L
)

// initialize both states with no buttons pressed
var last ButtonState = 0x3FF
var current ButtonState = 0x3FF

func (state ButtonState) isDown(button Button) bool {
	// key bits are active low
	return uint16(state)&uint16(button) == 0
}

func (button Button) IsPressed() bool {
	return current.isDown(button)
}

func Poll() {
	last = current
	current = ButtonState(registers.Key.KeyPad.Get())
}

func keyIsr(interrupt.Interrupt) {
	Poll()
}

func EnableAutoPolling() {
	// enable the interrupt
	registers.Key.KeyCnt.SetBits(1 << 0xE)
	// enable all keys
	registers.Key.KeyCnt.SetBits(0b1111111111)
	// create a new Interrupt and store the function
	interrupt.New(machine.IRQ_KEYPAD, keyIsr).Enable()
}
