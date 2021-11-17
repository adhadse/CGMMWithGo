package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
	"runtime"
	"testing"
)

func init() {
	runtime.LockOSThread()
}

func TestDrawPoints(t *testing.T) {
	screenWidth := int32(xMax)
	screenHeight := int32(yMax)
	rl.InitWindow(screenWidth, screenHeight, "DDA algorithm | Computer Graphics Demo algorithms")
	defer rl.CloseWindow()

	winImage := rl.LoadImage("res/golang-48.png")
	rl.SetWindowIcon(*winImage)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawPixel(3, 3, rl.Black)
		rl.DrawPixel(10, 10, rl.Blue)
		rl.DrawPixel(15, 15, rl.Brown)
		rl.DrawPixel(20, 20, rl.Beige)
		rl.DrawPixel(25, 25, rl.Pink)
		rl.DrawPixel(30, 30, rl.Green)
		rl.DrawPixel(35, 35, rl.DarkGreen)
		rl.EndDrawing()
	}
	rl.CloseWindow()
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
	rl.SetWindowIcon(*winImage)

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
	rl.InitWindow(screenWidth, screenHeight, "Bresenham algorithm | Computer Graphics Demo algorithms")
	defer rl.CloseWindow()
	winImage := rl.LoadImage("res/golang-48.png")
	rl.SetWindowIcon(*winImage)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		Bresenham(x1, y1, x2, y2)
		if rl.IsMouseButtonDown(0) {
			x1, y1 = rl.GetMouseX(), rl.GetMouseY()
		}
		if rl.IsMouseButtonPressed(0) {
			x2, y2 = rl.GetMouseX(), rl.GetMouseY()
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}


func TestCircleBresenham(t *testing.T) {
	var (
		xc int32 = 50
		yc int32 = 50
		r int32 = 30
	)
	x1, x2, y1, y2 := rl.GetMouseX(), rl.GetMouseX(), rl.GetMouseY(), rl.GetMouseY()

	screenWidth := int32(xMax)
	screenHeight := int32(yMax)
	rl.InitWindow(screenWidth, screenHeight, "Bresenham Circle drawing algorithm | Computer Graphics Demo algorithms")
	defer rl.CloseWindow()
	winImage := rl.LoadImage("res/golang-48.png")
	rl.SetWindowIcon(*winImage)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		CircleBresenham(xc, yc, r)
		if rl.IsMouseButtonPressed(0) {
			x1, y1 = rl.GetMouseX(), rl.GetMouseY()
		}
		if rl.IsMouseButtonDown(0) {
			x2, y2 = rl.GetMouseX(), rl.GetMouseY()
			xc, yc = (x1 + x2) / 2, (y1 + y2) / 2
			r = int32(math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2)))
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func TestCircleMidPoint(t *testing.T) {
	var (
		xc int32 = 50
		yc int32 = 50
		r int32 = 30
	)
	x1, x2, y1, y2 := rl.GetMouseX(), rl.GetMouseX(), rl.GetMouseY(), rl.GetMouseY()

	screenWidth := int32(xMax)
	screenHeight := int32(yMax)
	rl.InitWindow(screenWidth, screenHeight, "Circle Mid-Point algorithm | Computer Graphics Demo algorithms")
	defer rl.CloseWindow()
	winImage := rl.LoadImage("res/golang-48.png")
	rl.SetWindowIcon(*winImage)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		CircleMidPoint(xc, yc, r)
		if rl.IsMouseButtonPressed(0) {
			x1, y1 = rl.GetMouseX(), rl.GetMouseY()
		}
		if rl.IsMouseButtonDown(0) {
			x2, y2 = rl.GetMouseX(), rl.GetMouseY()
			xc, yc = (x1 + x2) / 2, (y1 + y2) / 2
			r = int32(math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2)))
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}


func TestDrawShape(t *testing.T) {
	screenWidth := int32(xMax)
	screenHeight := int32(yMax)
	rl.InitWindow(screenWidth, screenHeight, "Shape drawing program | Computer Graphics Demo algorithms")
	defer rl.CloseWindow()
	winImage := rl.LoadImage("res/golang-48.png")
	rl.SetWindowIcon(*winImage)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawCircleLines(40, 40, 30, rl.Black)
		rl.DrawRectangleLines(90, 90, 40, 30, rl.Black)
		rl.DrawEllipseLines(140, 140, 50, 20, rl.Black)
		rl.DrawLine(200, 200, 334, 442, rl.Black)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}



