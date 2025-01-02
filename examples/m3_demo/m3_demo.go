package main

import "gbalib"

func main() {
	var i, j int32

	gbalib.MemIo.RegDisplayControl.Init(gbalib.RegDisplayControlMode3, gbalib.RegDisplayControlBg2)

	// Fill screen with gray color
	gbalib.M3Fill(gbalib.MakeRgb15(12, 12, 14))

	// Rectangles:
	gbalib.M3Rect(12, 8, 108, 72, gbalib.ClrRed)
	gbalib.M3Rect(108, 72, 132, 88, gbalib.ClrLime)
	gbalib.M3Rect(132, 88, 228, 152, gbalib.ClrBlue)

	// Rectangle frames
	gbalib.M3Frame(132, 8, 228, 72, gbalib.ClrCyan)
	gbalib.M3Frame(109, 73, 131, 87, gbalib.ClrBlack)
	gbalib.M3Frame(12, 88, 108, 152, gbalib.ClrYellow)

	// Lines in top right frame
	for i = 0; i <= 8; i++ {
		j = 3*i + 7
		gbalib.M3Line(132+11*i, 9, 226, 12+7*i, gbalib.MakeRgb15(uint32(j), 0, uint32(j)))
		gbalib.M3Line(226-11*i, 70, 133, 69-7*i, gbalib.MakeRgb15(uint32(j), 0, uint32(j)))
	}

	// Lines in bottom left frame
	for i = 0; i <= 8; i++ {
		j = 3*i + 7
		gbalib.M3Line(15+11*i, 88, 104-11*i, 150, gbalib.MakeRgb15(0, uint32(j), uint32(j)))
	}

	for {
	}
}
