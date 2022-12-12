package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const START = 'S'
const END = 'E'

func key(x, y int) string {
	return fmt.Sprintf("%d %d", x, y)
}

func tuple(key string) (int, int) {
	broken := strings.Split(key, " ")
	x, _ := strconv.Atoi(broken[0])
	y, _ := strconv.Atoi(broken[1])
	return x, y
}

func get_min_dist(queue []string, dist map[string]int) (int, int, int) {
	min_dist := 999
	var min_key string
	var min_index int
	for i, k := range queue {
		if dist[k] < min_dist {
			min_dist = dist[k]
			min_index = i
			min_key = k
			//fmt.Println("found lower", min_dist, min_index, min_key)
		}
	}
	x, y := tuple(min_key)
	return x, y, min_index
}

func dijkstra(mapa [][]int, x_s int, y_s int, x_e int, y_e int) int {
	dist := map[string]int{}
	queue := []string{}
	x_count := len(mapa)
	y_count := len(mapa[0])

	for i := 0; i < x_count; i++ {
		for j := 0; j < y_count; j++ {
			key := key(i, j)
			dist[key] = 999
			queue = append(queue, key)
			//fmt.Println(queue)
		}
	}
	dist[key(x_s, y_s)] = 0

	for true {
		if len(queue) == 0 {
			return -1 //shouldn't happen
		}
		//fmt.Println(queue)
		x, y, index := get_min_dist(queue, dist)
		queue = append(queue[:index], queue[index+1:]...)
		//fmt.Println(queue)
		if x == x_e && y == y_e {
			//fmt.Println("FINISHED!?")
			return dist[key(x, y)]
		}
		neighbors := []string{key(x-1, y), key(x, y-1), key(x+1, y), key(x, y+1)}
		for _, k := range neighbors {
			vx, vy := tuple(k)
			if vx < 0 || vy < 0 || vx >= x_count || vy >= y_count {
				// exclude invalid moves
				continue
			}

			my_height := mapa[x][y]
			if mapa[vx][vy] > my_height+1 {
				//climb too big
				continue
			}
			alt := dist[key(x, y)] + 1
			if alt < dist[k] {
				dist[k] = alt
			}
		}
		//fmt.Println(dist)
	}

	return -2 //shouldn't happen
}

func main() {

	input, _ := os.ReadFile("./input/input12.txt")

	lines := strings.Split(string(input), "\r\n")
	x_count := len(lines[0])
	y_count := len(lines)

	var x_s, y_s, x_e, y_e int

	mapa := make([][]int, y_count)

	for i, line := range lines {
		//fmt.Println(line)
		mapa[i] = make([]int, x_count)
		for j, r := range line {
			if r == START {
				mapa[i][j] = 0
				x_s = i
				y_s = j
				continue
			}
			if r == END {
				mapa[i][j] = 25
				x_e = i
				y_e = j
				continue
			}
			mapa[i][j] = int(r - 'a')
		}
	}

	part1 := dijkstra(mapa, x_s, y_s, x_e, y_e)

	fmt.Println("Part 1:", part1)
}
