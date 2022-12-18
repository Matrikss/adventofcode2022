package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	x int
	y int
	z int
}

func is_adjacent(a Cube, b Cube) bool {
	x_dist := int(math.Abs(float64(a.x) - float64(b.x)))
	y_dist := int(math.Abs(float64(a.y) - float64(b.y)))
	z_dist := int(math.Abs(float64(a.z) - float64(b.z)))
	if x_dist+y_dist+z_dist == 1 {
		return true
	}
	return false
}

func main() {

	input, _ := os.ReadFile("./input/input18.txt")

	var cubes []Cube

	lines := strings.Split(string(input), "\r\n")
	for _, line := range lines {
		broken := strings.Split(line, ",")
		x, _ := strconv.Atoi(broken[0])
		y, _ := strconv.Atoi(broken[1])
		z, _ := strconv.Atoi(broken[2])
		cubes = append(cubes, Cube{x, y, z})
	}

	//fmt.Println(cubes)

	part1 := 0
	adjacent_faces := 0
	for i := 0; i < len(cubes); i++ {
		adjacent_faces = 0
		for j := 0; j < len(cubes); j++ {
			if i == j {
				continue
			}
			if is_adjacent(cubes[i], cubes[j]) {
				adjacent_faces++
			}
		}
		part1 += 6 - adjacent_faces
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", 0)
}
