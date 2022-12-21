package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_part1(monkeys map[string]string, monkey string) int {
	broken := strings.Split(monkeys[monkey], " ")
	if len(broken) == 1 {
		yell, _ := strconv.Atoi(broken[0])
		return yell
	} else {
		if broken[1] == "+" {
			return get_part1(monkeys, broken[0]) + get_part1(monkeys, broken[2])
		}
		if broken[1] == "-" {
			return get_part1(monkeys, broken[0]) - get_part1(monkeys, broken[2])
		}
		if broken[1] == "*" {
			return get_part1(monkeys, broken[0]) * get_part1(monkeys, broken[2])
		}
		return get_part1(monkeys, broken[0]) / get_part1(monkeys, broken[2])
	}
}

func main() {

	input, _ := os.ReadFile("./input/input21.txt")

	monkeys := map[string]string{}

	lines := strings.Split(string(input), "\r\n")
	for _, line := range lines {
		broken := strings.Split(line, ": ")
		monkeys[broken[0]] = broken[1]
	}

	fmt.Println("Part 1:", get_part1(monkeys, "root"))
	fmt.Println("Part 2:", 0)
}
