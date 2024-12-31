// Contains helpers interact with the video RAM

package gbalib

import "unsafe"

// Types

type Color uint16

// General constants

const (
	ScreenWidth  = 240
	ScreenHeight = 160
)

// Memory maps

const MemVramAddr = 0x06000000

// Memory handles

var MemVram = (*uint16)(unsafe.Pointer(uintptr(MemVramAddr)))
var MemVramMode3 = (*[ScreenWidth * ScreenHeight]Color)(unsafe.Pointer(MemVram))

// Helpers

func MakeRgb15(r, g, b uint32) Color {
	return Color((r & 0x1F) | ((g & 0x1F) << 5) | ((b & 0x1F) << 10))
}

func M3Plot(x, y int, c Color) {
	MemVramMode3[y*ScreenWidth+x] = c
}
