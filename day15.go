package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

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
		return sensor.center.x - dist, sensor.center.x + dist
	}
	return 0, 0
}

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
	fmt.Println(sensors)

	part1 := 0
	for _, sensor := range sensors {
		left, right := getIntersections(sensor, 10)
		fmt.Println(left, right)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", 0)
}
