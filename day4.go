package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func get_range(pair string) (int, int) {
	pas := strings.Split(pair, "-")
	xa, _ := strconv.Atoi(pas[0])
	ya, _ := strconv.Atoi(pas[1])
	return xa, ya
}

func is_full_overlap(pa string, pb string) bool {
	xa, ya := get_range(pa)
	xb, yb := get_range(pb)
	if xa <= xb && ya >= yb {
		return true
	}
	if xb <= xa && yb >= ya {
		return true
	}

	return false
}

func is_overlap(pa string, pb string) bool {
	xa, ya := get_range(pa)
	xb, yb := get_range(pb)
	if xa <= xb && xb <= ya {
		return true
	}
	if xa <= yb && yb <= ya {
		return true
	}
	if xb <= xa && xa <= yb {
		return true
	}
	if xb <= ya && ya <= yb {
		return true
	}
	return false
}

func main() {

	input, err := os.ReadFile("./input/input4.txt")
	check(err)

	lines := strings.Split(string(input), "\r\n")

	full_overlaps := 0
	overlaps := 0
	for _, line := range lines {
		pairs := strings.Split(line, ",")
		if is_full_overlap(pairs[0], pairs[1]) {
			full_overlaps++
		}
		if is_overlap(pairs[0], pairs[1]) {
			overlaps++
		}
	}

	fmt.Println("part 1: ", full_overlaps)
	fmt.Println("part 2: ", overlaps)
}
