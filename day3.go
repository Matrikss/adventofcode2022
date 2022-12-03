package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	priorities := map[string]int{}

	j := 1
	for i := 'a'; i <= 'z'; {
		priorities[fmt.Sprintf("%c", i)] = j
		i++
		j++
	}

	for i := 'A'; i <= 'Z'; {
		priorities[fmt.Sprintf("%c", i)] = j
		i++
		j++
	}

	input, err := os.ReadFile("./input/input3.txt")
	check(err)

	lines := strings.Split(string(input), "\r\n")

	sum := 0
	for _, line := range lines {
		f_half := line[:len(line)/2]
		s_half := line[len(line)/2:]

		regex := regexp.MustCompile("[" + f_half + "]")
		sum += priorities[regex.FindStringSubmatch(s_half)[0]]
	}

	fmt.Println("part 1: ", sum)
	fmt.Println("part 2: ", 0)
}
