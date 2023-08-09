package go_of_life

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

func DrawDevOverlay() {
	fps, frameTime, mousePos := rl.GetFPS(), rl.GetFrameTime(), rl.GetMousePosition()
	text := fmt.Sprintf("FPS: %d\nFrametime: %f\nMouse: %.2f %.2f", fps, frameTime, mousePos.X, mousePos.Y)

	rl.DrawRectangle(0, 0, 300, 200, color.RGBA{255, 255, 255, 120})
	rl.DrawText(text, 15, 12, 20, rl.Red)
}
