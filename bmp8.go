// Contains helpers to interact 8bits bitmaps

package gbalib

func Bmp8Plot(x, y, w int32, cid uint8, d *[]uint16) {
	v := (*d)[(y*w+x)/2]
	if x%2 == 0 {
		(*d)[(y*w+x)/2] = v&0xFF00 | uint16(cid)&0xFF
	} else {
		(*d)[(y*w+x)/2] = v&0xFF | uint16(cid)<<8
	}
}
