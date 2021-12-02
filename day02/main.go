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

	var directions []string
	var speed []int

	for _, line := range lines {
		split := strings.Split(line, " ")
		directions = append(directions, split[0])
		v, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
		speed = append(speed, v)
	}

	func() {
		var x int
		var z int
		for i, direction := range directions {
			switch direction {
			case "up":
				z -= speed[i]
			case "down":
				z += speed[i]
			case "forward":
				x += speed[i]
			}
		}

		fmt.Println("Part 1")
		fmt.Printf("X: %d\n", x)
		fmt.Printf("Z: %d\n", z)
		fmt.Printf("X * Z: %d\n", x*z)
	}()

	fmt.Println()

	func() {
		var x int
		var z int
		var aim int
		for i, direction := range directions {
			switch direction {
			case "up":
				aim -= speed[i]
			case "down":
				aim += speed[i]
			case "forward":
				x += speed[i]
				z += aim * speed[i]
			}
		}

		fmt.Println("Part 2")
		fmt.Printf("X: %d\n", x)
		fmt.Printf("Z: %d\n", z)
		fmt.Printf("X * Z: %d\n", x*z)
	}()

}
