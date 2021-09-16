package main

import (
	rl "github.com/chunqian/go-raylib/raylib"
)


type LineInt struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func DDA(x0, y0, x1, y1 float32) {
	dx := x1 - x0
	dy := y1 - y0
	var steps float32

	if absInt(dx) > absInt(dy) {
		steps = absInt(dx)
	} else {
		steps = absInt(dy)
	}

	// calculate increment in x & y for each steps
	var xinc = dx / steps
	var yinc = dy / steps

	// Put pixel for each step
	x := x0
	y := y0
	var i float32 = 0.0
	for ;i <= steps; i++ {
		rl.DrawPixel(int32(x), int32(y), rl.Black)
		x += xinc
		y += yinc
	}
}

func Bresenham(x1, y1, x2, y2 int32) {
	dx := x2 - x1
	dy := y2 - y1
	x := x1
	y := y1
	p := 2*dy-dx                            // Decision Parameter P_k
	for x < x2 {
		if p >= 0 {
			y += 1
			p += 2*dy-2*dx
		} else {
			p += 2*dy
		}
		rl.DrawPixel(x, y, rl.Black)
		x += 1
	}
}
