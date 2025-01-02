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

const (
	M3Width  = ScreenWidth
	M3Height = ScreenHeight
)

const (
	ClrBlack   = Color(0x0000)
	ClrRed     = Color(0x001F)
	ClrLime    = Color(0x03E0)
	ClrYellow  = Color(0x03FF)
	ClrBlue    = Color(0x7C00)
	ClrMagenta = Color(0x7C1F)
	ClrCyan    = Color(0x7FE0)
	ClrWhite   = Color(0xFFFF)
)

// Memory maps

const MemVramAddr = 0x06000000

// Memory handles

var MemVram = (*uint16)(unsafe.Pointer(uintptr(MemVramAddr)))
var M3MemVram = (*[M3Width * M3Height]Color)(unsafe.Pointer(MemVram))

// Helpers

func MakeRgb15(r, g, b uint32) Color {
	return Color((r & 0x1F) | ((g & 0x1F) << 5) | ((b & 0x1F) << 10))
}

func M3Plot(x, y int32, c Color) {
	s := unsafe.Slice(MemVram, M3Width*M3Height)
	Bmp16Plot(x, y, M3Width, uint16(c), &s)
}

func M3HLine(x1, y, x2 int32, c Color) {
	s := unsafe.Slice(MemVram, M3Width*M3Height)
	Bmp16HLine(x1, y, x2, M3Width, uint16(c), &s)
}

func M3VLine(x, y1, y2 int32, c Color) {
	s := unsafe.Slice(MemVram, M3Width*M3Height)
	Bmp16VLine(x, y1, y2, M3Width, uint16(c), &s)
}

func M3Line(x1, y1, x2, y2 int32, c Color) {
	s := unsafe.Slice(MemVram, M3Width*M3Height)
	Bmp16Line(x1, y1, x2, y2, M3Width, uint16(c), &s)
}

func M3Rect(left, top, right, bottom int32, c Color) {
	s := unsafe.Slice(MemVram, M3Width*M3Height)
	Bmp16Rect(left, top, right, bottom, M3Width, uint16(c), &s)
}

func M3Frame(left, top, right, bottom int32, c Color) {
	s := unsafe.Slice(MemVram, M3Width*M3Height)
	Bmp16Frame(left, top, right, bottom, M3Width, uint16(c), &s)
}

func M3Fill(c Color) {
	MemSet32((*uint32)(unsafe.Pointer(MemVram)), uint32(c)|uint32(c)<<16, uint32(M3Width*M3Height/2))
}
