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

	input, err := os.ReadFile("./input/input6.txt")
	check(err)

	lines := strings.Split(string(input), "\r\n")

	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			if line[i] != line[i+1] && line[i+1] != line[i+2] && line[i+2] != line[i+3] && line[i] != line[i+2] && line[i+1] != line[i+3] && line[i] != line[i+3] {
				fmt.Println("Part1: ", i+4)
				break
			}
		}
	}
}
