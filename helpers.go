package main

import (
	"os"
	"strconv"
)

// Panics in case of errors
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Reads the input file for vertices and edges
func readInput(filePath string) *os.File {
	input, err := os.Open(filePath)
	check(err)

	return input
}

// Converts string to integer
func stringToInt(str string) (val int) {
	val, err := strconv.Atoi(str)
	check(err)

	return val
}

// Splits input
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

// Counts the number of each color and returns their map
func countColors(landscape [][]int) map[int]int {
	res := make(map[int]int)

	for i := range landscape {
		for j := range landscape[i] {
			res[landscape[i][j]]++
		}
	}

	return res
}

// Copies matrix with pass by value
func copyMatrix(matrix [][]int) [][]int {
	duplicate := make([][]int, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]int, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}

	return duplicate
}

// Finds the max length within a matrix
func findMaxLength(s [][]int) int {
	max := 0

	for i := range s {
		if len(s[i]) > max {
			max = len(s[i])
		}
	}

	return max
}

// Fills the trimmed whitespaces with zeros
func fillZeros(s [][]int) [][]int {
	maxLength := findMaxLength(s)

	for i := range s {
		for len(s[i]) < maxLength {
			s[i] = append(s[i], 0)
		}
	}

	return s
}