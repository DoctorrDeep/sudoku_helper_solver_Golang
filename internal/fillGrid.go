package internal

import "math/rand"

var GridCellValues = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func FillGrid(sg *SudokuGrid) {
	rand.Shuffle(len(GridCellValues), func(i, j int) {
		GridCellValues[i], GridCellValues[j] = GridCellValues[j], GridCellValues[i]
	})
	for _, cellValue := range GridCellValues {
		for i := range []int{0, 1, 2} {
			for j := range []int{0, 1, 2} {
				if sg.TopRowGrid[i].TopRowCube[j] == 0 {
					sg.TopRowGrid[i].TopRowCube[j] = cellValue
					xPos := 3*i + j
					yPos := 0
					xyOkay := sg.CheckLineValidity(xPos, yPos)
					cubeOkay := sg.TopRowGrid[i].IsCubeValid()
					if !(xyOkay && cubeOkay) {
						sg.TopRowGrid[i].TopRowCube[j] = 0
					}
				}
				if sg.TopRowGrid[i].MiddleRowCube[j] == 0 {
					sg.TopRowGrid[i].MiddleRowCube[j] = cellValue
					xPos := 3*i + j
					yPos := 1
					xyOkay := sg.CheckLineValidity(xPos, yPos)
					cubeOkay := sg.TopRowGrid[i].IsCubeValid()
					if !(xyOkay && cubeOkay) {
						sg.TopRowGrid[i].MiddleRowCube[j] = 0
					}
				}
				if sg.TopRowGrid[i].BottomRowCube[j] == 0 {
					sg.TopRowGrid[i].BottomRowCube[j] = cellValue
					xPos := 3*i + j
					yPos := 2
					xyOkay := sg.CheckLineValidity(xPos, yPos)
					cubeOkay := sg.TopRowGrid[i].IsCubeValid()
					if !(xyOkay && cubeOkay) {
						sg.TopRowGrid[i].BottomRowCube[j] = 0
					}
				}
				if sg.MiddleRowGrid[i].TopRowCube[j] == 0 {
					sg.MiddleRowGrid[i].TopRowCube[j] = cellValue
					xPos := 3*i + j
					yPos := 3
					xyOkay := sg.CheckLineValidity(xPos, yPos)
					cubeOkay := sg.MiddleRowGrid[i].IsCubeValid()
					if !(xyOkay && cubeOkay) {
						sg.MiddleRowGrid[i].TopRowCube[j] = 0
					}
				}
				if sg.MiddleRowGrid[i].MiddleRowCube[j] == 0 {
					sg.MiddleRowGrid[i].MiddleRowCube[j] = cellValue
					xPos := 3*i + j
					yPos := 4
					xyOkay := sg.CheckLineValidity(xPos, yPos)
					cubeOkay := sg.MiddleRowGrid[i].IsCubeValid()
					if !(xyOkay && cubeOkay) {
						sg.MiddleRowGrid[i].MiddleRowCube[j] = 0
					}
				}
				if sg.MiddleRowGrid[i].BottomRowCube[j] == 0 {
					sg.MiddleRowGrid[i].BottomRowCube[j] = cellValue
					xPos := 3*i + j
					yPos := 5
					xyOkay := sg.CheckLineValidity(xPos, yPos)
					cubeOkay := sg.MiddleRowGrid[i].IsCubeValid()
					if !(xyOkay && cubeOkay) {
						sg.MiddleRowGrid[i].BottomRowCube[j] = 0
					}
				}
				if sg.BottomRowGrid[i].TopRowCube[j] == 0 {
					sg.BottomRowGrid[i].TopRowCube[j] = cellValue
					xPos := 3*i + j
					yPos := 6
					xyOkay := sg.CheckLineValidity(xPos, yPos)
					cubeOkay := sg.BottomRowGrid[i].IsCubeValid()
					if !(xyOkay && cubeOkay) {
						sg.BottomRowGrid[i].TopRowCube[j] = 0
					}
				}
				if sg.BottomRowGrid[i].MiddleRowCube[j] == 0 {
					sg.BottomRowGrid[i].MiddleRowCube[j] = cellValue
					xPos := 3*i + j
					yPos := 7
					xyOkay := sg.CheckLineValidity(xPos, yPos)
					cubeOkay := sg.BottomRowGrid[i].IsCubeValid()
					if !(xyOkay && cubeOkay) {
						sg.BottomRowGrid[i].MiddleRowCube[j] = 0
					}
				}
				if sg.BottomRowGrid[i].BottomRowCube[j] == 0 {
					sg.BottomRowGrid[i].BottomRowCube[j] = cellValue
					xPos := 3*i + j
					yPos := 8
					xyOkay := sg.CheckLineValidity(xPos, yPos)
					cubeOkay := sg.BottomRowGrid[i].IsCubeValid()
					if !(xyOkay && cubeOkay) {
						sg.BottomRowGrid[i].BottomRowCube[j] = 0
					}
				}
			}
		}
	}
}
