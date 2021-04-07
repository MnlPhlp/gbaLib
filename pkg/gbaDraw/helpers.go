package gbaDraw

import "github.com/MnlPhlp/gbaLib/pkg/registers"

// VSync waits for the next display refresh
func VSync() {
	for registers.Video.VCount.Get() >= 160 {
	}
	for registers.Video.VCount.Get() < 160 {
	}
	// this uses a bios call to halt the processor until the VBlank interrupt occurs
	//gbaBios.VBlankIntrWait()
}
