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
			[]int{56, 52, 58, 96, 70, 75, 72},
			func(worry int) (int, int) {
				new := relax(worry * 17)
				if new%11 == 0 {
					return 2, new
				} else {
					return 3, new
				}
			}, 0},
		monkey{
			[]int{75, 58, 86, 80, 55, 81},
			func(worry int) (int, int) {
				new := relax(worry + 7)
				if new%3 == 0 {
					return 6, new
				} else {
					return 5, new
				}
			}, 0},
		monkey{
			[]int{73, 68, 73, 90},
			func(worry int) (int, int) {
				new := relax(worry * worry)
				if new%5 == 0 {
					return 1, new
				} else {
					return 7, new
				}
			}, 0},
		monkey{
			[]int{72, 89, 55, 51, 59},
			func(worry int) (int, int) {
				new := relax(worry + 1)
				if new%7 == 0 {
					return 2, new
				} else {
					return 7, new
				}
			}, 0},
		monkey{
			[]int{76, 76, 91},
			func(worry int) (int, int) {
				new := relax(worry * 3)
				if new%19 == 0 {
					return 0, new
				} else {
					return 3, new
				}
			}, 0},
		monkey{
			[]int{88},
			func(worry int) (int, int) {
				new := relax(worry + 4)
				if new%2 == 0 {
					return 6, new
				} else {
					return 4, new
				}
			}, 0},
		monkey{
			[]int{64, 63, 56, 50, 77, 55, 55, 86},
			func(worry int) (int, int) {
				new := relax(worry + 8)
				if new%13 == 0 {
					return 4, new
				} else {
					return 0, new
				}
			}, 0},
		monkey{
			[]int{79, 58},
			func(worry int) (int, int) {
				new := relax(worry + 6)
				if new%17 == 0 {
					return 1, new
				} else {
					return 5, new
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
