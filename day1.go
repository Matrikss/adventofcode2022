package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	input, err := os.ReadFile("./input/input1.txt")
	check(err)

	lines := strings.Split(string(input), "\r\n")

	max_calories := 0.0
	calories := 0
	for _, line := range lines {
		if len(line) == 0 {
			max_calories = math.Max(float64(max_calories), float64(calories))
			calories = 0
			continue
		}
		calorie, err := strconv.ParseInt(line, 10, 32)
		check(err)
		calories += int(calorie)
	}
	fmt.Println(max_calories)
}
