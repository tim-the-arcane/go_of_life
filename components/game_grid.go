package components

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tim-the-arcane/go_of_life/helpers"
	"math"
	"math/rand"
)

type Game interface {
	Update()
	Draw()
}

type gameImpl struct {
	rows      int
	columns   int
	grid      [][]bool
	cellWidth int32
	running   bool
}

func NewGame(rows int, columns int, cellWidth int32) Game {
	newGame := &gameImpl{}
	newGame.rows = rows
	newGame.columns = columns
	newGame.cellWidth = cellWidth
	newGame.initialize()

	return newGame
}

func (g *gameImpl) Update() {
	g.processInput()

	if g.running {
		g.nextGeneration()
	}
}

func (g *gameImpl) Draw() {
	screenWidth, _ := int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight())

	for i := range g.grid {
		for j := range g.grid[i] {
			if g.grid[i][j] {
				rl.DrawRectangle(int32(i)*g.cellWidth, int32(j)*g.cellWidth, g.cellWidth, g.cellWidth, rl.Green)
			}
		}
	}

	fontsize := 30
	textsize := rl.MeasureText("Playing", int32(fontsize))
	var playstate string

	if g.running {
		playstate = "Playing"
	} else {
		playstate = "Paused"
	}

	rl.DrawText(playstate, int32(screenWidth)-30-int32(textsize), int32(fontsize), int32(fontsize), rl.RayWhite)
}

func (g *gameImpl) initialize() {
	fmt.Println("New Game initialized")
	g.grid = helpers.GenerateNewGrid(g.rows, g.columns)
}

func (g *gameImpl) processInput() {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		fmt.Println("CLICK", rl.MouseLeftButton)
		if !g.running {
			mousePos := rl.GetMousePosition()

			if inWindow, _ := helpers.IsMouseInWindow(); inWindow {
				col := int(math.Floor(float64(mousePos.X) / float64(g.cellWidth)))
				row := int(math.Floor(float64(mousePos.Y) / float64(g.cellWidth)))
				fmt.Printf("Clicked on row %d and col %d \n", row+1, col+1)
				g.toggleCell(row, col)
			}
		} else {
			g.pause()
		}
	}

	switch rl.GetKeyPressed() {
	case rl.KeySpace:
		fmt.Println("Pressed space")
		if g.running {
			g.pause()
		} else {
			g.play()
		}
	case rl.KeyK:
		g.killAll()
	case rl.KeyR:
		g.randomizeCells()
	case rl.KeyN:
		g.nextGeneration()
	}
}

func (g *gameImpl) nextGeneration() {
	newGeneration := helpers.GenerateNewGrid(g.rows, g.columns)

	for i := range g.grid {
		for j := range g.grid[i] {
			if alive := helpers.CountAliveNeighbours(&g.grid, j, i); alive < 2 || alive > 3 {
				newGeneration[i][j] = false
			} else if alive == 3 {
				newGeneration[i][j] = true
			} else {
				newGeneration[i][j] = g.grid[i][j]
			}
		}
	}

	g.grid = newGeneration
}

func (g *gameImpl) setCell(row int, column int, state bool) {
	g.grid[column][row] = state
}

func (g *gameImpl) toggleCell(row int, column int) {
	g.setCell(row, column, !g.grid[column][row])
}

func (g *gameImpl) randomizeCells() {
	for i := range g.grid {
		for j := range g.grid[i] {
			if rand.Intn(2) == 1 {
				g.grid[i][j] = true
			} else {
				g.grid[i][j] = false
			}
		}
	}

	fmt.Println("Randomized cells!")
}

func (g *gameImpl) killAll() {
	for i := range g.grid {
		for j := range g.grid[i] {
			g.grid[i][j] = false
		}
	}

	fmt.Println("Killed all cells!")
}

func (g *gameImpl) play() {
	if !g.running {
		fmt.Println("Playing")
		g.running = true
	}
}

func (g *gameImpl) pause() {
	if g.running {
		fmt.Println("Paused")
		g.running = false
	}
}
