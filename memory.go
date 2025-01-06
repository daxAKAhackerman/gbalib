// Contains helpers to interact with any part of the memory

package gbalib

import (
	"runtime/volatile"
	"unsafe"
)

// Types

type VolatileReg16 struct {
	volatile.Register16
}

type VolatileReg32 struct {
	volatile.Register32
}

// Helpers

func (r *VolatileReg16) Init(values ...uint16) {
	var newRegValue uint16

	for _, value := range values {
		newRegValue |= value
	}

	r.Set(newRegValue)
}

func (r *VolatileReg32) Init(values ...uint32) {
	var newRegValue uint32

	for _, value := range values {
		newRegValue |= value
	}

	r.Set(newRegValue)
}

func (r *VolatileReg16) Update(values ...uint16) {
	var newRegValue uint16

	for _, value := range values {
		newRegValue |= value
	}

	r.SetBits(newRegValue)
}

func (r *VolatileReg32) Update(values ...uint32) {
	var newRegValue uint32

	for _, value := range values {
		newRegValue |= value
	}

	r.SetBits(newRegValue)
}

func (r *VolatileReg16) Clear() {
	r.Set(0)
}

func (r *VolatileReg32) Clear() {
	r.Set(0)
}

func (r *VolatileReg16) Read() uint16 {
	return r.Get()
}

func (r *VolatileReg32) Read() uint32 {
	return r.Get()
}

func (r *VolatileReg16) SetField(value, mask uint16, shift uint8) {
	r.ReplaceBits(value, mask, shift)
}

func (r *VolatileReg32) SetField(value, mask uint32, shift uint8) {
	r.ReplaceBits(value, mask, shift)
}

func (r *VolatileReg16) UpdateField(value, mask uint16, shift uint8) {
	r.SetBits((value << shift) & mask)
}

func (r *VolatileReg32) UpdateField(value, mask uint32, shift uint8) {
	r.SetBits((value << shift) & mask)
}

func (r *VolatileReg16) GetField(mask, shift uint16) uint16 {
	return (r.Get() & mask) >> shift
}

func (r *VolatileReg32) GetField(mask, shift uint32) uint32 {
	return (r.Get() & mask) >> shift
}

func (r *VolatileReg16) ClearField(mask uint16) {
	r.Set(r.Get() &^ mask)
}

func (r *VolatileReg32) ClearField(mask uint32) {
	r.Set(r.Get() &^ mask)
}

func MemSet16(destination *uint16, value uint16, halfWordCount uint32) {
	if uintptr(unsafe.Pointer(destination))%4 != 0 {
		*destination = value
		destination = (*uint16)(unsafe.Pointer((uintptr(unsafe.Pointer(destination)) + 2)))
		halfWordCount--
	}

	leftover := halfWordCount%2 != 0
	wordCount := halfWordCount / 2

	MemSet32((*uint32)(unsafe.Pointer(destination)), uint32(value)|uint32(value)<<16, wordCount)

	if leftover {
		lastptr := uintptr(unsafe.Pointer(destination)) + uintptr(wordCount)*4
		*(*uint16)(unsafe.Pointer(lastptr)) = value
	}
}

func MemSet32(destination *uint32, value uint32, wordCount uint32) {
	destinationSlice := unsafe.Slice(destination, wordCount)

	for i := uint32(0); i < wordCount; i++ {
		destinationSlice[i] = uint32(value) | uint32(value)<<16
	}
}

func MemCpy16(destination *uint16, source *uint16, halfWordCount uint32) {
	if uintptr(unsafe.Pointer(destination))%4 != 0 {
		*destination = *source
		destination = (*uint16)(unsafe.Pointer((uintptr(unsafe.Pointer(destination)) + 2)))
		source = (*uint16)(unsafe.Pointer((uintptr(unsafe.Pointer(source)) + 2)))
		halfWordCount--
	}

	leftover := halfWordCount%2 != 0
	wordCount := halfWordCount / 2

	MemCpy32((*uint32)(unsafe.Pointer(destination)), (*uint32)(unsafe.Pointer(source)), wordCount)

	if leftover {
		lastdptr := uintptr(unsafe.Pointer(destination)) + uintptr(wordCount)*4
		lastsptr := uintptr(unsafe.Pointer(source)) + uintptr(wordCount)*4
		*(*uint16)(unsafe.Pointer(lastdptr)) = *(*uint16)(unsafe.Pointer(lastsptr))
	}
}

func MemCpy32(destination *uint32, source *uint32, wcount uint32) {
	destinationSlice := unsafe.Slice(destination, wcount)
	sourceSlice := unsafe.Slice(source, wcount)
	for i := uint32(0); i < wcount; i++ {
		destinationSlice[i] = sourceSlice[i]
	}
}
