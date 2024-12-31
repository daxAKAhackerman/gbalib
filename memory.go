// Contains helpers to interact with any part of the memory

package gbalib

import "runtime/volatile"

type VolatileReg32 struct {
	volatile.Register32
}

type Reg32 uint32

func (r *VolatileReg32) Init(v ...uint32) {
	var newRegValue uint32

	for _, value := range v {
		newRegValue |= value
	}

	r.Set(newRegValue)
}

func (r *Reg32) Init(v ...uint32) {
	var newRegValue uint32

	for _, value := range v {
		newRegValue |= value
	}

	*r = Reg32(newRegValue)
}

func (r *VolatileReg32) Update(v ...uint32) {
	var newRegValue uint32

	for _, value := range v {
		newRegValue |= value
	}

	r.SetBits(newRegValue)
}

func (r *Reg32) Update(v ...uint32) {
	var newRegValue uint32

	for _, value := range v {
		newRegValue |= value
	}

	*r |= Reg32(newRegValue)
}

func (r *VolatileReg32) Clear() {
	r.Set(0)
}

func (r *Reg32) Clear() {
	*r = Reg32(0)
}

func (r *VolatileReg32) Read() uint32 {
	return r.Get()
}

func (r *Reg32) Read() uint32 {
	return uint32(*r)
}

func (r *VolatileReg32) SetField(value, mask uint32, shift uint8) {
	r.ReplaceBits(value, mask, shift)
}

func (r *Reg32) SetField(value, mask uint32, shift uint8) {
	*r = Reg32((uint32(*r) &^ mask) | ((value << shift) & mask))
}

func (r *VolatileReg32) UpdateField(value, mask, shift uint32) {
	r.SetBits((value << shift) & mask)
}

func (r *Reg32) UpdateField(value, mask, shift uint32) {
	*r = Reg32(uint32(*r) | ((value << shift) & mask))
}

func (r *VolatileReg32) GetField(mask, shift uint32) uint32 {
	return (r.Get() & mask) >> shift
}

func (r *Reg32) GetField(mask, shift uint32) uint32 {
	return (uint32(*r) & mask) >> shift
}

func (r *VolatileReg32) ClearField(mask uint32) {
	r.Set(r.Get() &^ mask)
}

func (r *Reg32) ClearField(mask uint32) {
	*r = Reg32(uint32(*r) &^ mask)
}
