// Contains helpers to interact 8bits bitmaps

package gbalib

import "runtime/volatile"

func Bmp8Plot(x, y, w int32, cid uint8, d *[]uint16) {
	if x%2 == 0 {
		volatile.StoreUint16(&(*d)[(y*w+x)/2], (*d)[(y*w+x)/2]&0xFF00|uint16(cid))
	} else {
		volatile.StoreUint16(&(*d)[(y*w+x)/2], (*d)[(y*w+x)/2]&0xFF|uint16(cid)<<8)
	}
}

func Bmp8HLine(x1, y, x2, w int32, cid uint8, d *[]uint16) {
	// Normalize
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	width := x2 - x1

	if width == 0 {
		return
	}

	// Draw
	// Left unaligned pixel
	if x1%2 != 0 {
		volatile.StoreUint16(&(*d)[y*w+x1/2], (*d)[y*w+x1/2]&0xFF|uint16(cid)<<8)
		width--
		x1++
	}

	// Right unaligned pixel
	if width%2 != 0 {
		volatile.StoreUint16(&(*d)[y*w+x2/2], (*d)[y*w+x2/2]&0xFF00|uint16(cid))
		width--
		x2--
	}

	width /= 2

	// Aligned line
	if width > 0 {
		MemSet16(&(*d)[y*w+x1/2], uint16(cid)|uint16(cid)<<8, uint32(width))
	}
}

func Bmp8VLine(x, y1, y2, w int32, cid uint8, d *[]uint16) {
	// Normalize
	if y2 < y1 {
		y1, y2 = y2, y1
	}

	height := y2 - y1

	// Draw
	if x%2 == 0 {
		for i := (y1*w + x) / 2; height > 0; height-- {
			volatile.StoreUint16(&(*d)[i], (*d)[i]&0xFF00|uint16(cid))
			i += w / 2
		}
	} else {
		shiftedCid := uint16(cid) << 8
		for i := (y1*w + x) / 2; height > 0; height-- {
			volatile.StoreUint16(&(*d)[i], (*d)[i]&0xFF|shiftedCid)
			i += w / 2
		}
	}
}

func Bmp8Line(x1, y1, x2, y2, w int32, cid uint8, d *[]uint16) {
}

func Bmp8Rect(left, top, right, bottom, w int32, cid uint8, d *[]uint16) {
}

func Bmp8Frame(left, top, right, bottom, w int32, cid uint8, d *[]uint16) {
}
