package internal

func ConvertToGrid(problem [9][9]int) (SudokuGrid, error) {

	emptyGrid := SudokuGrid{}

	emptyGrid.TopRowGrid[0].TopRowCube = [3]int{problem[0][0], problem[0][1], problem[0][2]}
	emptyGrid.TopRowGrid[0].MiddleRowCube = [3]int{problem[1][0], problem[1][1], problem[1][2]}
	emptyGrid.TopRowGrid[0].BottomRowCube = [3]int{problem[2][0], problem[2][1], problem[2][2]}

	emptyGrid.TopRowGrid[1].TopRowCube = [3]int{problem[0][3], problem[0][4], problem[0][5]}
	emptyGrid.TopRowGrid[1].MiddleRowCube = [3]int{problem[1][3], problem[1][4], problem[1][5]}
	emptyGrid.TopRowGrid[1].BottomRowCube = [3]int{problem[2][3], problem[2][4], problem[2][5]}

	emptyGrid.TopRowGrid[2].TopRowCube = [3]int{problem[0][6], problem[0][7], problem[0][8]}
	emptyGrid.TopRowGrid[2].MiddleRowCube = [3]int{problem[1][6], problem[1][7], problem[1][8]}
	emptyGrid.TopRowGrid[2].BottomRowCube = [3]int{problem[2][6], problem[2][7], problem[2][8]}

	emptyGrid.MiddleRowGrid[0].TopRowCube = [3]int{problem[3][0], problem[3][1], problem[3][2]}
	emptyGrid.MiddleRowGrid[0].MiddleRowCube = [3]int{problem[4][0], problem[4][1], problem[4][2]}
	emptyGrid.MiddleRowGrid[0].BottomRowCube = [3]int{problem[5][0], problem[5][1], problem[5][2]}

	emptyGrid.MiddleRowGrid[1].TopRowCube = [3]int{problem[3][3], problem[3][4], problem[3][5]}
	emptyGrid.MiddleRowGrid[1].MiddleRowCube = [3]int{problem[4][3], problem[4][4], problem[4][5]}
	emptyGrid.MiddleRowGrid[1].BottomRowCube = [3]int{problem[5][3], problem[5][4], problem[5][5]}

	emptyGrid.MiddleRowGrid[2].TopRowCube = [3]int{problem[3][6], problem[3][7], problem[3][8]}
	emptyGrid.MiddleRowGrid[2].MiddleRowCube = [3]int{problem[4][6], problem[4][7], problem[4][8]}
	emptyGrid.MiddleRowGrid[2].BottomRowCube = [3]int{problem[5][6], problem[5][7], problem[5][8]}

	emptyGrid.BottomRowGrid[0].TopRowCube = [3]int{problem[6][0], problem[6][1], problem[6][2]}
	emptyGrid.BottomRowGrid[0].MiddleRowCube = [3]int{problem[7][0], problem[7][1], problem[7][2]}
	emptyGrid.BottomRowGrid[0].BottomRowCube = [3]int{problem[8][0], problem[8][1], problem[8][2]}

	emptyGrid.BottomRowGrid[1].TopRowCube = [3]int{problem[6][3], problem[6][4], problem[6][5]}
	emptyGrid.BottomRowGrid[1].MiddleRowCube = [3]int{problem[7][3], problem[7][4], problem[7][5]}
	emptyGrid.BottomRowGrid[1].BottomRowCube = [3]int{problem[8][3], problem[8][4], problem[8][5]}

	emptyGrid.BottomRowGrid[2].TopRowCube = [3]int{problem[6][6], problem[6][7], problem[6][8]}
	emptyGrid.BottomRowGrid[2].MiddleRowCube = [3]int{problem[7][6], problem[7][7], problem[7][8]}
	emptyGrid.BottomRowGrid[2].BottomRowCube = [3]int{problem[8][6], problem[8][7], problem[8][8]}

	var err error
	return emptyGrid, err
}
