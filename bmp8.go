// Contains helpers to interact 8bits bitmaps

package gbalib

func Bmp8Plot(x, y, w int32, cid uint8, d *[]uint16) {
	if x&0x01 == 1 {
		(*d)[(y*w+x)/2] = (*d)[(y*w+x)/2]&0xFF00 | uint16(cid)
	} else {
		(*d)[(y*w+x)/2] = (*d)[(y*w+x)/2]&0x00FF | uint16(cid)<<8
	}
}
