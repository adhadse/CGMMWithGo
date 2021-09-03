package main

import (
	"fmt"
	rl "github.com/chunqian/go-raylib/raylib"
)


func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func DDA(x0, y0, x1, y1 int) {
	dx := x1 - x0
	dy := y1 - y0
	steps := 0

	if absInt(dx) > absInt(dy) {
		steps = absInt(dx)
	} else {
		steps = absInt(dy)
	}

	// calculate increment in x & y for each steps
	xinc := dx / steps
	yinc := dy / steps

	// Put pixel for each step
	x := x0
	y := y0
	for i := 0; i <= steps; i++ {
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




