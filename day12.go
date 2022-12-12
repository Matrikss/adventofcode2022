package main

import (
	"fmt"
	"os"
	"strings"
)

const START = 'S'
const END = 'E'

func dijkstra(mapa [][]int, x_s int, y_s int, x_e int, y_e int) int {

	return 0
}

func main() {

	input, _ := os.ReadFile("./input/input12.txt")

	lines := strings.Split(string(input), "\n") // generated on Mac
	x_count := len(lines[0])
	y_count := len(lines)

	var x_s, y_s, x_e, y_e int

	mapa := make([][]int, y_count)

	for i, line := range lines {
		fmt.Println(line)
		mapa[i] = make([]int, x_count)
		for j, r := range line {
			if r == START {
				mapa[i][j] = 0
				x_s = i
				y_s = j
				continue
			}
			if r == END {
				mapa[i][j] = 25
				x_e = 0
				y_e = 0
				continue
			}
			mapa[i][j] = int(r - 'a')
		}
	}

	fmt.Println(mapa)

	part1 := dijkstra(mapa, x_s, y_s, x_e, y_e)

	fmt.Println("Part 1:", part1)
}
