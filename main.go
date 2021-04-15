package main

import (
	"fmt"
	"os"
)

func main() {
	input := readInput(os.Args[1]); defer input.Close()
	buildFromFile(input)
	fmt.Println("Solving...")
	backtrack(0, 0)
	showOutput()
}
