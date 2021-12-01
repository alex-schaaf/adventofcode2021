package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readFileLines(filepath string, split string) []string {
	content, _ := ioutil.ReadFile(filepath)
	return strings.Split(string(content), split)
}

func main() {
	lines := readFileLines("./input", "\n")

	depths, err := StringsToNumbers(lines)
	if err != nil {
		log.Fatal(err)
	}

	nIncreases := CountIncreases(depths)
	fmt.Printf("Part 1: # Measurements larger than the previous: %d", nIncreases)
	fmt.Println()
	nIncreasesWindow := CountIncreasesSlidingWindow(depths)
	fmt.Printf("Part 2: # Measurements larger than the previous: %d", nIncreasesWindow)
}

func StringsToNumbers(strings []string) ([]int, error) {
	var depths []int

	for _, str := range strings {
		depth, err := strconv.Atoi(str)
		if err != nil {
			return depths, err
		}
		depths = append(depths, depth)
	}
	return depths, nil
}

func CountIncreases(depths []int) int {
	var counter int
	for i := 1; i < len(depths); i++ {
		d0 := depths[i-1]
		d1 := depths[i]

		if d1 > d0 {
			counter++
		}
	}
	return counter
}

func CountIncreasesSlidingWindow(depths []int) int {
	var counter int
	windowSize := 3
	for i := windowSize; i < len(depths); i++ {
		dA0 := depths[i-windowSize]
		dA1 := depths[i-windowSize+1]
		dA2 := depths[i-windowSize+2]
		sumA := dA0 + dA1 + dA2

		dB0 := depths[i-windowSize+1]
		dB1 := depths[i-windowSize+2]
		dB2 := depths[i-windowSize+3]
		sumB := dB0 + dB1 + dB2

		if sumB > sumA {
			counter++
		}
	}
	return counter
}
