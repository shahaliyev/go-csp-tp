package main

import (
	"os"
	"sort"
	"strconv"
)

// Panics in case of errors
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// String to integer converter
// In order to not check for errors each time
func stringToInt(str string) (val int) {
	val, err := strconv.Atoi(str)
	check(err)

	return val
}

func split(s string) []int {
	var res []int

	for i := 0; i < len(s); i = i + 2 {
		if string(s[i]) == " " {
			res = append(res, 0)
		} else {
			res = append(res, stringToInt(string(s[i])))
		}
	}

	return res
}

func findMaxLength(s [][]int) int {
	max := 0

	for i := range s {
		if len(s[i]) > max {
			max = len(s[i])
		}
	}

	return max
}

func fillZeros(s [][]int) [][]int {
	maxLength := findMaxLength(s)

	for i := range s {
		for len(s[i]) < maxLength {
			s[i] = append(s[i], 0)
		}
	}

	return s
}


// Reads the input file for vertices and edges
func readInput(filePath string) *os.File {
	input, err := os.Open(filePath)
	check(err)

	return input
}

// Checks if slice contains a value
func contains(s []int, val int) bool {
	for _, i := range s {
		if i == val {
			return true
		}
	}

	return false
}

// Removes a value from slice
func remove(s []int, val int) []int {
	var res []int

	for _, i := range s {
		if i != val {
			res = append(res, i)
		}
	}

	return res
}

// Converts slice into map
func sliceToMap(s []int) map[int]int {
	res := make(map[int]int)

	for i := 0; i < len(s); i++ {
		res[s[i]] = 0
	}

	return res
}

// Sorting map by values
// https://stackoverflow.com/questions/18695346/how-to-sort-a-mapstringint-by-its-values
func sortMap(mp map[int]int) PairList{
	pl := make(PairList, len(mp))
	i := 0
	for k, v := range mp {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Color int
	Value int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }

func countColors(landscape [][]int) map[int]int {
	res := make(map[int]int)

	for i := range landscape {
		for j := range landscape[i] {
			res[landscape[i][j]]++
		}
	}

	return res
}

func copyMatrix(matrix [][]int) [][]int {
	duplicate := make([][]int, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]int, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}

	return duplicate
}



