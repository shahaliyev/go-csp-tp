package main

var tileNames = [3]string{"OUTER_BOUNDARY", "EL_SHAPE", "FULL_BLOCK"}

// Covers a specific row with tile
func cover(row []int, method string) {
	if method == "full" {
		for i := range row {
			row[i] = -1
		}
	} else if method == "start" {
		row[0] = -1
	} else if method == "both" {
		row[0] = -1
		row[3] = -1
	}
}

// Puts different types of tiles onto the landscape
func putTile(tileName string, landscape [][]int, X, Y int) [][]int {
	copy := copyMatrix(landscape)

	rows := copy[X:X+4]

	row1 := rows[0][Y:Y+4]
	row2 := rows[1][Y:Y+4]
	row3 := rows[2][Y:Y+4]
	row4 := rows[3][Y:Y+4]

	switch tileName {
	case "OUTER_BOUNDARY":
		cover(row1, "full")
		cover(row2, "both")
		cover(row3, "both")
		cover(row4, "full")
		break
	case "EL_SHAPE":
		cover(row1, "full")
		cover(row2, "start")
		cover(row3, "start")
		cover(row4, "start")
		break
	case "FULL_BLOCK":
		cover(row1, "full")
		cover(row2, "full")
		cover(row3, "full")
		cover(row4, "full")
		break
	}

	return copy
}

// Checks if putting a tile is alright with constraints
func canPutTile(tileName string, landscape [][]int, X, Y int) bool {
	possibleLandscape := putTile(tileName, landscape, X, Y)
	colors := countColors(possibleLandscape)

	for i := 1; i <= 4; i++ {
		if colors[i] < targets[i] {
			return false
		}
	}

	return true
}
