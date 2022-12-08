package main

import (
	"fmt"
	"os"
	"strings"
)

func is_visible(mapa [][]int, height, x int, y int, x_inc int, y_inc int) bool {
	side := len(mapa)
	if x <= 0 || x >= side || x+x_inc >= side {
		return true
	}
	if y <= 0 || y >= side || y+y_inc >= side {
		return true
	}
	if mapa[x+x_inc][y+y_inc] < height {
		return is_visible(mapa, height, x+x_inc, y+y_inc, x_inc, y_inc)
	}
	return false
}

func main() {

	input, _ := os.ReadFile("./input/input8.txt")

	lines := strings.Split(string(input), "\r\n")
	side_len := len(lines[0])
	mapa := make([][]int, side_len)

	for i, line := range lines {
		mapa[i] = make([]int, side_len)
		for j, r := range line {
			mapa[i][j] = int(r - '0')
		}
	}

	visible := 0
	for i := 1; i < side_len-1; i++ {
		for j := 1; j < side_len-1; j++ {
			height := mapa[i][j]
			if is_visible(mapa, height, i, j, -1, 0) || is_visible(mapa, height, i, j, 0, -1) || is_visible(mapa, height, i, j, 1, 0) || is_visible(mapa, height, i, j, 0, 1) {
				visible++
			}
		}
	}

	fmt.Println("Part 1: ", side_len*4-4+visible)
}
