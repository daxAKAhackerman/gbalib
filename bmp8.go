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
		volatile.StoreUint16(&(*d)[(y*w+x1)/2], (*d)[(y*w+x1)/2]&0xFF|uint16(cid)<<8)
		width--
		x1++
	}

	// Right unaligned pixel
	if width%2 != 0 {
		volatile.StoreUint16(&(*d)[(y*w+x2)/2], (*d)[(y*w+x2)/2]&0xFF00|uint16(cid))
		width--
		x2--
	}

	width /= 2

	// Aligned line
	if width > 0 {
		MemSet16(&(*d)[(y*w+x1)/2], uint16(cid)|uint16(cid)<<8, uint32(width))
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
	var dx, dy, xstep, ystep, dd int32

	if y1 == y2 { // Horizontal
		Bmp8HLine(x1, y1, x2, w, cid, d)
		return
	} else if x1 == x2 { // Vertical
		Bmp8VLine(x1, y1, y2, w, cid, d)
		return
	}

	// Normalize
	if x1 > x2 {
		xstep = -1
		dx = x1 - x2
	} else {
		xstep = +1
		dx = x2 - x1
	}

	if y1 > y2 {
		ystep = -w
		dy = y1 - y2
	} else {
		ystep = +w
		dy = y2 - y1
	}

	mask := uint16(255)

	cid16 := uint16(cid) | uint16(cid)<<8
	if x1%2 != 0 {
		mask = ^mask
	}

	if dx >= dy { // Diagonal, slope <= 1
		dd = 2*dy - dx

		ii := y1*w + x1
		for i := int32(0); i <= dx; i++ {
			volatile.StoreUint16(&(*d)[ii/2], ((*d)[ii/2]&^mask)|(uint16(cid16)&mask))

			if dd >= 0 {
				dd -= 2 * dx
				ii += ystep
			}

			dd += 2 * dy
			ii += xstep
			mask = ^mask
		}
	} else { // Diagonal, slope > 1
		dd = 2*dx - dy

		ii := y1*w + x1
		for i := int32(0); i <= dy; i++ {
			volatile.StoreUint16(&(*d)[ii/2], ((*d)[ii/2]&^mask)|(uint16(cid16)&mask))

			if dd >= 0 {
				dd -= 2 * dy
				ii += xstep
				mask = ^mask
			}

			dd += 2 * dx
			ii += ystep
		}
	}
}

func Bmp8Rect(left, top, right, bottom, w int32, cid uint8, d *[]uint16) {
	// Normalize
	if right < left {
		left, right = right, left
	}

	if bottom < top {
		top, bottom = bottom, top
	}

	width, height := right-left, bottom-top

	if width == 0 {
		return
	}

	// Unaligned left
	if left%2 != 0 {
		theight := height
		shiftedCid := uint16(cid) << 8
		for i := (top*w + left) / 2; theight > 0; theight-- {
			volatile.StoreUint16(&(*d)[i], (*d)[i]&0xFF|shiftedCid)
			i += w / 2
		}
		width--
		left++
	}

	// Unaligned right
	if width%2 != 0 {
		theight := height
		for i := (top*w + right) / 2; theight > 0; theight-- {
			volatile.StoreUint16(&(*d)[i], (*d)[i]&0xFF00|uint16(cid))
			i += w / 2
		}
		width--
	}

	// Center
	for i := top; height > 0; height-- {
		MemSet16(&(*d)[(i*w+left)/2], uint16(cid)|uint16(cid)<<8, uint32(width/2))
		i++
	}
}

func Bmp8Frame(left, top, right, bottom, w int32, cid uint8, d *[]uint16) {
	// Normalize
	if right < left {
		left, right = right, left
	}

	if bottom < top {
		top, bottom = bottom, top
	}

	// Draw
	Bmp8HLine(left, top, right, w, cid, d)
	Bmp8HLine(left, bottom-1, right, w, cid, d)

	Bmp8VLine(left, top+1, bottom, w, cid, d)
	Bmp8VLine(right-1, top+1, bottom, w, cid, d)
}
