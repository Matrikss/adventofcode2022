package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const TRIM = 450 // 0 for final solution!!!!!

func print_map(mapa [][]string) {
	for j := 0; j < len(mapa[0]); j++ {
		for i := TRIM; i < len(mapa); i++ {
			char := mapa[i][j]
			if char == "" {
				char = "."
			}
			fmt.Print(char)
		}
		fmt.Println()
	}
}

func draw_line(mapa [][]string, xa, ya, xb, yb string) [][]string {
	xai, _ := strconv.Atoi(xa)
	yai, _ := strconv.Atoi(ya)
	xbi, _ := strconv.Atoi(xb)
	ybi, _ := strconv.Atoi(yb)

	for j := yai; j <= ybi; j++ {
		for i := xbi; i <= xai; i++ {
			mapa[i][j] = "#"
		}
	}
	return mapa
}

func add_sand(mapa [][]string) bool {
	mapa[500][0] = "+"
	return true //if it goes off the map
}

func main() {

	input, _ := os.ReadFile("./input/input14.txt")

	x_count := 504
	y_count := 10
	mapa := make([][]string, x_count)
	for i := 0; i < x_count; i++ {
		mapa[i] = make([]string, y_count)
	}

	lines := strings.Split(string(input), "\r\n")
	for _, line := range lines {
		broken := strings.Split(line, " -> ")
		for i := 0; i < len(broken)-1; i++ {
			a := strings.Split(broken[i], ",")
			b := strings.Split(broken[i+1], ",")
			mapa = draw_line(mapa, a[0], a[1], b[0], b[1])
		}
	}

	part1 := 0
	finished := false
	for true {
		finished = add_sand(mapa)
		if finished {
			break
		}
		part1++
	}
	print_map(mapa)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", 0)
}
