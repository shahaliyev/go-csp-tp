package main

import (
	"bufio"
	"os"
	"strings"
)

// Builds graphs by scanning each line of the input file
func buildFromFile(input *os.File) {
	sections := [3]string{"landscape", "tiles", "targets"}
	cnt := -2 // number of initial comments

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "#") {
			cnt++
			continue
		}

		if cnt == 4 {
			break
		}

		section := sections[cnt]

		if section == "landscape" {
			tokens := split(line)
			landscape = append(landscape, tokens)
		} else if section == "tiles" {
			for _, tile := range tileNames {
				index := strings.Index(line, tile) + len(tile) + 1
				tiles[tile] = stringToInt(string(line[index]))
			}
		} else if section == "targets" {
			tokens := strings.Split(line, ":")
			targets[stringToInt(tokens[0])] = stringToInt(tokens[1])
		}

		err := scanner.Err();
		check(err)
	}

	landscape = fillZeros(landscape)
}