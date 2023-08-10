package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	// "github.com/tim-the-arcane/go_of_life/pkg/screens"
	"math"
)

// Config
const (
	fps         int32 = 60
	screenWidth int32 = 800
	rows        int32 = 40
	columns     int32 = 40
)

var (
	cellWidth    int32   = int32(math.Ceil(float64(screenWidth) / float64(columns)))
	screenHeight int32   = rows * cellWidth
	cells        [][]int = make([][]int, rows)
	currentFrame int32   = 0
	// screenManager screens.ScreenManager = *screens.NewScreenManager()
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Go of Life")
	rl.SetTargetFPS(fps)

	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		rl.DrawRectangle(30, 30, 30, 30, rl.Red)
	}
}
