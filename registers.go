package gbaLib

import (
	"runtime/volatile"
	"unsafe"
)

type videoRegisters struct {
	DispCnt,
	VCount,
	BG0Cnt,
	BG1Cnt,
	BG2Cnt,
	BG3Cnt,
	DispStat *volatile.Register16
}

type keyRegister struct {
	KeyCnt,
	KeyPad *volatile.Register16
}

type soundRegister struct {
	Sound1Cnt_L,
	Sound1Cnt_H,
	Sound1Cnt_X,
	Sound2Cnt_L,
	Sound2Cnt_H,
	Sound3Cnt_L,
	Sound3Cnt_H,
	Sound3Cnt_X,
	Sound4Cnt_L,
	Sound4Cnt_H,
	SoundCnt_L,
	SoundCnt_H,
	SoundCnt_X,
	SoundBias,
	WaveRam0_L,
	WaveRam0_H,
	WaveRam1_L,
	WaveRam1_H,
	WaveRam2_L,
	WaveRam2_H,
	WaveRam3_L,
	WaveRam3_H,
	FIFO_A_L,
	FIFO_A_H,
	FIFO_B_L,
	FIFO_B_H *volatile.Register16
}

type registers struct {
	Video videoRegisters
	Key   keyRegister
	Sound soundRegister
	IE,
	IF *volatile.Register16
}

var Register = registers{
	Video: videoRegisters{
		DispCnt:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000000))),
		DispStat: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000004))),
		VCount:   (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000006))),
		BG0Cnt:   (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000008))),
		BG1Cnt:   (*volatile.Register16)(unsafe.Pointer(uintptr(0x400000A))),
		BG2Cnt:   (*volatile.Register16)(unsafe.Pointer(uintptr(0x400000C))),
		BG3Cnt:   (*volatile.Register16)(unsafe.Pointer(uintptr(0x400000E))),
	},
	Sound: soundRegister{
		Sound1Cnt_L: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000060))),
		Sound1Cnt_H: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000062))),
		Sound1Cnt_X: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000064))),
		Sound2Cnt_L: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000068))),
		Sound2Cnt_H: (*volatile.Register16)(unsafe.Pointer(uintptr(0x400006C))),
		Sound3Cnt_L: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000070))),
		Sound3Cnt_H: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000072))),
		Sound3Cnt_X: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000074))),
		Sound4Cnt_L: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000078))),
		Sound4Cnt_H: (*volatile.Register16)(unsafe.Pointer(uintptr(0x400007C))),
		SoundCnt_L:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000080))),
		SoundCnt_H:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000082))),
		SoundCnt_X:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000084))),
		SoundBias:   (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000088))),
		WaveRam0_L:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000090))),
		WaveRam0_H:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000092))),
		WaveRam1_L:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000094))),
		WaveRam1_H:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000096))),
		WaveRam2_L:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000098))),
		WaveRam2_H:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x400009A))),
		WaveRam3_L:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x400009C))),
		WaveRam3_H:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x400009E))),
		FIFO_A_L:    (*volatile.Register16)(unsafe.Pointer(uintptr(0x40000A0))),
		FIFO_A_H:    (*volatile.Register16)(unsafe.Pointer(uintptr(0x40000A2))),
		FIFO_B_L:    (*volatile.Register16)(unsafe.Pointer(uintptr(0x40000A4))),
		FIFO_B_H:    (*volatile.Register16)(unsafe.Pointer(uintptr(0x40000A6))),
	},
	Key: keyRegister{
		KeyPad: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000130))),
		KeyCnt: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000132))),
	},
	IE: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000200))),
	IF: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000202))),
}
