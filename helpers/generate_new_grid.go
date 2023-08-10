package helpers

func GenerateNewGrid(rows int, colums int) (generatedGrid [][]bool) {
	generatedGrid = make([][]bool, rows)

	for i := range generatedGrid {
		generatedGrid[i] = make([]bool, colums)
	}

	return generatedGrid
}
