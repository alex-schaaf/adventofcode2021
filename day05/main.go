package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFileLines(filepath string, split string) []string {
	content, _ := ioutil.ReadFile(filepath)
	return strings.Split(string(content), split)
}

func main() {
	lines := readFileLines("./input", "\n")

	coords := ParseCoordinates(lines)
	xMax, yMax := GetMaxCoords(coords)
	fmt.Println()

	diagram := make([][]uint8, yMax+1)
	for i := range diagram {
		diagram[i] = make([]uint8, xMax+1)
	}

	for _, xy := range coords {
		x := xy[0]
		y := xy[1]
		diagram[y][x] += 1
	}

	var overlaps int
	for _, row := range diagram {
		fmt.Println(row)
		for _, col := range row {
			if col > 1 {
				overlaps++
			}
		}
	}

	fmt.Println(overlaps)
}

func GetMaxCoords(coords [][]int) (int, int) {
	var xMax int
	var yMax int

	for _, xy := range coords {
		x := xy[0]
		y := xy[1]
		if x > xMax {
			xMax = x
		}
		if y > yMax {
			yMax = y
		}
	}

	return xMax, yMax
}

func ParseCoordinates(lines []string) [][]int {
	var coords [][]int

	for _, line := range lines {
		split := strings.Split(line, " -> ")
		startXY := strings.Split(split[0], ",")
		endXY := strings.Split(split[1], ",")
		x1, _ := strconv.Atoi(startXY[0])
		y1, _ := strconv.Atoi(startXY[1])
		x2, _ := strconv.Atoi(endXY[0])
		y2, _ := strconv.Atoi(endXY[1])

		dx := GetStep(x2 - x1)
		dy := GetStep(y2 - y1)

		coords = append(coords, []int{x1, y1})
		for x1 != x2 || y1 != y2 {
			x1 += dx
			y1 += dy
			coords = append(coords, []int{x1, y1})
		}
	}

	return coords
}

func GetStep(d int) int {
	if d > 0 {
		return 1
	} else if d < 0 {
		return -1
	} else {
		return 0
	}
}
