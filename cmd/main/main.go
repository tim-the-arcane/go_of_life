package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	components "github.com/tim-the-arcane/go_of_life/components"
	"time"
)

// Config
const (
	rows        int32 = 80
	columns     int32 = 80
	screenWidth int32 = 800
	fps         int32 = 60
)

var (
	showDevPanel   bool = false
	game           components.Game
	currentFrame   int
	lastUpdateTime time.Time
)

func main() {
	initialize()
	defer rl.CloseWindow()

	lastUpdateTime = time.Now()

	for !rl.WindowShouldClose() {
		currentTime := time.Now()
		elapsed := currentTime.Sub(lastUpdateTime).Seconds()

		if elapsed >= 1.0 {
			update()
			draw()
		}
	}
}

func initialize() {
	// Calculate dimensions
	cellWidth := int32(screenWidth / columns)
	screenWidth := columns * cellWidth
	screenHeight := rows * cellWidth

	// Initialize Raylib window
	rl.InitWindow(screenWidth, screenHeight, "Go of Life")
	rl.SetTargetFPS(fps)

	// Initialize game
	game = components.NewGame(int(rows), int(columns), cellWidth)
}

func update() {
	game.Update()

	if rl.IsKeyPressed(rl.KeyD) {
		showDevPanel = !showDevPanel
	}
}

func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	game.Draw()

	if showDevPanel {
		components.DrawDevOverlay()
	}

	rl.EndDrawing()
}
