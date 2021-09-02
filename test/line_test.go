package test

import (
	"CGMM_Practical_programs_in_Go"
	"testing"
)

func Test_DDA(t *testing.T) {
	x0 := 2.0
	y0 := 2.0
	x1 := 16.0
	y1 := 16.0
	CGMM_Practical_programs_in_Go.DDA(x0, y0, x1, y1)
}

func Test_Bresenham(t *testing.T) {
	x0 := 2
	y0 := 3
	x1 := 6
	y1 := 15
	CGMM_Practical_programs_in_Go.Bresenham(x0, y0, x1, y1)
}
