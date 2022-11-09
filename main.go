package main

import (
	"example/sudoku_solver/internal"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	sudokuProblem := [9][9]int{
		{2, 3, 9, 1, 5, 1, 2, 8, 2},
		{7, 5, 1, 1, 3, 1, 1, 5, 6},
		{8, 6, 4, 4, 7, 5, 1, 9, 8},
		{8, 9, 6, 7, 3, 6, 6, 6, 1},
		{2, 1, 3, 5, 9, 4, 8, 6, 7},
		{8, 7, 3, 8, 5, 6, 8, 1, 4},
		{9, 5, 2, 5, 5, 8, 7, 6, 5},
		{3, 3, 9, 2, 5, 7, 6, 1, 9},
		{9, 9, 4, 9, 3, 4, 9, 9, 3},
	}

	grid, _ := internal.ConvertToGrid(sudokuProblem)
	fmt.Println(grid.CheckCompleteness())
	grid.PrintGrid()

	emptyGrid := internal.CreateEmptyGrid()
	internal.FillGrid(emptyGrid)
	fmt.Println(emptyGrid.CheckCompleteness())
	emptyGrid.PrintGrid()
}
