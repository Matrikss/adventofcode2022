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

func join_multi(s [][]string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			res += s[i][j]
		}
	}
	return res
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

	sum1 := 0
	for _, line := range lines {
		f_half := line[:len(line)/2]
		s_half := line[len(line)/2:]

		regex := regexp.MustCompile("[" + f_half + "]")
		sum1 += priorities[regex.FindStringSubmatch(s_half)[0]]
	}

	sum2 := 0
	for i := 0; i < len(lines); i += 3 {
		regex := regexp.MustCompile("[" + lines[i] + "]")
		common := regex.FindAllStringSubmatch(lines[i+1], -1)
		regex = regexp.MustCompile("[" + lines[i+2] + "]")
		common = regex.FindAllStringSubmatch(join_multi(common), -1)
		sum2 += priorities[common[0][0]]
	}

	fmt.Println("part 1: ", sum1)
	fmt.Println("part 2: ", sum2)
}
