package screens

import (
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth int32 = 800
	rows              = 40
	columns           = 40
)

type StartScreen struct {
	cells         [][]int
	screenManager *ScreenManager
	switchToGame  bool
	cellWidth     int32
}

func NewStartScreen(screenManager *ScreenManager) *StartScreen {
	cells := make([][]int, rows)
	switchToGame := false
	cellWidth := int32(30)

	return &StartScreen{
		cells,
		screenManager,
		switchToGame,
		cellWidth,
	}
}

func (s *StartScreen) Process() {
	s.update()
	s.draw()
}

func (s *StartScreen) update() {
	if rg.Button(rl.NewRectangle(200, 200, 200, 50), "Start") {
		s.screenManager.screens = append(s.screenManager.screens, NewGameScreen(s.cells, s.cellWidth))
	}
}

func (s *StartScreen) draw() {
	rl.DrawRectangle(32, 23, 40, 40, rl.Red)
}
