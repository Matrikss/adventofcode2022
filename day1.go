package main

import (
	"fmt"
	"math"
	"os"
	"sort"
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

	total_per_elf := []int{}
	max_calories := 0.0
	calories := 0
	for _, line := range lines {
		if len(line) == 0 {
			max_calories = math.Max(float64(max_calories), float64(calories))
			total_per_elf = append(total_per_elf, calories)
			calories = 0
			continue
		}
		calorie, err := strconv.ParseInt(line, 10, 32)
		check(err)
		calories += int(calorie)
	}
	sort.Ints(total_per_elf)
	count := len(total_per_elf)
	fmt.Println("part 1: ", max_calories)
	fmt.Println("part 2: ", total_per_elf[count-1]+total_per_elf[count-2]+total_per_elf[count-3])
}
