// Contains helpers to interact with any part of the memory

package gbalib

type Reg32 uint32

func (r *Reg32) Init(v ...uint32) {
	var newRegValue uint32

	for _, value := range v {
		newRegValue |= value
	}

	*r = Reg32(newRegValue)
}

func (r *Reg32) Update(v ...uint32) {
	var newRegValue uint32

	for _, value := range v {
		newRegValue |= value
	}

	*r |= Reg32(newRegValue)
}

func (r *Reg32) Clear() {
	*r = 0
}

func (r *Reg32) Read() uint32 {
	return uint32(*r)
}

func (r *Reg32) SetField(value, mask, shift uint32) {
	*r = Reg32((r.Read() &^ mask) | ((value << shift) & mask))
}

func (r *Reg32) UpdateField(value, mask, shift uint32) {
	*r = Reg32(r.Read() | ((value << shift) & mask))
}

func (r *Reg32) GetField(mask, shift uint32) uint32 {
	return (r.Read() & mask) >> shift
}

func (r *Reg32) ClearField(mask uint32) {
	*r = Reg32(r.Read() &^ mask)
}
