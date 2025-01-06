// Contains helpers to interact with the memory-mapped I/O registers

package gbalib

import "unsafe"

// Types

type Key uint16

type MemIoType struct {
	RegDisplayControl  *VolatileReg32
	RegDisplayStatus   *VolatileReg16
	RegScanlineCounter *VolatileReg16
	RegKeyInput        *VolatileReg16
}

// Memory maps

const MemIoAddr = 0x04000000
const RegDisplayControlAddr = MemIoAddr + 0x00
const RegDisplayStatusAddr = MemIoAddr + 0x04
const RegScanlineCounterAddr = MemIoAddr + 0x06
const RegKeyInputAddr = MemIoAddr + 0x0130

// Memory defs

const (
	DisplayControlModeMask  = 0x07
	DisplayControlModeShift = 0

	DisplayControlMode0 = 0x00 << DisplayControlModeShift
	DisplayControlMode1 = 0x01 << DisplayControlModeShift
	DisplayControlMode2 = 0x02 << DisplayControlModeShift
	DisplayControlMode3 = 0x03 << DisplayControlModeShift
	DisplayControlMode4 = 0x04 << DisplayControlModeShift
	DisplayControlMode5 = 0x05 << DisplayControlModeShift
)

const DisplayControlGameBoyColor = 0x08
const DisplayControlPageSelect = 0x10
const DisplayControlHBlankOAM = 0x20
const DisplayControl1DObjectMapping = 0x40
const DisplayControlScreenBlank = 0x80

const (
	DisplayControlBgMask  = 0x0F00
	DisplayControlBgShift = 8

	DisplayControlBg0             = 0x01 << DisplayControlBgShift
	DisplayControlBg1             = 0x02 << DisplayControlBgShift
	DisplayControlBg2             = 0x04 << DisplayControlBgShift
	DisplayControlBg3             = 0x08 << DisplayControlBgShift
	DisplayControlObjectRendering = 0x10 << DisplayControlBgShift
)

const (
	DisplayControlWindowMask  = 0xE000
	DisplayControlWindowShift = 13

	DisplayControlWindow0      = 0x01 << DisplayControlWindowShift
	DisplayControlWindow1      = 0x02 << DisplayControlWindowShift
	DisplayControlObjectWindow = 0x04 << DisplayControlWindowShift
)

const (
	DisplayStatusVBlank        = 0x01
	DisplayStatusHBlank        = 0x02
	DisplayStatusVCountTrigger = 0x04
	DisplayStatusVBlankIRQ     = 0x08
	DisplayStatusHBlankIRQ     = 0x10
	DisplayStatusVCountIRQ     = 0x20
)

const (
	DisplayStatusVCountTriggerMask  = 0xFF00
	DisplayStatusVCountTriggerShift = 8
)

const (
	KeyInputA      = 0x01
	KeyInputB      = 0x02
	KeyInputSelect = 0x04
	KeyInputStart  = 0x08
	KeyInputRight  = 0x10
	KeyInputLeft   = 0x20
	KeyInputUp     = 0x40
	KeyInputDown   = 0x80
	KeyInputR      = 0x0100
	KeyInputL      = 0x0200
)

// Memory handles

var MemIo = MemIoType{
	RegDisplayControl:  (*VolatileReg32)(unsafe.Pointer(uintptr(RegDisplayControlAddr))),
	RegDisplayStatus:   (*VolatileReg16)(unsafe.Pointer(uintptr(RegDisplayStatusAddr))),
	RegScanlineCounter: (*VolatileReg16)(unsafe.Pointer(uintptr(RegScanlineCounterAddr))),
	RegKeyInput:        (*VolatileReg16)(unsafe.Pointer(uintptr(RegKeyInputAddr))),
}

// Helpers

func KeyDownNow(key Key) bool {
	return (^MemIo.RegKeyInput.Get())&uint16(key) != 0
}
