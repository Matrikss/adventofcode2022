package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type node struct {
	size     int
	name     string
	parent   *node
	children []*node
}

func get_child_node(parent *node, folder string) *node {
	for _, no := range parent.children {
		if no.name == folder {
			return no
		}
	}
	return nil
}

func calculate_folder_sizes(this *node) int {
	res := 0
	if len(this.children) == 0 {
		res = this.size
	} else {
		for _, no := range this.children {
			res += calculate_folder_sizes(no)
		}
		this.size = res
	}
	return res
}

func main() {

	input, err := os.ReadFile("./input/input7.txt")
	check(err)

	lines := strings.Split(string(input), "\r\n")

	root := node{0, "/", nil, []*node{}}
	dirs := []*node{&root}
	current_dir := &root
	for _, line := range lines {
		if line[0] == '$' {
			if line == "$ ls" {
				// NOOP
			} else if line[5] == '.' {
				current_dir = current_dir.parent
			} else {
				current_dir = get_child_node(current_dir, strings.Split(line, " ")[2])
			}
		} else {
			broken := strings.Split(line, " ")
			size := 0
			if broken[0] != "dir" {
				size, _ = strconv.Atoi(broken[0])
			}
			new_node := node{size, broken[1], current_dir, []*node{}}
			current_dir.children = append(current_dir.children, &new_node)
			if broken[0] == "dir" {
				dirs = append(dirs, &new_node)
			}
		}
	}
	calculate_folder_sizes(&root)

	part1 := 0
	for _, no := range dirs {
		if no.size <= 100000 {
			part1 += no.size
		}
	}

	fmt.Println("Part 1: ", part1)

	need_to_free := 30000000 - (70000000 - root.size)

	min := 30000000
	for _, no := range dirs {
		if no.size >= need_to_free {
			if no.size < min {
				min = no.size
			}
		}
	}

	fmt.Println("Part 2: ", min)
}
