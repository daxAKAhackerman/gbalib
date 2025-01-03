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
	M4Width  = ScreenWidth
	M4Height = ScreenHeight
	M5Width  = 160
	M5Height = 128
)

const VramPageSize = 0xA000

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

const (
	MemVramAddr     = 0x06000000
	MemVramBackAddr = MemVramAddr + VramPageSize
)

// Memory handles

var MemVram = (*uint16)(unsafe.Pointer(uintptr(MemVramAddr)))
var MemVramBack = (*uint16)(unsafe.Pointer(uintptr(MemVramBackAddr)))
var M3MemVram = (*[M3Height][M3Width]Color)(unsafe.Pointer(MemVram))
var M4MemVram = (*[M4Height][M4Width]uint8)(unsafe.Pointer(MemVram))
var M4MemVramBack = (*[M4Height][M4Width]uint8)(unsafe.Pointer(MemVramBack))
var M5MemVram = (*[M5Height][M5Width]Color)(unsafe.Pointer(MemVram))
var M5MemVramBack = (*[M5Height][M5Width]Color)(unsafe.Pointer(MemVramBack))

var VidPage = MemVramBack

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

func M3Clear() {
	MemSet32((*uint32)(unsafe.Pointer(MemVram)), 0x00, uint32(M3Width*M3Height/2))
}

func M4Plot(x, y int32, cid uint8) {
	s := unsafe.Slice(VidPage, M4Width*M4Height/2)
	Bmp8Plot(x, y, M4Width, cid, &s)
}

func M5Plot(x, y int32, c Color) {
	s := unsafe.Slice(VidPage, M5Width*M5Height)
	Bmp16Plot(x, y, M5Width, uint16(c), &s)
}

func M5HLine(x1, y, x2 int32, c Color) {
	s := unsafe.Slice(VidPage, M5Width*M5Height)
	Bmp16HLine(x1, y, x2, M5Width, uint16(c), &s)
}

func M5VLine(x, y1, y2 int32, c Color) {
	s := unsafe.Slice(VidPage, M5Width*M5Height)
	Bmp16VLine(x, y1, y2, M5Width, uint16(c), &s)
}

func M5Line(x1, y1, x2, y2 int32, c Color) {
	s := unsafe.Slice(VidPage, M5Width*M5Height)
	Bmp16Line(x1, y1, x2, y2, M5Width, uint16(c), &s)
}

func M5Rect(left, top, right, bottom int32, c Color) {
	s := unsafe.Slice(VidPage, M5Width*M5Height)
	Bmp16Rect(left, top, right, bottom, M5Width, uint16(c), &s)
}

func M5Frame(left, top, right, bottom int32, c Color) {
	s := unsafe.Slice(VidPage, M5Width*M5Height)
	Bmp16Frame(left, top, right, bottom, M5Width, uint16(c), &s)
}

func M5Fill(c Color) {
	MemSet32((*uint32)(unsafe.Pointer(VidPage)), uint32(c)|uint32(c)<<16, uint32(M5Width*M5Height/2))
}

func M5Clear() {
	MemSet32((*uint32)(unsafe.Pointer(VidPage)), 0x00, uint32(M5Width*M5Height/2))
}

func VidFlip() *uint16 {
	VidPage = (*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(VidPage)) ^ VramPageSize))
	MemIo.RegDisplayControl.Set(MemIo.RegDisplayControl.Get() ^ DisplayControlPageSelect)

	return VidPage
}
