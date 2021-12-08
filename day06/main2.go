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

	ages := make(map[int]int)
	for _, n := range strings.Split(lines[0], ",") {
		age, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		ages[age]++
	}

	var day int
	for day < 256 {
		zeros := ages[0]
		for i := 1; i <= 9; i++ {
			ages[i-1] = ages[i]
		}
		ages[6] += zeros
		ages[8] += zeros

		day++
	}

	var sum int
	for _, n := range ages {
		sum += n
	}
	fmt.Println(sum)
}
