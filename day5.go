package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	crates := [][]string{
		{},
		{"S", "Z", "P", "D", "L", "B", "F", "C"},
		{"N", "V", "G", "P", "H", "W", "B"},
		{"F", "W", "B", "J", "G"},
		{"G", "J", "N", "F", "L", "W", "C", "S"},
		{"W", "J", "L", "T", "P", "M", "S", "H"},
		{"B", "C", "W", "G", "F", "S"},
		{"H", "T", "P", "M", "Q", "B", "W"},
		{"F", "S", "W", "T"},
		{"N", "C", "R"},
	}

	input, err := os.ReadFile("./input/input5.txt")
	check(err)

	lines := strings.Split(string(input), "\r\n")

	for _, line := range lines {
		instructions := strings.Split(line, " ")
		moves, _ := strconv.Atoi(instructions[1])
		from, _ := strconv.Atoi(instructions[3])
		to, _ := strconv.Atoi(instructions[5])
		if false { // part 1
			for i := 0; i < moves; i++ {
				pop := len(crates[from]) - 1
				crates[to] = append(crates[to], crates[from][pop])
				crates[from] = crates[from][:pop]
			}
		} else { // part 2
			pop := len(crates[from])
			crates[to] = append(crates[to], crates[from][pop-moves:pop]...)
			crates[from] = crates[from][:pop-moves]
		}
	}

	result := ""
	for i := 1; i < len(crates); i++ {
		pop := len(crates[i]) - 1
		result += crates[i][pop]
	}

	fmt.Println("solution: ", result)
}
