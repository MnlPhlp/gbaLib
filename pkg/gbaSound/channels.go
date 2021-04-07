package gbaSound

import (
	"runtime/volatile"

	"github.com/MnlPhlp/gbaLib/pkg/registers"
)

type SoundChannel struct {
	BuffLen   volatile.Register16
	BuffFreq  volatile.Register16
	regLen    *volatile.Register16
	regFreq   *volatile.Register16
	enableBit int16
}

const (
	soundEnable1 = iota + 8
	soundEnable2
	soundEnable3
	soundEnable4

	freqMask = 0x7ff
	lenMask  = 0x3f
	wdcMask  = 0x3
	wdcPos   = 0x6
)

func (c *SoundChannel) SetFrequency(freq uint16) {
	value := uint16(2048 - 131072/int(freq))
	c.BuffFreq.ReplaceBits(value&freqMask, freqMask, 0)
}

// set sound lenght in milliseconds
func (c *SoundChannel) SetLength(milliSec int) {
	c.SetLoop(false)
	value := uint16(-((milliSec >> 2) - 64))
	c.BuffLen.ReplaceBits(value&lenMask, lenMask, 0)
	c.SoundReset()
}

func (c *SoundChannel) SetWaveDutyCyle(value uint16) {
	c.BuffLen.ReplaceBits(value&wdcMask, wdcMask, wdcPos)
}

func (c *SoundChannel) SetLoop(loop bool) {
	if loop {
		c.BuffFreq.ClearBits(1 << 0xE)
	} else {
		c.BuffFreq.SetBits(1 << 0xE)
	}
}

func (c SoundChannel) SoundReset() {
	c.regFreq.SetBits(1 << 0xF)
}

func (c SoundChannel) Flush() {
	c.regFreq.Set(c.BuffFreq.Get())
	c.regLen.Set(c.BuffLen.Get())
}

func (c SoundChannel) Play() {
	c.Flush()
	c.SoundReset()
}

// Set the initial volume.
// 15 is the max value
func (c *SoundChannel) SetInitVol(vol uint16) {
	c.BuffLen.SetBits((vol & 0xf) << 0xC)
}

func (c *SoundChannel) EnableLeft() {
	registers.Sound.SoundCnt_L.SetBits(1 << c.enableBit)
}

func (c *SoundChannel) EnableRight() {
	registers.Sound.SoundCnt_L.SetBits(1 << (c.enableBit + 4))
}

func (c *SoundChannel) Enable() {
	c.EnableLeft()
	c.EnableRight()
	c.SetInitVol(15)
	c.SetWaveDutyCyle(2)
}
