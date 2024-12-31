// Contains helpers to interact with the memory-mapped I/O registers

package gbalib

import (
	"unsafe"
)

// Types

type MemIoType struct {
	RegDisplayControl *VolatileReg32
}

// Memory maps

const MemIoAddr = 0x04000000
const RegDisplayControlAddr = MemIoAddr + 0x0000

// Memory defs

const (
	RegDisplayControlModeMask  = 0x00000007
	RegDisplayControlModeShift = 0

	RegDisplayControlMode0 = 0x00 << RegDisplayControlModeShift
	RegDisplayControlMode1 = 0x01 << RegDisplayControlModeShift
	RegDisplayControlMode2 = 0x02 << RegDisplayControlModeShift
	RegDisplayControlMode3 = 0x03 << RegDisplayControlModeShift
	RegDisplayControlMode4 = 0x04 << RegDisplayControlModeShift
	RegDisplayControlMode5 = 0x05 << RegDisplayControlModeShift
)

const (
	RegDisplayControlBgMask  = 0x00000F00
	RegDisplayControlBgShift = 8

	RegDisplayControlBg0 = 0x01 << RegDisplayControlBgShift
	RegDisplayControlBg1 = 0x02 << RegDisplayControlBgShift
	RegDisplayControlBg2 = 0x04 << RegDisplayControlBgShift
	RegDisplayControlBg3 = 0x08 << RegDisplayControlBgShift
)

var MemIo = MemIoType{
	RegDisplayControl: (*VolatileReg32)(unsafe.Pointer(uintptr(RegDisplayControlAddr))),
}
