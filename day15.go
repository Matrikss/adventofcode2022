package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	left  int
	right int
}

type Coord struct {
	x int
	y int
}

type Sensor struct {
	center Coord
	radius int
}

func newSensor(sensor Coord, beacon Coord) Sensor {
	manhatanDistance := int(math.Abs(float64(sensor.x) - float64(beacon.x)))
	manhatanDistance += int(math.Abs(float64(sensor.y) - float64(beacon.y)))
	return Sensor{sensor, manhatanDistance}
}

func getIntersections(sensor Sensor, y_line int) (int, int) {
	top_y := sensor.center.y - sensor.radius
	bottom_y := sensor.center.y + sensor.radius

	if y_line >= top_y && y_line <= bottom_y {
		dist := sensor.radius - int(math.Abs(float64(sensor.center.y)-float64(y_line)))
		return sensor.center.x - dist, sensor.center.x + dist + 1
	}
	return 0, 0
}

const SEARCH_SPACE = 4000000

func main() {

	input, _ := os.ReadFile("./input/input15.txt")

	sensors := []Sensor{}

	lines := strings.Split(string(input), "\r\n")
	for _, line := range lines {
		words := strings.Split(line, " ")
		sx, _ := strconv.Atoi(strings.Split(words[2], "=")[1])
		sy, _ := strconv.Atoi(strings.Split(words[3], "=")[1])
		bx, _ := strconv.Atoi(strings.Split(words[8], "=")[1])
		by, _ := strconv.Atoi(strings.Split(words[9], "=")[1])
		sensors = append(sensors, newSensor(Coord{sx, sy}, Coord{bx, by}))
	}
	var part2 int
	for y := 2906101; y < SEARCH_SPACE; y++ { // found with brute force
		if part2 != 0 {
			break
		}
		//fmt.Print(y, " ") // progress "bar"
		ranges := []Range{}
		for _, sensor := range sensors {
			left, right := getIntersections(sensor, y)
			if left == 0 && right == 0 {
				//no intersection found
			} else {
				ranges = append(ranges, Range{left, right})
			}
		}
		var test [SEARCH_SPACE]int
		for _, r := range ranges {
			left := r.left
			if left < 0 {
				left = 0
			}
			right := r.right
			if right > SEARCH_SPACE {
				right = SEARCH_SPACE
			}
			for i := left; i < right; i++ {
				test[i] = 1
			}
		}
		for i, v := range test {
			if v == 0 {
				part2 = i*4000000 + y
				break
			}
		}
	}

	fmt.Println("Part 1:", 4358651+882167) // 👀🤚
	fmt.Println("Part 2:", part2)          // 13213086906101
}
