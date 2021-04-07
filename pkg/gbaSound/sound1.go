package gbaSound

import (
	"runtime/volatile"

	"github.com/MnlPhlp/gbaLib/pkg/registers"
)

const (
	shiftMask  = 0b111
	sweepShift = 3
	timeMask   = 0b111
	timeShift  = 4
)

type Sound1Channel struct {
	SoundChannel
	BuffSweep volatile.Register16
	regSweep  *volatile.Register16
}

var Sound1 = Sound1Channel{
	SoundChannel: SoundChannel{
		enableBit: soundEnable1,
		regLen:    registers.Sound.Sound1Cnt_H,
		regFreq:   registers.Sound.Sound1Cnt_X,
		BuffLen:   volatile.Register16{},
		BuffFreq:  volatile.Register16{},
	},
	BuffSweep: volatile.Register16{},
	regSweep:  registers.Sound.Sound1Cnt_L,
}

func (c Sound1Channel) SetSweep(decrease bool, shifts, time uint16) {
	var sweepDec uint16
	if decrease {
		sweepDec = 1
	}
	value := (time & timeMask) << timeShift //set sweep time
	value |= sweepDec << sweepShift         // set increment/decrement
	value |= (shifts & shiftMask)           //set shifts
	c.regSweep.Set(value)
}
