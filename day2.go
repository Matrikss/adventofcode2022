package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	combination := map[string]int{
		"A X": 1 + 3,
		"A Y": 2 + 6,
		"A Z": 3 + 0,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 1 + 6,
		"C Y": 2 + 0,
		"C Z": 3 + 3,
	}

	input, err := os.ReadFile("./input/input2.txt")
	check(err)

	lines := strings.Split(string(input), "\r\n")

	points := 0
	for _, line := range lines {
		points += combination[line]
	}

	fmt.Println("part 1: ", points)
	fmt.Println("part 2: ", 0)
}
