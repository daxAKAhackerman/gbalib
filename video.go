// Contains helpers interact with the video RAM

package gbalib

import "unsafe"

const (
	ScreenWidth  = 240
	ScreenHeight = 160
)

const MemVramAddr = 0x06000000

type Color uint16

var MemVram = (*uint16)(unsafe.Pointer(uintptr(MemVramAddr)))
var MemVramMode3 = (*[ScreenWidth * ScreenHeight]Color)(unsafe.Pointer(MemVram))

func MakeRgb15(r, g, b uint32) Color {
	return Color((r & 0x1F) | ((g & 0x1F) << 5) | ((b & 0x1F) << 10))
}

func M3Plot(x, y int, c Color) {
	MemVramMode3[y*ScreenWidth+x] = c
}
