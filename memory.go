// Contains helpers to interact with any part of the memory

package gbalib

import "runtime/volatile"

type VolatileReg16 struct {
	volatile.Register16
}

type VolatileReg32 struct {
	volatile.Register32
}

func (r *VolatileReg16) Init(v ...uint16) {
	var newRegValue uint16

	for _, value := range v {
		newRegValue |= value
	}

	r.Set(newRegValue)
}

func (r *VolatileReg32) Init(v ...uint32) {
	var newRegValue uint32

	for _, value := range v {
		newRegValue |= value
	}

	r.Set(newRegValue)
}

func (r *VolatileReg16) Update(v ...uint16) {
	var newRegValue uint16

	for _, value := range v {
		newRegValue |= value
	}

	r.SetBits(newRegValue)
}

func (r *VolatileReg32) Update(v ...uint32) {
	var newRegValue uint32

	for _, value := range v {
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
