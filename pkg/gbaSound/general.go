package gbaSound

import "github.com/MnlPhlp/gbaLib/pkg/registers"

func Enable() {
	// turn on Sound circuit
	registers.Sound.SoundCnt_X.SetBits(1 << 7)
	// set Volume to full
	SetVolume(2)
}

func Disable() {
	registers.Sound.SoundCnt_X.ClearBits(1 << 7)
}

func SetVolume(lvl uint16) {
	if lvl > 2 {
		lvl = 2
	}
	registers.Sound.SoundCnt_H.SetBits(lvl)
}
