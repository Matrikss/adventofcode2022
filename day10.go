package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, _ := os.ReadFile("./input/input10.txt")

	var x_per_cycle [250]int
	x_per_cycle[0] = 1

	lines := strings.Split(string(input), "\r\n")

	i := 1
	for _, line := range lines {
		if line == "noop" {
			x_per_cycle[i] = x_per_cycle[i-1]
		} else {
			broken := strings.Split(line, " ")
			inc, _ := strconv.Atoi(broken[1])
			x_per_cycle[i] = x_per_cycle[i-1]
			i++
			x_per_cycle[i] = x_per_cycle[i-1] + inc
		}
		i++
	}

	part1 := 0
	for i := 20; i < 225; i += 40 {
		part1 += i * x_per_cycle[i-1]
	}
	//fmt.Println(x_per_cycle)
	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", 0)
}
