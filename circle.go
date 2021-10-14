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
	p := 3 - 2 * r
	DrawCircle(xc, yc, x, y)

	for x <= y {
		// for each pixel we will
		// draw all eight pixels
		x++
		// check for decision parameter
		// and correspondingly
		// update d, x, y
		if p > 0 {
			y--
			p = p + 4*(x - y) + 10
		} else {
			p = p + 4*x + 6
		}
		DrawCircle(xc, yc, x, y)
	}
}


func CircleMidPoint(xc, yc, r int32) {
	var x, y int32 = 0, r

	p := 1 - r
	for x <= y {
		// for each pixel we will
		// draw all eight pixels
		x++
		// check for decision parameter
		// and correspondingly
		// update d, x, y
		if p >= 0 {
			// Mid-point outside the perimeter
			y--
			p = p + 2*x - 2*y + 1
		} else {
			// Mid-point inside or on the perimeter
			p = p + 2*x + 1
		}
		DrawCircle(xc, yc, x, y)
		// All the perimeter points have already been printed
		if x > y {
			break
		}
	}
}
