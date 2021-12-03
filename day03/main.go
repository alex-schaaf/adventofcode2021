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
	lines := readFileLines("./test", "\n")

	func() {
		positions := len(lines[0])
		countZeros := make([]int, positions)
		countOnes := make([]int, positions)

		for _, line := range lines {
			for i, bit := range line {
				if bit == '0' {
					countZeros[i] += 1
				} else {
					countOnes[i] += 1
				}
			}
		}

		gammaRate := make([]rune, positions)
		epsilonRate := make([]rune, positions)

		for i, zeros := range countZeros {
			if zeros > countOnes[i] {
				gammaRate[i] = '0'
				epsilonRate[i] = '1'
			} else {
				gammaRate[i] = '1'
				epsilonRate[i] = '0'
			}
		}

		gamma, _ := strconv.ParseInt(string(gammaRate), 2, 64)
		epsilon, _ := strconv.ParseInt(string(epsilonRate), 2, 64)

		fmt.Printf("Part 1: %d\n", gamma*epsilon)
	}()

}
