package CGMM_Practical_programs_in_Go

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
	"math"
)

func DDA(x0, y0, x1, y1 float64) {
	dx := x1 - x0
	dy := y1 - y0
	steps := 0.0
	dc := gg.NewContext(1000, 1000)
	dc.SetColor(color.White)

	if math.Abs(dx) > math.Abs(dy) {
		steps = math.Abs(dx)
	} else {
		steps = math.Abs(dy)
	}

	// calculate increment in x & y for each steps
	xinc := dx / steps
	yinc := dy / steps

	// Put pixel for each step
	x := x0
	y := y0
	for i := 0.0; i <= steps; i++ {
		dc.SetPixel(int(math.Round(x)), int(math.Round(y)))
		x += xinc
		y += yinc
		fmt.Println("Point", math.Round(i), "X:", x, " | Y:", y)
	}
	dc.SavePNG("output/out_DDA.png")
}

func Bresenham(x1, y1, x2, y2 int) {
	dc := gg.NewContext(1000, 1000)
	dc.SetColor(color.White)
	mNew := 2*(y2-y1)
	slopeErrorNew := mNew - (x2 - x1)
	for x, y:=x1, y1; x1<=x2; x++ {
		fmt.Println("Point", "X", x, " | Y:", y)
		slopeErrorNew += mNew // add slope to increment angle formed
		dc.SetPixel(x, y)
		if slopeErrorNew >= 0 { // if slope error reaches limit
			y++                          	// increment y, and update slope error
			slopeErrorNew -= 2*(x2-x1)
		}
	}
	dc.SavePNG("output/out_bresenham.png")
}



