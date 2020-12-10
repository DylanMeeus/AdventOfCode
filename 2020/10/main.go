package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%v\n", solve1())
}

func solve1() int {
	is := append([]int{0}, getInput()...)
	sort.Slice(is, func(i, j int) bool { return is[i] < is[j] })

	// find all the differences
	j1, j3 := 0, 1

	fmt.Printf("%v\n", is)
	for i := 0; i < len(is)-1; i++ {
		delta := is[i+1] - is[i]
		if delta == 1 {
			j1++
		}
		if delta == 3 {
			j3++
		}
	}
	return j1 * j3
}

func getInput() []int {
	in, _ := ioutil.ReadFile("input.txt")
	out := []int{}
	for _, line := range strings.Split(string(in), "\n") {
		if line == "" {
			continue
		}
		i, _ := strconv.Atoi(line)
		out = append(out, i)
	}
	return out
}
