package components

import (
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Toolbar interface {
	Update()
	AddTool(tool Tool)
}

type toolbarImpl struct {
	game       *Game
	buttonSize int32
	spacing    int32
	tools      []Tool
}

func (tb *toolbarImpl) NewToolbar(game *Game) Toolbar {
	toolbar := &toolbarImpl{}
	toolbar.game = game

	return toolbar
}

func (tb *toolbarImpl) AddTool(tool Tool) {
	tb.tools = append(tb.tools, tool)
}

func (tb *toolbarImpl) Update() {

}

type Tool interface {
	Update()
}

type buttonImpl struct {
	label   string
	x       int32
	y       int32
	spacing int32
	onClick func()
}

func (btn *buttonImpl) NewButton(label string, x int32, y int32, onClick func()) Tool {
	button := &buttonImpl{}
	button.label = label
	button.x = x
	button.y = y
	button.onClick = onClick

	return button
}

func (btn *buttonImpl) Update() {
	if rg.Button(rl.NewRectangle(float32(btn.x), float32(btn.y), 100, 20), btn.label) {

	}
}
