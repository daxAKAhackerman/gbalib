package main

import (
	"gbalib"
	"unsafe"
)

func loadGfx() {
	frontptr := unsafe.Slice((*uint32)(unsafe.Pointer(gbalib.MemVram.FrontPagePtr)), (gbalib.MemVramBackAddr-gbalib.MemVramAddr)/4)
	backptr := unsafe.Slice((*uint32)(unsafe.Pointer(gbalib.MemVram.BackPagePtr)), (gbalib.MemVramBackAddr-gbalib.MemVramAddr)/4)
	for i := 0; i < 16; i++ {
		gbalib.MemCpy32(&frontptr[i*240/4], &FrontBitmap[i*144/4], 144/4)
		gbalib.MemCpy32(&backptr[i*240/4], &BackBitmap[i*144/4], 144/4)
	}

	// You don't have to do everything with memcpy.
	// In fact, for small blocks it might be better if you didn't.
	// Just mind your types, though. No sense in copying from a 32bit
	// array to a 16bit one.
	for i := 0; i < 8; i++ {
		gbalib.MemPal.Bg[i] = (*[16]gbalib.Color)(unsafe.Pointer(&FrontPal))[i]
	}
}

func main() {
	var i = 0

	loadGfx()
	gbalib.MemIo.RegDisplayControl.Init(gbalib.DisplayControlBg2, gbalib.DisplayControlMode4)

	for {
		for gbalib.KeyDownNow(gbalib.KeyInputStart) {
		}

		gbalib.VidVSync()
		i++
		if i == 60 {
			i = 0
			gbalib.VidFlip()
		}
	}
}
