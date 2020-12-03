package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Printf("%v\n", solve1())
	fmt.Printf("%v\n", solve2())
}

func traverse(down, right int, input []string) int {
	trees := 0
	xpos := 0

	for i := 0; i < len(input); i += down {
		line := input[i]
		if string(line[xpos]) == "#" {
			trees++
		}
		xpos = (xpos + right) % len(line)
	}

	return trees
}

func solve1() int {
	input := getInput()
	return traverse(1, 3, input)
}

func solve2() int {
	input := getInput()

	slopes := []struct {
		right, down int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	mult := 1
	for _, slope := range slopes {
		mult *= traverse(slope.down, slope.right, input)
	}

	return mult
}

func getInput() []string {
	in, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(in), "\n")
	return lines[:len(lines)-1]
}
