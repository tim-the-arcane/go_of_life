package main

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"math"
	"math/rand"
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
)

func main() {
	initialize()

	for !rl.WindowShouldClose() {
		if currentFrame > 6 {
			update()
			currentFrame = 0
		}
		draw()
	}

	rl.CloseWindow()
}

func initialize() {
	fmt.Println("Initializing...")
	fmt.Println("Rows: ", rows)
	fmt.Println("Columns: ", columns)

	for i := range cells {
		cells[i] = make([]int, columns)

		for j := range cells[i] {
			if rand.Intn(2) == 1 {
				cells[i][j] = 1
			} else {
				cells[i][j] = 0
			}
		}
	}

	rl.InitWindow(screenWidth, screenHeight, "Go of Life")
	rl.SetTargetFPS(fps)
}

func update() {
	var newCells [][]int = make([][]int, rows)
	copy(newCells, cells)

	for i := range cells {
		for j := range cells[i] {
			aliveNeighbours, _ := countAliveNeighbors(int32(i), int32(j))

			if aliveNeighbours < 2 || aliveNeighbours > 3 {
				newCells[i][j] = 0
			} else if aliveNeighbours == 3 {
				newCells[i][j] = 1
			} else {
				newCells[i][j]++
			}
		}
	}

	cells = newCells
}

func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	drawCells()
	drawFps()

	rl.EndDrawing()
	currentFrame++
}

// counts dead and alive neighbors
func countAliveNeighbors(row int32, column int32) (int32, error) {
	if row < 0 || column < 0 || row >= rows || column >= columns {
		return 0, fmt.Errorf("Invalid coordinates")
	}

	var alive int32 = 0

	// check top left
	if row > 0 && column > 0 {
		if cells[row-1][column-1] > 0 {
			alive++
		}
	}

	// check top
	if row > 0 {
		if cells[row-1][column] > 0 {
			alive++
		}
	}

	// check top right
	if row > 0 && column < columns-1 {
		if cells[row-1][column+1] > 0 {
			alive++
		}
	}

	// check left
	if column > 0 {
		if cells[row][column-1] > 0 {
			alive++
		}
	}

	// check right
	if column < columns-1 {
		if cells[row][column+1] > 0 {
			alive++
		}
	}

	// check bottom left
	if row < rows-1 && column > 0 {
		if cells[row+1][column-1] > 0 {
			alive++
		}
	}

	// check bottom
	if row < rows-1 {
		if cells[row+1][column] > 0 {
			alive++
		}
	}

	// check bottom right
	if row < rows-1 && column < columns-1 {
		if cells[row+1][column+1] > 0 {
			alive++
		}
	}

	return alive, nil
}

func drawCells() {
	for i := range cells {
		for j := range cells[i] {
			var color rl.Color

			if cells[i][j] > 0 {
				switch cells[i][j] {
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

			rl.DrawRectangle(int32(j)*cellWidth, int32(i)*cellWidth, cellWidth, cellWidth, color)
		}
	}
}

func drawFps() {
	fps, frameTime := rl.GetFPS(), rl.GetFrameTime()

	rl.DrawText(fmt.Sprintf("FPS: %d\nFrametime: %f", fps, frameTime), 15, 12, 20, rl.Red)
}
