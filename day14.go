package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const TRIM = 300
const X_SAND_DROP = 500

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

	if yai > ybi {
		t := yai
		yai = ybi
		ybi = t
	}

	if xai > xbi {
		t := xai
		xai = xbi
		xbi = t
	}

	for j := yai; j <= ybi; j++ {
		for i := xai; i <= xbi; i++ {
			mapa[i][j] = "#"
		}
	}
	return mapa
}

func add_sand(mapa [][]string) bool {
	x_size := len(mapa)
	y_size := len(mapa[0])
	x := X_SAND_DROP
	if mapa[x][0] != "" {
		return true //if it reaches the source
	}
	for y := 0; y < y_size-1; y++ {
		if mapa[x][y+1] != "" {
			if x-1 >= 0 && mapa[x-1][y+1] == "" {
				x--
				continue
			}
			if x+1 < x_size && mapa[x+1][y+1] == "" {
				x++
				continue
			}
			mapa[x][y] = "o"
			return false
		}
	}
	return true //if it goes off the map
}

func main() {

	input, _ := os.ReadFile("./input/input14.txt")

	// didn't feel like parsing twice 😅:
	x_count := 700
	y_count := 163
	y_floor := 162
	mapa1 := make([][]string, x_count)
	for i := 0; i < x_count; i++ {
		mapa1[i] = make([]string, y_count)
	}

	lines := strings.Split(string(input), "\r\n")
	for _, line := range lines {
		broken := strings.Split(line, " -> ")
		for i := 0; i < len(broken)-1; i++ {
			a := strings.Split(broken[i], ",")
			b := strings.Split(broken[i+1], ",")
			mapa1 = draw_line(mapa1, a[0], a[1], b[0], b[1])
		}
	}

	mapa2 := make([][]string, x_count)
	for i := 0; i < x_count; i++ {
		mapa2[i] = make([]string, y_count)
		copy(mapa2[i], mapa1[i])
	}

	part1 := 0
	finished := false
	for {
		finished = add_sand(mapa1)
		if finished {
			break
		}
		part1++
	}

	draw_line(mapa2, strconv.Itoa(0), strconv.Itoa(y_floor), strconv.Itoa(x_count-1), strconv.Itoa(y_floor)) // 🤦
	part2 := 0
	finished = false
	for {
		finished = add_sand(mapa2)
		if finished {
			break
		}
		part2++
	}
	//print_map(mapa2)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