func TestAnimation(t *testing.T) {
	screenWidth := int32(xMax)
	screenHeight := int32(yMax)

	rl.InitWindow(screenWidth, screenHeight, "Collision detection")
	defer rl.CloseWindow()
	winImage := rl.LoadImage("res/golang-48.png")
	rl.SetWindowIcon(*winImage)

	// Box A: Moving box
	boxA := rl.Rectangle{X: 10, Y: float32(rl.GetScreenHeight() / 2.0) - 50, Width:200, Height:100 }
	var boxASpeedX float32 = 4

	// Box B: Mouse moved box
	boxB := rl.Rectangle{X:float32(rl.GetScreenWidth() / 2.0) - 30, Y:float32(rl.GetScreenHeight()/2.0) - 30,
						 Width:60, Height: 60 }

	boxCollision := rl.Rectangle{} // Collision rectangle

	var screenUpperLimit float32 = 40      // Top menu limits

	pause := false             // Movement pause
	collision := false         // Collision detection

	rl.SetTargetFPS(60)               // Set our game to run at 60 frames-per-second

	for !rl.WindowShouldClose() { // Detect window close button or ESC key
		// Move box if not paused
		if !pause {
			boxA.X += boxASpeedX
		}

		// Bounce box on x screen limits
		if ((boxA.X + boxA.Width) >= float32(rl.GetScreenWidth())) || (boxA.X <= 0) {
			boxASpeedX *= -1
		}

		// Update player-controlled-box (box02)
		boxB.X = float32(rl.GetMouseX()) - boxB.Width/2
		boxB.Y = float32(rl.GetMouseY()) - boxB.Height/2

		// Make sure Box B does not go out of move area limits
		if (boxB.X + boxB.Width) >= float32(rl.GetScreenWidth()) {
			boxB.X = float32(rl.GetScreenWidth()) - boxB.Width
		} else if boxB.X <= 0 {
			boxB.X = 0
		}

		if (boxB.Y + boxB.Height) >= float32(rl.GetScreenHeight()) {
			boxB.Y = float32(rl.GetScreenHeight()) - boxB.Height
		} else if boxB.Y <= screenUpperLimit {
			boxB.Y = screenUpperLimit
		}

		// Check boxes collision
		collision = rl.CheckCollisionRecs(boxA, boxB)

		// Get collision rectangle (only on collision)
		if collision {
			boxCollision = rl.GetCollisionRec(boxA, boxB)
		}

		// Pause Box A movement
		if rl.IsKeyPressed(' ') {
			pause = !pause
		}

		// Draw
		//-----------------------------------------------------
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rec_color := rl.Red
		if collision { rec_color = rl.Red } else { rec_color = rl.Black}

		rl.DrawRectangle(0, 0, screenWidth, int32(screenUpperLimit), rec_color)

		rl.DrawRectangleRec(boxA, rl.Gold)
		rl.DrawRectangleRec(boxB, rl.Blue)
		if collision {
			// Draw collision area
			rl.DrawRectangleRec(boxCollision, rl.Lime)

			// Draw collision message
			rl.DrawText("COLLISION!", int32(rl.GetScreenWidth()/2) - rl.MeasureText("COLLISION!", 20/2), int32(screenUpperLimit/2-10), 20, rl.Black)
		}
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}

func TestFloodFill(t *testing.T) {
	// Try filling black boxes by pressing mouse clicks
	screenWidth := int32(xMax)
	screenHeight := int32(yMax)
	rl.InitWindow(screenWidth, screenHeight, "Flood fill algorithm | Computer Graphics Demo algorithms")
	defer rl.CloseWindow()

	winImage := rl.LoadImage("res/golang-48.png")
	defer rl.UnloadImage(winImage)
	rl.SetWindowIcon(*winImage)

	checkboard := rl.LoadImage("res/checkboard.png")
	checkboardText := rl.LoadTextureFromImage(checkboard)
	defer rl.UnloadTexture(checkboardText)
	defer rl.UnloadImage(checkboard)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawTexture(checkboardText, 0, 0, rl.White)

		if rl.IsMouseButtonPressed(0) {
			x1, y1 := rl.GetMouseX(), rl.GetMouseY()
			FloodFill(x1, y1, rl.Lime, rl.Black, *checkboard)
			checkboardText = rl.LoadTextureFromImage(checkboard)
		}
		rl.EndDrawing()
	}
	rl.EndDrawing()
}

