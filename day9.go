package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type rope struct {
	head_x         float64
	head_y         float64
	tail_x         float64
	tail_y         float64
	prev_hx        float64
	prev_hy        float64
	tail_positions map[string]int
}

func move_tail(corda *rope) {
	dist_x := math.Abs(corda.tail_x - corda.head_x)
	dist_y := math.Abs(corda.tail_y - corda.head_y)
	if dist_x > 1 || dist_y > 1 {
		corda.tail_x = corda.prev_hx
		corda.tail_y = corda.prev_hy
		corda.tail_positions[fmt.Sprintf("%d %d", int(corda.tail_x), int(corda.tail_y))] += 1
	}
}

func move_head(corda *rope, x_inc float64, y_inc float64) {
	abs_move := math.Abs(x_inc) + math.Abs(y_inc)
	for i := 0; i < int(abs_move); i++ {
		corda.prev_hx = corda.head_x
		corda.prev_hy = corda.head_y
		corda.head_x += x_inc / abs_move
		corda.head_y += y_inc / abs_move
		move_tail(corda)
	}
}

func main() {

	input, _ := os.ReadFile("./input/input9.txt")

	lines := strings.Split(string(input), "\r\n")

	corda := rope{0, 0, 0, 0, 0, 0, map[string]int{"0 0": 1}}
	for _, line := range lines {
		broken := strings.Split(line, " ")
		quantity, _ := strconv.Atoi(broken[1])
		if broken[0] == "R" {
			move_head(&corda, float64(quantity), 0)
		}
		if broken[0] == "L" {
			move_head(&corda, -float64(quantity), 0)
		}
		if broken[0] == "U" {
			move_head(&corda, 0, -float64(quantity))
		}
		if broken[0] == "D" {
			move_head(&corda, 0, float64(quantity))
		}
	}

	fmt.Println("Part 1: ", len(corda.tail_positions))
	fmt.Println("Part 2: ", 0)
}
