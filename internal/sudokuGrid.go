package internal

import "fmt"

type SudokuGrid struct {
	TopRowGrid    [3]SingleCube
	MiddleRowGrid [3]SingleCube
	BottomRowGrid [3]SingleCube
}

func (sg SudokuGrid) GetRowOfCubes() [3]*[3]SingleCube {
	cubeRows := [3]*[3]SingleCube{&sg.TopRowGrid, &sg.MiddleRowGrid, &sg.BottomRowGrid}
	return cubeRows
}

func (sg SudokuGrid) CheckCompleteness() (bool, error) {
	cubeCompleteness := true
	verticalCompleteness := true
	horizontalCompleteness := true

	for _, cubeRow := range sg.GetRowOfCubes() {
		for _, v := range cubeRow {
			tempCubeCompleteness, _ := v.IsCubeComplete()
			fmt.Println(v, tempCubeCompleteness)
			if !tempCubeCompleteness {
				cubeCompleteness = false
			}
		}
	}

	for _, v := range sg.GetHorizontalLines() {
		tempVerticalCompleteness, _ := OneToNine(v)
		fmt.Println(v, tempVerticalCompleteness)
		if !tempVerticalCompleteness {
			verticalCompleteness = false
		}
	}

	return cubeCompleteness && verticalCompleteness && horizontalCompleteness, nil

}

func (sg SudokuGrid) GetHorizontalLines() [][]int {

	var (
		topRow          []int
		middleRow       []int
		bottomRow       []int
		horizontalLines [][]int
	)

	for _, cubeRow := range sg.GetRowOfCubes() {
		for _, v := range cubeRow {
			topRow = append(topRow, v.TopRowCube[:]...)
			middleRow = append(middleRow, v.MiddleRowCube[:]...)
			bottomRow = append(bottomRow, v.BottomRowCube[:]...)
		}
		horizontalLines = append(horizontalLines, topRow)
		horizontalLines = append(horizontalLines, middleRow)
		horizontalLines = append(horizontalLines, bottomRow)
		topRow = nil
		middleRow = nil
		bottomRow = nil
	}
	return horizontalLines
}

func (sg SudokuGrid) PrintGrid() {
	for _, v := range sg.GetHorizontalLines() {
		fmt.Println(v)
	}
}
