package screens

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

var currentFrame int32 = 0

type GameScreen struct {
	cellWidth int32
	cells     [][]int
}

func NewGameScreen(cells [][]int, cellWidth int32) *GameScreen {
	var gameScreen GameScreen
	gameScreen.cellWidth = cellWidth
	gameScreen.initialize()
	return &gameScreen
}

func (s *GameScreen) Process() {
	s.update()
	s.draw()
}

func (s *GameScreen) initialize() {
	s.initializeCells()
}

func (s *GameScreen) initializeCells() {
	for i := range s.cells {
		s.cells[i] = make([]int, columns)

		for j := range s.cells[i] {
			if rand.Intn(2) == 1 {
				s.cells[i][j] = 1
			} else {
				s.cells[i][j] = 0
			}
		}
	}
}

func (s *GameScreen) update() {
	var newCells [][]int = make([][]int, rows)
	copy(newCells, s.cells)

	for i := range s.cells {
		for j := range s.cells[i] {
			aliveNeighbours, _ := s.countAliveNeighbors(int32(i), int32(j))

			if aliveNeighbours < 2 || aliveNeighbours > 3 {
				newCells[i][j] = 0
			} else if aliveNeighbours == 3 {
				newCells[i][j] = 1
			} else {
				newCells[i][j]++
			}
		}
	}

	s.cells = newCells
}

func (s *GameScreen) draw() {
	s.drawCells()
}

// counts dead and alive neighbors
func (s *GameScreen) countAliveNeighbors(row int32, column int32) (int32, error) {
	if row < 0 || column < 0 || row >= rows || column >= columns {
		return 0, fmt.Errorf("Invalid coordinates")
	}

	var alive int32 = 0

	// check top left
	if row > 0 && column > 0 {
		if s.cells[row-1][column-1] > 0 {
			alive++
		}
	}

	// check top
	if row > 0 {
		if s.cells[row-1][column] > 0 {
			alive++
		}
	}

	// check top right
	if row > 0 && column < columns-1 {
		if s.cells[row-1][column+1] > 0 {
			alive++
		}
	}

	// check left
	if column > 0 {
		if s.cells[row][column-1] > 0 {
			alive++
		}
	}

	// check right
	if column < columns-1 {
		if s.cells[row][column+1] > 0 {
			alive++
		}
	}

	// check bottom left
	if row < rows-1 && column > 0 {
		if s.cells[row+1][column-1] > 0 {
			alive++
		}
	}

	// check bottom
	if row < rows-1 {
		if s.cells[row+1][column] > 0 {
			alive++
		}
	}

	// check bottom right
	if row < rows-1 && column < columns-1 {
		if s.cells[row+1][column+1] > 0 {
			alive++
		}
	}

	return alive, nil
}

func (s *GameScreen) drawCells() {
	for i := range s.cells {
		for j := range s.cells[i] {
			var color rl.Color

			if s.cells[i][j] > 0 {
				switch s.cells[i][j] {
				case 1:
					color = rl.Lime
				case 2:
					color = rl.Green
				case 3:
					color = rl.DarkGreen
				case 4:
					color = rl.DarkBlue
				case 5:
					color = rl.Blue
				case 6:
					color = rl.SkyBlue
				case 7:
					color = rl.Purple
				case 8:
					color = rl.DarkPurple
				default:
					color = rl.Gold
				}
			} else {
				color = rl.Black
			}

			rl.DrawRectangle(int32(j)*s.cellWidth, int32(i)*s.cellWidth, s.cellWidth, s.cellWidth, color)
		}
	}
}
