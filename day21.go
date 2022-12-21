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

const FOUND = -1

func get_part2(monkeys map[string]string, monkey string, humn int) int {
	if monkey == "humn" {
		return humn
	}
	broken := strings.Split(monkeys[monkey], " ")
	if len(broken) == 1 {
		yell, _ := strconv.Atoi(broken[0])
		return yell
	} else {
		if monkey == "root" {
			a := get_part2(monkeys, broken[0], humn)
			b := get_part2(monkeys, broken[2], humn)
			if a == b {
				fmt.Println(get_part2(monkeys, broken[0], humn), "==", get_part2(monkeys, broken[2], humn))
				return FOUND
			}
			//fmt.Println("\n", get_part2(monkeys, broken[0], humn), "!=\n", get_part2(monkeys, broken[2], humn))
			return 0
		}
		if broken[1] == "+" {
			return get_part2(monkeys, broken[0], humn) + get_part2(monkeys, broken[2], humn)
		}
		if broken[1] == "-" {
			return get_part2(monkeys, broken[0], humn) - get_part2(monkeys, broken[2], humn)
		}
		if broken[1] == "*" {
			return get_part2(monkeys, broken[0], humn) * get_part2(monkeys, broken[2], humn)
		}
		return get_part2(monkeys, broken[0], humn) / get_part2(monkeys, broken[2], humn)
	}
}

func find_part2(monkeys map[string]string, start int) {
	part2 := start
	var res int
	for {
		res = get_part2(monkeys, "root", part2)
		if res == FOUND {
			break
		}
		part2++
	}
	fmt.Println("Part 2:", part2)
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

	find_part2(monkeys, 3759566890000) // 3759566892641 👀🤚
}
