package main

import "strconv"

func backtrack(locX, locY int) bool {

	if isFound() {
		return true
	}

	// Iterating over domains
	for _, tileName := range tileNames {

		// checking if a tile type is left
		if tiles[tileName] == 0 {
			continue
		}

		copy := copyMatrix(landscape)

		// If putting the tile won't break the constraints
		if canPutTile(tileName, landscape, locX, locY) {

			// removing tile from the domain
			tiles[tileName]--

			// putting tile on the landscape
			landscape = putTile(tileName, landscape, locX, locY)

			// building solution map
			solution[strconv.Itoa(locX) + " " + strconv.Itoa(locY)] = tileName

			// memorizing location and getting next one
			prevLocY, prevLocX := locY, locX
			locX, locY = getNextLocation(locX, locY)

			if backtrack(locX, locY) {
				return true
			}

			// backtracking
			locX, locY = prevLocX, prevLocY
			landscape = copyMatrix(copy)
			tiles[tileName]++
		}
	}

	return false
}

// Checks if the solution is found
func isFound() bool {
	colors := countColors(landscape)
	cnt := 0

	for i := 1; i <= 4; i++ {
		if colors[i] == targets[i] {
			cnt++
		}
	}

	if cnt == 4 {
		return true
	}

	return false
}

// Gets next location to put a tile
func getNextLocation(locX, locY int) (int, int) {
	if locY + 4 < len(landscape[0]) {
		locY += 4
	} else {
		locY = 0

		if locX + 4 < len(landscape) {
			locX += 4
		} else {
			locX = 0
		}
	}

	return locX, locY
}
