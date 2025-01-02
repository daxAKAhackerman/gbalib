// Contains helpers to interact with the memory-mapped I/O registers

package gbalib

import "unsafe"

// Types

type MemIoType struct {
	RegDisplayControl  *VolatileReg32
	RegDisplayStatus   *VolatileReg16
	RegScanlineCounter *VolatileReg16
}

// Memory maps

const MemIoAddr = 0x04000000
const RegDisplayControlAddr = MemIoAddr + 0x00
const RegDisplayStatusAddr = MemIoAddr + 0x04
const RegScanlineCounterAddr = MemIoAddr + 0x06

// Memory defs

const (
	RegDisplayControlModeMask  = 0x07
	RegDisplayControlModeShift = 0

	RegDisplayControlMode0 = 0x00 << RegDisplayControlModeShift
	RegDisplayControlMode1 = 0x01 << RegDisplayControlModeShift
	RegDisplayControlMode2 = 0x02 << RegDisplayControlModeShift
	RegDisplayControlMode3 = 0x03 << RegDisplayControlModeShift
	RegDisplayControlMode4 = 0x04 << RegDisplayControlModeShift
	RegDisplayControlMode5 = 0x05 << RegDisplayControlModeShift
)

const RegDisplayControlGameBoyColor = 0x08
const RegDisplayControlPageSelect = 0x10
const RegDisplayControlHBlankOAM = 0x20
const RegDisplayControl1DObjectMapping = 0x40
const RegDisplayControlScreenBlank = 0x80

const (
	RegDisplayControlBgMask  = 0x0F00
	RegDisplayControlBgShift = 8

	RegDisplayControlBg0             = 0x01 << RegDisplayControlBgShift
	RegDisplayControlBg1             = 0x02 << RegDisplayControlBgShift
	RegDisplayControlBg2             = 0x04 << RegDisplayControlBgShift
	RegDisplayControlBg3             = 0x08 << RegDisplayControlBgShift
	RegDisplayControlObjectRendering = 0x10 << RegDisplayControlBgShift
)

const (
	RegDisplayControlWindowMask  = 0xE000
	RegDisplayControlWindowShift = 13

	RegDisplayControlWindow0      = 0x01 << RegDisplayControlWindowShift
	RegDisplayControlWindow1      = 0x02 << RegDisplayControlWindowShift
	RegDisplayControlObjectWindow = 0x04 << RegDisplayControlWindowShift
)

const (
	RegDisplayStatusVBlank        = 0x01
	RegDisplayStatusHBlank        = 0x02
	RegDisplayStatusVCountTrigger = 0x04
	RegDisplayStatusVBlankIRQ     = 0x08
	RegDisplayStatusHBlankIRQ     = 0x10
	RegDisplayStatusVCountIRQ     = 0x20
)

const (
	RegDisplayStatusVCountTriggerMask  = 0xFF00
	RegDisplayStatusVCountTriggerShift = 8
)

// Memory handles

var MemIo = MemIoType{
	RegDisplayControl:  (*VolatileReg32)(unsafe.Pointer(uintptr(RegDisplayControlAddr))),
	RegDisplayStatus:   (*VolatileReg16)(unsafe.Pointer(uintptr(RegDisplayStatusAddr))),
	RegScanlineCounter: (*VolatileReg16)(unsafe.Pointer(uintptr(RegScanlineCounterAddr))),
}
