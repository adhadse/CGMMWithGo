package main

import (
	"fmt"
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
	var steps float32 = 0.0

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
		fmt.Println("Point", i, "X:", x, " | Y:", y)
	}
}

func OldBresenham(x1, y1, x2, y2 int32) {
	mNew := 2*(y2-y1)
	slopeErrorNew := mNew - (x2 - x1)
	for x, y:=x1, y1; x1<=x2; x++ {
		fmt.Println("Point", "X:", x, " | Y:", y)
		slopeErrorNew += mNew               // add slope to increment angle formed
		rl.DrawPixel(x, y, rl.Black)
		if slopeErrorNew >= 0 {             // if slope error reaches limit
			y++                          	// increment y, and update slope error
			slopeErrorNew -= 2*(x2-x1)
		}
	}
}

func Bresenham(x1, y1, x2, y2 int32) {
	dx := x2 - x1
	dy := y2 - y1
	x := x1
	y := y1
	p := 2*dy-dx
	for x < x2 {
		fmt.Println("Point", "X", x, " | Y:", y)
		if p >= 0 {
			rl.DrawPixel(x, y, rl.Black)
			y += 1
			p += 2*dy-2*dx
		} else {
			rl.DrawPixel(x, y, rl.Black)
			p += 2*dy
		}
		x += 1
	}
}
