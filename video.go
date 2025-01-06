// Contains helpers interact with the video RAM

package gbalib

import "unsafe"

// Types

type Color uint16

type MemVramType struct {
	Ptr          *uint16
	FrontPagePtr *uint16
	BackPagePtr  *uint16
	VidPage      *uint16
	M3           *[M3Height][M3Width]Color
	M4FrontPage  *[M4Height][M4Width]ColorId
	M4BackPage   *[M4Height][M4Width]ColorId
	M5FrontPage  *[M5Height][M5Width]Color
	M5BackPage   *[M5Height][M5Width]Color
}

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
	ClrBlack   = 0x0000
	ClrRed     = 0x001F
	ClrLime    = 0x03E0
	ClrYellow  = 0x03FF
	ClrBlue    = 0x7C00
	ClrMagenta = 0x7C1F
	ClrCyan    = 0x7FE0
	ClrWhite   = 0xFFFF
)

// Memory maps

const MemVramAddr = 0x06000000
const MemVramFrontAddr = MemVramAddr
const MemVramBackAddr = MemVramAddr + VramPageSize

// Memory handles

var MemVram = MemVramType{
	Ptr:          (*uint16)(unsafe.Pointer(uintptr(MemVramAddr))),
	FrontPagePtr: (*uint16)(unsafe.Pointer(uintptr(MemVramAddr))),
	BackPagePtr:  (*uint16)(unsafe.Pointer(uintptr(MemVramBackAddr))),
	VidPage:      (*uint16)(unsafe.Pointer(uintptr(MemVramBackAddr))),
	M3:           (*[M3Height][M3Width]Color)(unsafe.Pointer(uintptr(MemVramAddr))),
	M4FrontPage:  (*[M4Height][M4Width]ColorId)(unsafe.Pointer(uintptr(MemVramAddr))),
	M4BackPage:   (*[M4Height][M4Width]ColorId)(unsafe.Pointer(uintptr(MemVramBackAddr))),
	M5FrontPage:  (*[M5Height][M5Width]Color)(unsafe.Pointer(uintptr(MemVramAddr))),
	M5BackPage:   (*[M5Height][M5Width]Color)(unsafe.Pointer(uintptr(MemVramBackAddr))),
}

// Helpers

func MakeRgb15(r, g, b uint32) Color {
	return Color((r & 0x1F) | ((g & 0x1F) << 5) | ((b & 0x1F) << 10))
}

func M3Plot(x, y int32, color Color) {
	vramSlice := unsafe.Slice(MemVram.Ptr, M3Width*M3Height)
	Bmp16Plot(x, y, M3Width, uint16(color), &vramSlice)
}

func M3HLine(x1, y, x2 int32, color Color) {
	vramSlice := unsafe.Slice(MemVram.Ptr, M3Width*M3Height)
	Bmp16HLine(x1, y, x2, M3Width, uint16(color), &vramSlice)
}

func M3VLine(x, y1, y2 int32, color Color) {
	vramSlice := unsafe.Slice(MemVram.Ptr, M3Width*M3Height)
	Bmp16VLine(x, y1, y2, M3Width, uint16(color), &vramSlice)
}

func M3Line(x1, y1, x2, y2 int32, color Color) {
	vramSlice := unsafe.Slice(MemVram.Ptr, M3Width*M3Height)
	Bmp16Line(x1, y1, x2, y2, M3Width, uint16(color), &vramSlice)
}

func M3Rect(left, top, right, bottom int32, color Color) {
	vramSlice := unsafe.Slice(MemVram.Ptr, M3Width*M3Height)
	Bmp16Rect(left, top, right, bottom, M3Width, uint16(color), &vramSlice)
}

func M3Frame(left, top, right, bottom int32, color Color) {
	vramSlice := unsafe.Slice(MemVram.Ptr, M3Width*M3Height)
	Bmp16Frame(left, top, right, bottom, M3Width, uint16(color), &vramSlice)
}

func M3Fill(color Color) {
	MemSet32((*uint32)(unsafe.Pointer(MemVram.Ptr)), uint32(color)|uint32(color)<<16, uint32(M3Width*M3Height/2))
}

func M3Clear() {
	MemSet32((*uint32)(unsafe.Pointer(MemVram.Ptr)), ClrBlack, uint32(M3Width*M3Height/2))
}

