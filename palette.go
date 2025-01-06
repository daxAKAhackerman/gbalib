// Contains helpers interact with the palette memory

package gbalib

import "unsafe"

// Types

type ColorId uint8

type MemPalType struct {
	BgPtr      *uint16
	Bg         *[256]Color
	BgBank     *[16][16]Color
	SpritePtr  *uint16
	Sprite     *[256]Color
	SpriteBank *[16][16]Color
}

// Memory maps

const MemPalAddr = 0x05000000
const MemPalBgAddr = MemPalAddr + 0x00
const MemPalSpriteAddr = MemPalAddr + 0x0200

// Memory handles

var MemPal = MemPalType{
	BgPtr:      (*uint16)(unsafe.Pointer(uintptr(MemPalBgAddr))),
	Bg:         (*[256]Color)(unsafe.Pointer(uintptr(MemPalBgAddr))),
	BgBank:     (*[16][16]Color)(unsafe.Pointer(uintptr(MemPalBgAddr))),
	SpritePtr:  (*uint16)(unsafe.Pointer(uintptr(MemPalSpriteAddr))),
	Sprite:     (*[256]Color)(unsafe.Pointer(uintptr(MemPalSpriteAddr))),
	SpriteBank: (*[16][16]Color)(unsafe.Pointer(uintptr(MemPalSpriteAddr))),
}
