package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func print_screen(screen [40][6]string) {
	for j := 0; j < 6; j++ {
		for i := 0; i < 40; i++ {
			char := screen[i][j]
			fmt.Print(char)
		}
		fmt.Println()
	}
}

func main() {

	input, _ := os.ReadFile("./input/input10.txt")

	var x_per_cycle [241]int
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
	fmt.Println("Part 1: ", part1)

	var screen [40][6]string
	for i := 0; i < 240; i++ {
		x := i % 40
		y := i / 40
		if x_per_cycle[i] <= x+1 && x_per_cycle[i] >= x-1 {
			screen[x][y] = "#"
		} else {
			screen[x][y] = "_"
		}
	}

	fmt.Println("Part 2:")
	print_screen(screen)
}