func M4Plot(x, y int32, colorId ColorId) {
	vramSlice := unsafe.Slice(MemVram.VidPage, M4Width*M4Height/2)
	Bmp8Plot(x, y, M4Width, uint8(colorId), &vramSlice)
}

func M4HLine(x1, y, x2 int32, colorId ColorId) {
	vramSlice := unsafe.Slice(MemVram.VidPage, M4Width*M4Height/2)
	Bmp8HLine(x1, y, x2, M4Width, uint8(colorId), &vramSlice)
}

func M4VLine(x, y1, y2 int32, colorId ColorId) {
	vramSlice := unsafe.Slice(MemVram.VidPage, M4Width*M4Height/2)
	Bmp8VLine(x, y1, y2, M4Width, uint8(colorId), &vramSlice)
}

func M4Line(x1, y1, x2, y2 int32, colorId ColorId) {
	vramSlice := unsafe.Slice(MemVram.VidPage, M4Width*M4Height/2)
	Bmp8Line(x1, y1, x2, y2, M4Width, uint8(colorId), &vramSlice)
}

func M4Rect(left, top, right, bottom int32, colorId ColorId) {
	vramSlice := unsafe.Slice(MemVram.VidPage, M4Width*M4Height/2)
	Bmp8Rect(left, top, right, bottom, M4Width, uint8(colorId), &vramSlice)
}

func M4Frame(left, top, right, bottom int32, colorId ColorId) {
	vramSlice := unsafe.Slice(MemVram.VidPage, M4Width*M4Height/2)
	Bmp8Frame(left, top, right, bottom, M4Width, uint8(colorId), &vramSlice)
}

func M4Fill(colorId ColorId) {
	MemSet32((*uint32)(unsafe.Pointer(MemVram.VidPage)), uint32(colorId)|uint32(colorId)<<8|uint32(colorId)<<16|uint32(colorId)<<24, uint32(M4Width*M4Height/4))
}

func M4Clear() {
	MemSet32((*uint32)(unsafe.Pointer(MemVram.VidPage)), ClrBlack, uint32(M4Width*M4Height/4))
}

func M5Plot(x, y int32, color Color) {
	vramSlice := unsafe.Slice(MemVram.VidPage, M5Width*M5Height)
	Bmp16Plot(x, y, M5Width, uint16(color), &vramSlice)
}

func M5HLine(x1, y, x2 int32, color Color) {
	vramSlice := unsafe.Slice(MemVram.VidPage, M5Width*M5Height)
	Bmp16HLine(x1, y, x2, M5Width, uint16(color), &vramSlice)
}

func M5VLine(x, y1, y2 int32, color Color) {
	vramSlice := unsafe.Slice(MemVram.VidPage, M5Width*M5Height)
	Bmp16VLine(x, y1, y2, M5Width, uint16(color), &vramSlice)
}

func M5Line(x1, y1, x2, y2 int32, color Color) {
	vramSlice := unsafe.Slice(MemVram.VidPage, M5Width*M5Height)
	Bmp16Line(x1, y1, x2, y2, M5Width, uint16(color), &vramSlice)
}

func M5Rect(left, top, right, bottom int32, color Color) {
	vramSlice := unsafe.Slice(MemVram.VidPage, M5Width*M5Height)
	Bmp16Rect(left, top, right, bottom, M5Width, uint16(color), &vramSlice)
}

func M5Frame(left, top, right, bottom int32, color Color) {
	vramSlice := unsafe.Slice(MemVram.VidPage, M5Width*M5Height)
	Bmp16Frame(left, top, right, bottom, M5Width, uint16(color), &vramSlice)
}

func M5Fill(color Color) {
	MemSet32((*uint32)(unsafe.Pointer(MemVram.VidPage)), uint32(color)|uint32(color)<<16, uint32(M5Width*M5Height/2))
}

func M5Clear() {
	MemSet32((*uint32)(unsafe.Pointer(MemVram.VidPage)), ClrBlack, uint32(M5Width*M5Height/2))
}

func VidFlip() *uint16 {
	MemVram.VidPage = (*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(MemVram.VidPage)) ^ VramPageSize))
	MemIo.RegDisplayControl.Set(MemIo.RegDisplayControl.Get() ^ DisplayControlPageSelect)

	return MemVram.VidPage
}

func VidVSync() {
	for MemIo.RegVCount.Get() >= 160 { // wait till VDraw
	}
	for MemIo.RegVCount.Get() < 160 { // wait till VBlank
	}
}
