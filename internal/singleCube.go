package internal

type SingleCube struct {
	TopRowCube    [3]int
	MiddleRowCube [3]int
	BottomRowCube [3]int
}

func (sc SingleCube) IsCubeComplete() (bool, error) {
	var allNums []int
	allNums = append(allNums, sc.TopRowCube[:]...)
	allNums = append(allNums, sc.MiddleRowCube[:]...)
	allNums = append(allNums, sc.BottomRowCube[:]...)
	return OneToNine(allNums)
}

func (sc SingleCube) IsCubeValid() bool {
	var allNums []int
	allNums = append(allNums, sc.TopRowCube[:]...)
	allNums = append(allNums, sc.MiddleRowCube[:]...)
	allNums = append(allNums, sc.BottomRowCube[:]...)
	return IncompleteValidity(allNums)
}
