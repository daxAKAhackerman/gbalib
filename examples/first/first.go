package main

import "unsafe"

func main() {
	*(*uint32)(unsafe.Pointer((uintptr(0x04000000)))) = 0x0403

	*(*uint16)(unsafe.Pointer(uintptr(0x06000000 + (120+80*240)*2))) = 0x001f
	*(*uint16)(unsafe.Pointer(uintptr(0x06000000 + (136+80*240)*2))) = 0x03e0
	*(*uint16)(unsafe.Pointer(uintptr(0x06000000 + (120+96*240)*2))) = 0x7c00

	for {
	}
}
