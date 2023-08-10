package helpers

import "fmt"

func CountAliveNeighbours(grid *[][]bool, rowIndex int, columnIndex int) (alive int) {
	columns, rows := len(*grid), len((*grid)[columnIndex])
	if columnIndex < 0 || rowIndex < 0 || columnIndex > columns-1 || rowIndex > rows-1 {
		panic(fmt.Sprintf("Index out of bounds!"))
	}

	// Top Left
	if rowIndex > 0 && columnIndex > 0 && (*grid)[columnIndex-1][rowIndex-1] {
		alive++
	}

	// Top
	if rowIndex > 0 && (*grid)[columnIndex][rowIndex-1] {
		alive++
	}

	// Top Right
	if rowIndex > 0 && columnIndex < columns-1 && (*grid)[columnIndex+1][rowIndex-1] {
		alive++
	}

	// Right
	if columnIndex < columns-1 && (*grid)[columnIndex+1][rowIndex] {
		alive++
	}

	// Bottom Right
	if rowIndex < rows-1 && columnIndex < columns-1 && (*grid)[columnIndex+1][rowIndex+1] {
		alive++
	}

	// Bottom
	if rowIndex < rows-1 && (*grid)[columnIndex][rowIndex+1] {
		alive++
	}

	// Bottom Left
	if rowIndex < rows-1 && columnIndex > 0 && (*grid)[columnIndex-1][rowIndex+1] {
		alive++
	}

	// Left
	if columnIndex > 0 && (*grid)[columnIndex-1][rowIndex] {
		alive++
	}

	return alive
}
