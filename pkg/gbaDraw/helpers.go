package gbaDraw

import "github.com/MnlPhlp/gbaLib/pkg/gbaBios"

// VSync waits for the next display refresh
func VSync() {
	// this uses a bios call to halt the processor until the VBlank interrupt occurs
	gbaBios.VBlankIntrWait()
}
