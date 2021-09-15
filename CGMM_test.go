package main

import (
	"fmt"
	rl "github.com/chunqian/go-raylib/raylib"
	"runtime"
	"testing"
)

func init() {
	runtime.LockOSThread()
}

func TestDDA(t *testing.T) {
	var x1 float32 = 2
	var y1 float32 = 3
	var x2 float32 = 231
	var y2 float32 = 344

	screenWidth := int32(xMax)
	screenHeight := int32(yMax)
	rl.InitWindow(screenWidth, screenHeight, "DDA algorithm | Computer Graphics Demo algorithms")
	defer rl.CloseWindow()
	winImage := rl.LoadImage("res/golang-48.png")
	rl.SetWindowIcon(winImage)


	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		DDA(x1, y1, x2, y2)
		if rl.IsMouseButtonDown(0) {
			x1, y1 = float32(rl.GetMouseX()), float32(rl.GetMouseY())
		}
		if rl.IsMouseButtonPressed(0) {
			x2, y2 = float32(rl.GetMouseX()), float32(rl.GetMouseY())
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}


func TestBresenham(t *testing.T) {
	var x1 int32 = 2
	var y1 int32 = 3
	var x2 int32 = 231
	var y2 int32 = 500

	screenWidth := int32(xMax)
	screenHeight := int32(yMax)
	rl.InitWindow(screenWidth, screenHeight, "Bresenham algorithm | Computer Graphics Demo algorithms | " +
		fmt.Sprintf("Line Draw between (%d, %d) and (%d, %d)", x1, y1, x2, y2))
	defer rl.CloseWindow()
	winImage := rl.LoadImage("res/golang-48.png")
	rl.SetWindowIcon(winImage)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		Bresenham(x1, y1, x2, y2)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func TestCohenSutherlandClip(T *testing.T) {
	l := LineInt{3, 3, 500, 440}
	CohenSutherlandClip(l)
	fmt.Printf("Region code for Point 1 (%d, %d) = %04b \n", l.x1, l.y1, computeRegionCode(l.x1, l.y1))
	fmt.Printf("Region code for Point 2 (%d, %d) = %04b \n", l.x2, l.y2, computeRegionCode(l.x2, l.y2))
}
