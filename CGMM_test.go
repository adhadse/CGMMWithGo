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
	x1 := 2
	y1 := 2
	x2 := 430
	y2 := 430

	screenWidth := int32(800)
	screenHeight := int32(600)
	rl.InitWindow(screenWidth, screenHeight, "DDA algorithm | Computer Graphics Demo algorithms | " +
		fmt.Sprintf("Line Draw between (%d, %d) and (%d, %d)", x1, y1, x2, y2))
	defer rl.CloseWindow()
	winImage := rl.LoadImage("res/golang-48.png")
	rl.SetWindowIcon(winImage)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		DDA(x1, y1, x2, y2)
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
