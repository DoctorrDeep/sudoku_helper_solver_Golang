package internal

import "sort"

type DeDuplicatedSlice struct {
	m map[int]bool
}

func (d *DeDuplicatedSlice) AddValues(i int) {
	d.m[i] = true
}

func (d *DeDuplicatedSlice) GetKeys() []int {
	var allKeys []int
	for mapIndex := range d.m {
		allKeys = append(allKeys, mapIndex)
	}
	return allKeys
}

func makeDeDuplicatedSlices(allNums []int) []int {
	var d DeDuplicatedSlice
	d.m = make(map[int]bool)
	for _, i := range allNums {
		d.AddValues(i)
	}
	return d.GetKeys()
}

func OneToNine(allNums []int) (bool, error) {
	allNums = makeDeDuplicatedSlices(allNums)
	sort.Ints(allNums)
	if len(allNums) != 9 || allNums[0] < 1 || allNums[8] > 9 {
		return false, nil
	}
	return true, nil
}
