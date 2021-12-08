package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func readFileLines(filepath string, split string) []string {
	content, _ := ioutil.ReadFile(filepath)
	return strings.Split(string(content), split)
}

func main() {
	lines := readFileLines("./input", "\n")
	var positions []int
	for _, n := range strings.Split(lines[0], ",") {
		pos, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		positions = append(positions, pos)
	}

	fuels := make([]int, len(positions))
	for i := 0; i < len(positions); i++ {
		for _, pos := range positions {
			// part 1
			// diff := int(math.Abs(float64(pos - i)))
			// fuels[i] += diff

			// part 2
			for j := 1; j <= int(math.Abs(float64(pos-i))); j++ {
				fuels[i] += j
			}
		}
	}

	minValue := fuels[0]
	var minPos int
	for i, fuelValue := range fuels {
		if fuelValue < minValue {
			minValue = fuelValue
			minPos = i
		}
	}

	fmt.Println(minPos, minValue)

}
