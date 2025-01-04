// Contains helpers to interact 16bits bitmaps

package gbalib

func Bmp16Plot(x, y, w int32, c uint16, d *[]uint16) {
	(*d)[y*w+x] = c
}

func Bmp16HLine(x1, y, x2, w int32, c uint16, d *[]uint16) {
	// Normalize
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Draw
	MemSet16(&(*d)[y*w+x1], c, uint32(x2-x1))
}

func Bmp16VLine(x, y1, y2, w int32, c uint16, d *[]uint16) {
	// Normalize
	if y2 < y1 {
		y1, y2 = y2, y1
	}

	height := y2 - y1

	// Draw
	for i := y1*w + x; height > 0; height-- {
		(*d)[i] = c
		i += w
	}
}

func Bmp16Line(x1, y1, x2, y2, w int32, c uint16, d *[]uint16) {
	var dx, dy, xstep, ystep, dd int32

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

	// Draw
	switch {
	case dy == 0: // Horizontal
		for i := int32(0); i < dx; i++ {
			(*d)[y1*w+i*xstep] = c
		}
	case dx == 0: // Vertical
		for i := int32(0); i < dy; i++ {
			(*d)[(y1+i)*w] = c
		}
	case dx >= dy: // Diagonal, slope <= 1
		dd = 2*dy - dx

		ii := y1*w + x1
		for i := int32(0); i <= dx; i++ {
			(*d)[ii] = c
			if dd >= 0 {
				dd -= 2 * dx
				ii += ystep
			}

			dd += 2 * dy
			ii += xstep
		}
	default: // Diagonal, slope > 1
		dd = 2*dx - dy

		ii := y1*w + x1
		for i := int32(0); i <= dy; i++ {
			(*d)[ii] = c
			if dd >= 0 {
				dd -= 2 * dy
				ii += xstep
			}

			dd += 2 * dx
			ii += ystep
		}
	}
}

func Bmp16Rect(left, top, right, bottom, w int32, c uint16, d *[]uint16) {
	// Normalize
	if right < left {
		left, right = right, left
	}

	if bottom < top {
		top, bottom = bottom, top
	}

	width, height := right-left, bottom-top

	// Draw
	for i := top; height > 0; height-- {
		MemSet16(&(*d)[i*w+left], c, uint32(width))
		i++
	}
}

func Bmp16Frame(left, top, right, bottom, w int32, c uint16, d *[]uint16) {
	// Normalize
	if right < left {
		left, right = right, left
	}

	if bottom < top {
		top, bottom = bottom, top
	}

	width, height := right-left+1, bottom-top

	// Draw
	// Top line
	MemSet16(&(*d)[top*w+left], c, uint32(width-1))

	if height < 2 {
		return
	}

	top++
	height--

	// Left and right lines
	for i := top; height > 0; height-- {
		(*d)[i*w+left] = c
		(*d)[i*w+right-1] = c
		i++
	}

	// Bottom line
	MemSet16(&(*d)[(bottom-1)*w+left], c, uint32(width-1))
}
