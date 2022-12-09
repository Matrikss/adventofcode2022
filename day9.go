package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const CORDA = 10
const START_X = 0
const START_Y = 0

type knot struct {
	head_x         float64
	head_y         float64
	next           *knot
	tail_positions map[string]int
}

func move_tail(corda *knot, prev *knot) {
	dist_x := math.Abs(prev.head_x - corda.head_x)
	dist_y := math.Abs(prev.head_y - corda.head_y)
	if dist_x > 1 || dist_y > 1 {
		if prev.head_x != corda.head_x && prev.head_y != corda.head_y {
			corda.head_x += (prev.head_x - corda.head_x) / dist_x
			corda.head_y += (prev.head_y - corda.head_y) / dist_y
		} else {
			if dist_x > 1 {
				corda.head_x += (prev.head_x - corda.head_x) / dist_x
			}
			if dist_y > 1 {
				corda.head_y += (prev.head_y - corda.head_y) / dist_y
			}
		}
		corda.tail_positions[fmt.Sprintf("%d %d", int(corda.head_x), int(corda.head_y))] += 1
	}
	if corda.next != nil {
		move_tail(corda.next, corda)
	}
}

func move_head(corda *knot, x_inc float64, y_inc float64) {
	abs_move := math.Abs(x_inc) + math.Abs(y_inc)
	for i := 0; i < int(abs_move); i++ {
		corda.head_x += x_inc / abs_move
		corda.head_y += y_inc / abs_move
		move_tail(corda.next, corda)
	}
}

func print_visits(positions map[string]int) {
	fmt.Println(positions)
	for i := 0; i < START_X*2; i++ {
		for j := 0; j < START_Y*2; j++ {
			val := positions[fmt.Sprintf("%d %d", i, j)]
			if val > 0 {
				fmt.Print("#")
			} else {
				fmt.Print("_")
			}
		}
		fmt.Println()
	}
}

func print_snapshot(corda [10]*knot) {
	fmt.Println("snapshot")
	var mapa [START_X * 2][START_Y * 2]string
	for i := CORDA - 1; i >= 0; i-- {
		mapa[int(corda[i].head_x)][int(corda[i].head_y)] = fmt.Sprintf("%d", i)
	}
	for i := 0; i < START_X*2; i++ {
		for j := 0; j < START_Y*2; j++ {
			char := mapa[i][j]
			if char == "" {
				fmt.Print("_")
			} else {
				fmt.Print(char)
			}
		}
		fmt.Println()
	}
}

func main() {

	input, _ := os.ReadFile("./input/input9.txt")

	lines := strings.Split(string(input), "\r\n")

	var corda [CORDA]*knot
	for i := 0; i < CORDA; i++ {
		new := knot{START_X, START_Y, nil, map[string]int{fmt.Sprintf("%d %d", START_X, START_Y): 1}}
		corda[i] = &new
	}
	for i := 0; i < CORDA-1; i++ {
		corda[i].next = corda[i+1]
	}
	for _, line := range lines {
		broken := strings.Split(line, " ")
		quantity, _ := strconv.Atoi(broken[1])
		if broken[0] == "R" {
			move_head(corda[0], float64(quantity), 0)
		}
		if broken[0] == "L" {
			move_head(corda[0], -float64(quantity), 0)
		}
		if broken[0] == "U" {
			move_head(corda[0], 0, -float64(quantity))
		}
		if broken[0] == "D" {
			move_head(corda[0], 0, float64(quantity))
		}
	}

	fmt.Println("Part 1: ", len(corda[9].tail_positions)) //6498
	fmt.Println("Part 2: ", 0)                            //2531
}
