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

func is_full_overlap(pa string, pb string) bool {
	pas := strings.Split(pa, "-")
	pbs := strings.Split(pb, "-")
	xa, _ := strconv.Atoi(pas[0])
	ya, _ := strconv.Atoi(pas[1])
	xb, _ := strconv.Atoi(pbs[0])
	yb, _ := strconv.Atoi(pbs[1])
	if xa <= xb && ya >= yb {
		return true
	}
	if xb <= xa && yb >= ya {
		return true
	}

	return false
}

func main() {

	input, err := os.ReadFile("./input/input4.txt")
	check(err)

	lines := strings.Split(string(input), "\r\n")

	full_overlaps := 0
	for _, line := range lines {
		pairs := strings.Split(line, ",")
		if is_full_overlap(pairs[0], pairs[1]) {
			full_overlaps++
		}
	}

	fmt.Println("part 1: ", full_overlaps)
	fmt.Println("part 2: ", 0)
}
