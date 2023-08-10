package components

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

func DrawDevOverlay() {
	fps, mousePos := rl.GetFPS(), rl.GetMousePosition()
	text := fmt.Sprintf("FPS: %d\nMouse: %.2f %.2f", fps, mousePos.X, mousePos.Y)
	textDimensions := rl.MeasureTextEx(rl.GetFontDefault(), text, 20, 10)

	rl.DrawRectangle(0, 0, int32(textDimensions.X+30), int32(textDimensions.Y+30), color.RGBA{255, 255, 255, 255})
	rl.DrawText(text, 15, 12, 20, rl.Red)
}
