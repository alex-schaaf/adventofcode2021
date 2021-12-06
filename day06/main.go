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
	var ages []int
	for _, n := range strings.Split(lines[0], ",") {
		age, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		ages = append(ages, age)
	}

	var day int
	var newAges []int
	for day < 256 {
		// fmt.Println(day, ages)
		for i, age := range ages {
			if age == 0 {
				ages[i] = 6
				newAges = append(newAges, 8)
			} else {
				ages[i]--
			}
			newAges = append(newAges, ages[i])
		}
		day++
		ages = newAges
		newAges = []int{}
	}
	fmt.Println("Part 1: #Fish: ", len(ages))

}
