package gbaDraw

import "github.com/MnlPhlp/gbaLib/pkg/gbaBios"

var vSyncEnabled = false

// VSync waits for the next display refresh
func VSync() {
	if !vSyncEnabled {
		gbaBios.VBlankIntrWait_Enable()
		vSyncEnabled = true
	}
	// this uses a bios call to halt the processor until the VBlank interrupt occurs
	gbaBios.VBlankIntrWait()
}
