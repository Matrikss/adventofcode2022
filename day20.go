package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	x    int
	prev *Number
	next *Number
}

const SIZE = 7

func print_sequence(seq []*Number, start int, end int) {
	now := seq[start]
	for i := 0; i < SIZE*2; i++ {
		fmt.Print(now.x, " ")
		now = now.next
	}
	fmt.Println()
	now = seq[end]
	for i := 0; i < SIZE*2; i++ {
		fmt.Print(now.x, " ")
		now = now.prev
	}
	fmt.Println()
}

func get_value_for_position(seq []*Number, start int, length int) int {
	now := seq[start]
	for i := 0; i < length; i++ {
		now = now.next
	}
	return now.x
}

func main() {

	input, _ := os.ReadFile("./input/input20.txt")

	var sequencia [SIZE]*Number
	var zero_pos int

	lines := strings.Split(string(input), "\r\n")
	for i, line := range lines {
		x, _ := strconv.Atoi(line)
		if x == 0 {
			zero_pos = i
		}
		sequencia[i] = &Number{x, nil, nil}
	}
	for i, _ := range sequencia {
		prev := i - 1
		if prev < 0 {
			prev = SIZE - 1
		}
		sequencia[i].next = sequencia[(i+1)%SIZE]
		sequencia[i].prev = sequencia[prev]
	}

	for i := 0; i < SIZE; i++ {

		move := sequencia[i].x
		starting_pos := sequencia[i]
		new_pos := sequencia[i]
		before_start := starting_pos.prev
		after_start := starting_pos.next
		if move > 0 {
			for {
				if move > 0 {
					new_pos = new_pos.next
					move--
				} else {
					break
				}
			}
			if new_pos == starting_pos {
				fmt.Println("Edge Case?")
				continue
			}
			if new_pos == starting_pos.prev {
				fmt.Println("Edge Case?")
				continue
			}
			after_new_pos := new_pos.next

			starting_pos.next = after_new_pos
			starting_pos.prev = new_pos

			new_pos.next = starting_pos
			if new_pos.prev == starting_pos {
				new_pos.prev = before_start
			}

			after_new_pos.prev = starting_pos

			before_start.next = after_start
			after_start.prev = before_start
		} else if move < 0 {
			move--
			for {
				if move < 0 {
					new_pos = new_pos.prev
					move++
				} else {
					break
				}
			}
			if new_pos == starting_pos {
				fmt.Println("Edge Case?")
				continue
			}
			if new_pos == starting_pos.prev {
				fmt.Println("Edge Case?")
				continue
			}
			after_new_pos := new_pos.next

			starting_pos.next = after_new_pos
			starting_pos.prev = new_pos

			new_pos.next = starting_pos

			after_new_pos.prev = starting_pos

			before_start.next = after_start
			after_start.prev = before_start
		}
	}

	//print_sequence(sequencia[:], 0, 4)

	part1 := get_value_for_position(sequencia[:], zero_pos, 1000)
	part1 += get_value_for_position(sequencia[:], zero_pos, 2000)
	part1 += get_value_for_position(sequencia[:], zero_pos, 3000)

	fmt.Println("Part 1:", part1) // 7238 too low // correct according to HAM: 7584
	fmt.Println("Part 2:", 0)
}
