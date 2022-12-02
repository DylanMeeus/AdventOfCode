package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type lines []string

func main() {
	fmt.Println(solve1())
	fmt.Println(solve2())
}

func solve1() int {
	lines := getData()

	max := 0
	currentMax := 0
	for _, line := range lines {
		if line == "" {
			if currentMax > max {
				max = currentMax
			}
			currentMax = 0
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		currentMax += num

	}
	return max
}

func solve2() int {
	lines := getData()

	sums := []int{}

	currentMax := 0
	for _, line := range lines {
		if line == "" {
			sums = append(sums, currentMax)
			currentMax = 0
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		currentMax += num

	}

	sort.Slice(sums, func(i, j int) bool { return sums[i] > sums[j] })

	return sums[0] + sums[1] + sums[2]
}

func getData() []string {
	f, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(f), "\n")
}
