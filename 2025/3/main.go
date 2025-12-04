package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	lines := readInput()
	fmt.Println(solve1(lines))
}

// find the max of 2 batteries combined
func findMax(line string) int {
	max := 0

	for i := 0; i < len(line); i++ {
		for j := i + 1; j < len(line); j++ {
			combined := string(line[i]) + string(line[j])
			num, err := strconv.Atoi(combined)
			if err != nil {
				panic(err)
			}
			if num > max {
				max = num
			}
		}
	}
	return max
}

func solve1(lines []string) int {
	t := 0

	for _, line := range lines {
		t += findMax(line)
	}

	return t
}

func readInput() []string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}
