package main

import (
	"gbalib"
)

func main() {
	gbalib.MemIo.RegDisplayControl.Init(
		gbalib.RegDisplayControlMode3,
		gbalib.RegDisplayControlBg2,
	)

	gbalib.M3Plot(120, 80, gbalib.MakeRgb15(31, 0, 0))
	gbalib.M3Plot(136, 80, gbalib.MakeRgb15(0, 31, 0))
	gbalib.M3Plot(120, 96, gbalib.MakeRgb15(0, 0, 31))
	for {
	}
}
