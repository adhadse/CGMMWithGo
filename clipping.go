package main

import (
	"fmt"
)

// Defining region codes
const (
	inside = 0   //0000
	left = 1     //0001
	right = 2    //0010
	bottom = 4   //0100
	top = 8      //1000
)

// Function to compute region code for a point(x, y)
func computeRegionCode(x, y int) int {
	 code := inside
	 if x < xMin {        // to the left of rectangle
		 code |= left
	 } else if x > xMax { // to the right of rectangle
		 code |= right
	 }
	 if y < yMin {        // below the rectangle
		 code |= bottom
 	 } else if y > yMax { // above the rectangle
		  code |= top
	 }
	 return code
}

func CohenSutherlandClip(l LineInt) {
	code1, code2 := computeRegionCode(l.x1, l.y1), computeRegionCode(l.x2, l.y2)

	for true {
		if code1 == 0 && code2 == 0 {                  // If both endpoints lie within rectangle
			fmt.Printf("Line accepted (%d,%d) and (%d,%d)\n", l.x1, l.y1, l.x2, l.y2)
			break
		} else if itob(code1 & code2) {                // If both endpoints lie outside rectangle
			fmt.Printf("Line rejected (%d,%d) and (%d,%d)\n", l.x1, l.y1, l.x2, l.y2)
			break
		} else {
			// If Some segment of line lies within the
			// rectangle
			var codeOut int
			if code1 != 0 {                            // At least one endpoint is outside the
				codeOut = code1 // rectangle, pick it.
			} else {
				codeOut = code2
			}

			// Find intersection point
			// x = x1 + (1 / slope) * (y_<yMin/max> - y1)
			// y = y1 + slope * (x_<yMin/max> - x1)
			var x, y int
			if itob(codeOut & top) { // point is above the clip rectangle
				x = l.x1 + (l.x2 - l.x1) * (yMax- l.y1) / (l.y2 - l.y1)
				y = yMax
			} else if itob(codeOut & bottom) { // point is below the rectangle
				x = l.x1 + (l.x2 - l.x1) * (yMin- l.y1) / (l.y2 - l.y1)
				y = yMin
			}
			if itob(codeOut & right) { // point is to the right of rectangle
				x = xMax
				y = l.y1 + (l.y2 - l.y1) * (xMax- l.x1) / (l.x2 - l.x1)
			} else if itob(codeOut & left) { // point is to the left of the rectangle
				x = xMin
				y = l.y1 + (l.y2 - l.y1) * (xMin- l.x1) / (l.x2 - l.x1)
			}

			// Replace point outside rectangle by intersection point
			if codeOut == code1 {
				l.x1, l.y1 = x, y
			} else {
				l.x2, l.y2 = x, y
			}
		}
	}
}
