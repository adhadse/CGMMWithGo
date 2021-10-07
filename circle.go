package main

import (
	rl "github.com/chunqian/go-raylib/raylib"
)


func DrawCircle(xc, yc, x, y int32) {
	rl.DrawPixel(xc+x, yc+y, rl.Black)
	rl.DrawPixel(xc-x, yc+y, rl.Black)
	rl.DrawPixel(xc+x, yc-y, rl.Black)
	rl.DrawPixel(xc-x, yc-y, rl.Black)
	rl.DrawPixel(xc+y, yc+x, rl.Black)
	rl.DrawPixel(xc-y, yc+x, rl.Black)
	rl.DrawPixel(xc+y, yc-x, rl.Black)
	rl.DrawPixel(xc-y, yc-x, rl.Black)
}

// CircleBresenham Function for circle-generation
// using Bresenham's algorithm
func CircleBresenham(xc, yc, r int32) {
	var x, y int32 = 0, r
	d := 3 - 2 * r
	DrawCircle(xc, yc, x, y)

	for x <= y {
		// for each pixel we will
		// draw all eight pixels
		x++
		// check for decision parameter
		// and correspondingly
		// update d, x, y
		if d > 0 {
			y--
			d = d + 4 * (x - y) + 10
		} else {
			d = d + 4*x + 6
		}
		DrawCircle(xc, yc, x, y)
	}
}
