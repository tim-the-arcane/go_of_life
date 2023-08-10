package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	components "github.com/tim-the-arcane/go_of_life/components"
	"time"
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
	// Initialize Raylib window
	rl.InitWindow(800, 800, "Go of Life")
	rl.SetTargetFPS(60)

	screenWidth := rl.GetMonitorWidth(rl.GetCurrentMonitor())
	screenHeight := rl.GetMonitorHeight(rl.GetCurrentMonitor())

	rl.SetWindowSize(screenWidth, screenHeight)
	rl.ToggleFullscreen()

	// Initialize game
	game = components.NewGame(int(screenHeight), int(screenWidth), 1)
}

func update() {
	game.Update()

	if rl.IsKeyPressed(rl.KeyD) {
		showDevPanel = !showDevPanel
	}
}

func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	game.Draw()

	if showDevPanel {
		components.DrawDevOverlay()
	}

	rl.EndDrawing()
}
