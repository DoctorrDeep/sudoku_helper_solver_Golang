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
	horizontalCompleteness := true
	verticalCompleteness := true

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
		tempHorizontalCompleteness, _ := OneToNine(v)
		fmt.Println(v, tempHorizontalCompleteness)
		if !tempHorizontalCompleteness {
			horizontalCompleteness = false
		}
	}

	for _, v := range sg.GetVerticalLines() {
		tempVerticalCompleteness, _ := OneToNine(v)
		fmt.Println(v, tempVerticalCompleteness)
		if !tempVerticalCompleteness {
			verticalCompleteness = false
		}
	}

	return cubeCompleteness && horizontalCompleteness && verticalCompleteness, nil

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

func (sg SudokuGrid) GetVerticalLines() [][]int {

	var verticalLines [][]int
	verticalSlices := make(map[int][]int)

	for _, cubeRow := range sg.GetRowOfCubes() {
		verticalSlices[0] = append(verticalSlices[0], []int{cubeRow[0].TopRowCube[0], cubeRow[0].MiddleRowCube[0], cubeRow[0].BottomRowCube[0]}...)
		verticalSlices[1] = append(verticalSlices[1], []int{cubeRow[0].TopRowCube[1], cubeRow[0].MiddleRowCube[1], cubeRow[0].BottomRowCube[1]}...)
		verticalSlices[2] = append(verticalSlices[2], []int{cubeRow[0].TopRowCube[2], cubeRow[0].MiddleRowCube[2], cubeRow[0].BottomRowCube[2]}...)

		verticalSlices[3] = append(verticalSlices[3], []int{cubeRow[1].TopRowCube[0], cubeRow[1].MiddleRowCube[0], cubeRow[1].BottomRowCube[0]}...)
		verticalSlices[4] = append(verticalSlices[4], []int{cubeRow[1].TopRowCube[1], cubeRow[1].MiddleRowCube[1], cubeRow[1].BottomRowCube[1]}...)
		verticalSlices[5] = append(verticalSlices[5], []int{cubeRow[1].TopRowCube[2], cubeRow[1].MiddleRowCube[2], cubeRow[1].BottomRowCube[2]}...)

		verticalSlices[6] = append(verticalSlices[6], []int{cubeRow[2].TopRowCube[0], cubeRow[2].MiddleRowCube[0], cubeRow[2].BottomRowCube[0]}...)
		verticalSlices[7] = append(verticalSlices[7], []int{cubeRow[2].TopRowCube[1], cubeRow[2].MiddleRowCube[1], cubeRow[2].BottomRowCube[1]}...)
		verticalSlices[8] = append(verticalSlices[8], []int{cubeRow[2].TopRowCube[2], cubeRow[2].MiddleRowCube[2], cubeRow[2].BottomRowCube[2]}...)
	}

	for _, v := range verticalSlices {
		verticalLines = append(verticalLines, v)
	}

	return verticalLines
}

func (sg SudokuGrid) PrintGrid() {
	for _, v := range sg.GetHorizontalLines() {
		fmt.Println(v)
	}
}
