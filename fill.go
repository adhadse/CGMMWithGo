package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func FloodFill(x int32, y int32, fillColor rl.Color, oldColor rl.Color, image rl.Image){
	if rl.GetImageColor(image, x, y) == oldColor {
		rl.ImageDrawPixel(&image, x, y, fillColor)
		FloodFill(x+1, y, fillColor, oldColor, image)
		FloodFill(x-1, y, fillColor, oldColor, image)
		FloodFill(x, y+1, fillColor, oldColor, image)
		FloodFill(x, y-1, fillColor, oldColor, image)
	}
}

func BoundaryFill(x int32, y int32, color rl.Color, boundaryColor rl.Color, image rl.Image) {
	c := rl.GetImageColor(image, x, y)

	if c != color && c != boundaryColor {
		rl.ImageDrawPixel(&image, x, y, color)
		BoundaryFill(x+1, y, color, boundaryColor, image)
		BoundaryFill(x-1, y, color, boundaryColor, image)
		BoundaryFill(x, y+1, color, boundaryColor, image)
		BoundaryFill(x, y-1, color, boundaryColor, image)
	}
}