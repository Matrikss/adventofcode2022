package main

import (
	"fmt"
	"os"
	"strings"
)

// const MARKER = 4 // Part 1
const MARKER = 14 // Part 2

func main() {

	input, _ := os.ReadFile("./input/input6.txt")

	lines := strings.Split(string(input), "\r\n")

	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			seen := map[byte]int{line[i]: 1}
			for j := i + 1; j < i+MARKER; j++ {
				if _, ok := seen[line[j]]; ok {
					break
				}
				seen[line[j]] = 1
			}
			if len(seen) == MARKER {
				fmt.Println("Solution: ", i+MARKER)
				break
			}
		}
	}
}
