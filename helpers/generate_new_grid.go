package helpers

func GenerateNewGrid(rows int, colums int) (generatedGrid [][]bool) {
	generatedGrid = make([][]bool, colums)

	for i := range generatedGrid {
		generatedGrid[i] = make([]bool, rows)
	}

	return generatedGrid
}
