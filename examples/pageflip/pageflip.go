package main

import (
	"gbalib"
	"unsafe"
)

func loadGfx() {
	frontptr := *(*[]uint32)(unsafe.Pointer(gbalib.M4MemVramFront))
	backptr := *(*[]uint32)(unsafe.Pointer(gbalib.M4MemVramBack))
	for i := 0; i < 16; i++ {
		gbalib.MemCpy32(&frontptr[i*120], (*uint32)(unsafe.Pointer(&FrontBitmap[i*144/4])), 144)
		gbalib.MemCpy32(&backptr[i*120], (*uint32)(unsafe.Pointer(&BackBitmap[i*144/4])), 144)
	}

	// You don't have to do everything with memcpy.
	// In fact, for small blocks it might be better if you didn't.
	// Just mind your types, though. No sense in copying from a 32bit
	// array to a 16bit one.
	for i := 0; i < 8; i++ {
		gbalib.MemPalBg[i] = gbalib.Color(FrontPal[i])
	}
}

func main() {
	loadGfx()
	gbalib.MemIo.RegDisplayControl.Init(gbalib.DisplayControlBg2, gbalib.DisplayControlMode4)
	for {
	}
}
