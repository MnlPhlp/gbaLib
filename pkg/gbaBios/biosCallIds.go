package gbaBios

// list of IDs of GBA Bios Calls
// taken from https://www.coranac.com/tonc/text/swi.htm#sec-funs
const (
	Instr_SoftReset = iota << 16
	Instr_RegisterRamReset
	Instr_Halt           = "swi 0x20000"
	Instr_Stop           = "swi 0x30000"
	Instr_IntrWait       = "swi 0x40000"
	Instr_VBlankIntrWait = "swi 0x50000"
	Instr_Div
	Instr_DivArm
	Instr_Sqrt
	Instr_ArcTan
	Instr_ArcTan2
	Instr_CPUSet
	Instr_CPUFastSet
	Instr_BiosChecksum
	Instr_BgAffineSet
	Instr_ObjAffineSet
	Instr_BitUnPack
	Instr_LZ77UnCompWRAM
	Instr_LZ77UnCompVRAM
	Instr_HuffUnComp
	Instr_RLUnCompWRAM
	Instr_RLUnCompVRAM
	Instr_Diff8bitUnFilterWRAM
	Instr_Diff8bitUnFilterVRAM
	Instr_Diff16bitUnFilter
	Instr_SoundBiasChange
	Instr_SoundDriverInit
	Instr_SoundDriverMode
	Instr_SoundDriverMain
	Instr_SoundDriverVSync
	Instr_SoundChannelClear
	Instr_MIDIKey2Freq
	Instr_MusicPlayerOpen
	Instr_MusicPlayerStart
	Instr_MusicPlayerStop
	Instr_MusicPlayerContinue
	Instr_MusicPlayerFadeOut
	Instr_MultiBoot
	Instr_HardReset
	Instr_CustomHalt
	Instr_SoundDriverVSyncOff
	Instr_SoundDriverVSyncOn
	Instr_GetJumpList
)
