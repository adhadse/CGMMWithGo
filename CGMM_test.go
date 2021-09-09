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

func Test_DDA(t *testing.T) {
	var x1 float32 = 2
	var y1 float32 = 3
	var x2 float32 = 231
	var y2 float32 = 344

	screenWidth := int32(800)
	screenHeight := int32(600)
	rl.InitWindow(screenWidth, screenHeight, "DDA algorithm | Computer Graphics Demo algorithms | " +
		fmt.Sprintf("Line Draw between (%.2f, %.2f) and (%.2f, %.2f)", x1, y1, x2, y2))
	defer rl.CloseWindow()
	winImage := rl.LoadImage("res/golang-48.png")
	rl.SetWindowIcon(winImage)


	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		DDA(x1, y1, x2, y2)

		if rl.IsMouseButtonDown(0) {
			x1, y1 = float32(rl.GetMouseX()), float32(rl.GetMouseY())
			if rl.IsMouseButtonReleased(0) {
				x2, y2 = float32(rl.GetMouseX()), float32(rl.GetMouseY())
			}
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}


func Test_Bresenham(t *testing.T) {
	var x1 int32 = 2
	var y1 int32 = 3
	var x2 int32 = 231
	var y2 int32 = 500

	screenWidth := int32(800)
	screenHeight := int32(600)
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
