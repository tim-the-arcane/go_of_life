package helpers

import (
	"errors"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func IsMouseInWindow() (result bool, err error) {
	if !rl.IsWindowReady() {
		err = errors.New("Window was not initialized!")
		return
	}

	if mousePos := rl.GetMousePosition(); mousePos.X >= 0 && mousePos.Y >= 0 && mousePos.X <= float32(rl.GetScreenWidth()) && mousePos.Y <= float32(rl.GetScreenHeight()) {
		return true, nil
	}

	return false, nil
}
