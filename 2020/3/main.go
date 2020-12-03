package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	right = 3
)

func main() {
	fmt.Printf("%v\n", solve1())
}

func solve1() int {
	trees := 0
	xpos := 0

	input := getInput()
	for _, line := range input {
		if string(line[xpos]) == "#" {
			trees++
		}
		xpos = (xpos + right) % len(line)
	}
	return trees
}

func getInput() []string {
	in, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(in), "\n")
	return lines[:len(lines)-1]
}
