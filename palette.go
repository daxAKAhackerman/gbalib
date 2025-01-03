// Contains helpers interact with the palette memory

package gbalib

import "unsafe"

// Types

// Memory maps

const MemPalAddr = 0x05000000
const MemPalBgAddr = MemPalAddr + 0x00
const MemPalSpriteAddr = MemPalAddr + 0x0200

// Memory handles

var MemPalBgPtr = (*uint16)(unsafe.Pointer(uintptr(MemPalAddr)))
var MemPalBg = (*[256]Color)(unsafe.Pointer(MemPalBgPtr))
var MemPalBankBg = (*[16][16]Color)(unsafe.Pointer(MemPalBgPtr))
var MemPalSpritePtr = (*uint16)(unsafe.Pointer(uintptr(MemPalSpriteAddr)))
var MemPalSprite = (*[256]Color)(unsafe.Pointer(MemPalSpritePtr))
var MemPalBankSprite = (*[16][16]Color)(unsafe.Pointer(MemPalSpritePtr))
