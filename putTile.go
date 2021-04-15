package main

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

func putTile(tile string, landscape [][]int, X, Y int) [][]int {
	copy := copyMatrix(landscape)

	rows := copy[X:X+4]

	row1 := rows[0][Y:Y+4]
	row2 := rows[1][Y:Y+4]
	row3 := rows[2][Y:Y+4]
	row4 := rows[3][Y:Y+4]

	switch tile {
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