func TestBoundaryFill(t *testing.T) {
	screenWidth := int32(xMax)
	screenHeight := int32(yMax)
	rl.InitWindow(screenWidth, screenHeight, "Boundary fill algorithm | Computer Graphics Demo algorithms")
	defer rl.CloseWindow()

	winImage := rl.LoadImage("res/golang-48.png")
	defer rl.UnloadImage(winImage)
	rl.SetWindowIcon(*winImage)

	checkboard := rl.LoadImage("res/checkboard.png")
	checkboardText := rl.LoadTextureFromImage(checkboard)
	defer rl.UnloadTexture(checkboardText)
	defer rl.UnloadImage(checkboard)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawTexture(checkboardText, 0, 0, rl.White)
		if rl.IsMouseButtonPressed(0) {
			x1, y1 := rl.GetMouseX(), rl.GetMouseY()
			BoundaryFill(x1, y1, rl.Lime, rl.Black, *checkboard)
			checkboardText = rl.LoadTextureFromImage(checkboard)
		}
		rl.EndDrawing()
	}
	rl.EndDrawing()
}

func TestTransformation(t *testing.T) {
	screenWidth := int32(xMax)
	screenHeight := int32(yMax)
	rl.InitWindow(screenWidth, screenHeight, "3 Transformation | Computer Graphics Demo algorithms")
	defer rl.CloseWindow()
	winImage := rl.LoadImage("res/golang-48.png")
	rl.SetWindowIcon(*winImage)

	state := 0
	var (
		posX   int32 = 90
		posY   int32 = 90
		width  int32 = 40
		height int32 = 30
		screenUpperLimit int32 = 50

		startPosX int32 = 40
		startPosY int32 = 40
		endPosX   int32 = 200
		endPosY   int32 = 200

		thetha  float64 = 60 // angle to rotate line

		centerX   int32 = 450
		centerY   int32 = 450
		t_x int32       = 40
		t_y int32       = 50
	)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawRectangleLines(posX, posY, width, height, rl.Black)
		rl.DrawLine(startPosX, startPosY, endPosX, endPosY, rl.Black)
		rl.DrawCircle(centerX, centerY, 40, rl.Blue)
		if rl.IsMouseButtonPressed(0) {
			if state == 0 {
				height = (height + posY) * 4 - posY
				width  = (width  + posX) * 4 - posX
				state += 1
				rl.DrawText("Performed 4X Scaling on Rectangle!", int32(rl.GetScreenWidth()/2), int32(screenUpperLimit/2-10), 20, rl.Black)
			} else
			if state == 1 {
				endPosX = startPosX * int32(math.Cos(thetha)) - startPosY * int32(math.Sin(thetha))
				endPosY = startPosX * int32(math.Sin(thetha)) + startPosY * int32(math.Cos(thetha))
				state += 1
				rl.DrawText("Performed 60 degree rotation on line!", int32(rl.GetScreenWidth()/2), int32(screenUpperLimit/2-10), 20, rl.Black)
			} else
			if state == 2 {
				centerX = centerX + t_x
				centerY = centerY + t_y
				state += 1
				rl.DrawText("Performed 40x and 50y translation on circle !", int32(rl.GetScreenWidth()/2), int32(screenUpperLimit/2-10), 20, rl.Black)
			}
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func TestCohenSutherlandClip(t *testing.T) {
	l := LineInt{-3, -3, 5000, 4440}
	CohenSutherlandClip(l)
	fmt.Printf("Region code for Point 1 (%d, %d) = %04b \n", l.x1, l.y1, computeRegionCode(l.x1, l.y1))
	fmt.Printf("Region code for Point 2 (%d, %d) = %04b \n", l.x2, l.y2, computeRegionCode(l.x2, l.y2))
}
