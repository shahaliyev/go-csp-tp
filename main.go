package main

import (
	"fmt"
	"os"
)

func main() {
	input := readInput(os.Args[1]); defer input.Close()

	buildFromFile(input)

	tiles["FULL_BLOCK"] = 12
	fmt.Println(tiles)
	//backtrack()
	fmt.Println(backtrack())
	fmt.Println(targets)
	for _, i := range landscape {
		fmt.Println(i)
	}
	fmt.Println()
	//fmt.Println(canPutTile(tileNames[0], landscape, 0, 0))

}
