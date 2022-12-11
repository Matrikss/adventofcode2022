package main

import (
	"fmt"
	"math"
	"sort"
)

type operation_throw func(int) (int, int)

type monkey struct {
	worry_levels []int
	operation    operation_throw
	inspections  int
}

func relax(worry int) int {
	return int(math.Floor(float64(worry) / 3))
}

func main() {

	monkeys := []monkey{
		monkey{
			[]int{79, 98},
			func(worry int) (int, int) {
				new := relax(worry * 19)
				if new%23 == 0 {
					return 2, new
				} else {
					return 3, new
				}
			}, 0},
		monkey{
			[]int{54, 65, 75, 74},
			func(worry int) (int, int) {
				new := relax(worry + 6)
				if new%19 == 0 {
					return 2, new
				} else {
					return 0, new
				}
			}, 0},
		monkey{
			[]int{79, 60, 97},
			func(worry int) (int, int) {
				new := relax(worry * worry)
				if new%13 == 0 {
					return 1, new
				} else {
					return 3, new
				}
			}, 0},
		monkey{
			[]int{74},
			func(worry int) (int, int) {
				new := relax(worry + 3)
				if new%17 == 0 {
					return 0, new
				} else {
					return 1, new
				}
			}, 0},
	}

	for round := 1; round <= 20; round++ {
		for i := 0; i < len(monkeys); i++ {
			items := monkeys[i].worry_levels
			monkeys[i].worry_levels = []int{}
			monkeys[i].inspections += len(items)
			for _, item := range items {
				throw_to, new_worry := monkeys[i].operation(item)
				//println("monkey", i, "throws", item, "as", new_worry, "to", throw_to)
				monkeys[throw_to].worry_levels = append(monkeys[throw_to].worry_levels, new_worry)
			}
		}
	}

	monkey_business := []int{}
	for _, m := range monkeys {
		monkey_business = append(monkey_business, m.inspections)
	}
	sort.Ints(monkey_business)
	count := len(monkey_business)

	fmt.Println("Part 1:", monkey_business[count-1]*monkey_business[count-2])
}
