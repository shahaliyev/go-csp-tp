package main

func canPutTile(tile string, landscape [][]int, X, Y int) bool {
	possibleLandscape := putTile(tile, landscape, X, Y)
	colors := countColors(possibleLandscape)

	for i := 1; i <= 4; i++ {
		if colors[i] < targets[i] {
			return false
		}
	}

	return true
}

var locX int
var locY int
var iter int
func backtrack() bool {
	//if iter == 15 {
	//	return true
	//}
	//iter++

	colors := countColors(landscape)
	//fmt.Println("colors: ", colors)

	cnt := 0
	for i := 1; i <= 4; i++ {
		if colors[i] == targets[i] {
			cnt++
		}
	}
	if cnt == 4 {
		return true
	}

	for _, tileName := range tileNames {
		if tiles[tileName] == 0 {
			continue
		}

		//fmt.Println("tileName: ", tileName)

		copy := copyMatrix(landscape)

		if canPutTile(tileName, landscape, locX, locY) {
			tiles[tileName]--
			landscape = putTile(tileName, landscape, locX, locY)

			prevLocY := locY
			prevLocX := locX

			if locY + 4 < len(landscape[0]) {
				locY += 4
			} else {
				locY = 0

				if locX + 4 < len(landscape) {
					locX += 4
				}
			}

			//for _, i := range landscape {
			//	fmt.Println(i)
			//}
			//fmt.Println()

			if backtrack() {
				return true
			}

			locY = prevLocY
			locX = prevLocX

			landscape = copyMatrix(copy)

			tiles[tileName]++
		}
	}

	//fmt.Println("NOW!")
	//fmt.Println(countColors(landscape))
	//fmt.Println(tiles)
	//for _, i := range landscape {
	//	fmt.Println(i)
	//}
	//fmt.Println()

	return false
}
