package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%v\n", solve1())
	fmt.Printf("part2: %v\n", solve2())
	fmt.Println("done")
}

func solve2() int {
	is := append([]int{0}, getInput()...)
	sort.Slice(is, func(i, j int) bool { return is[i] < is[j] })
	is = append(is, is[len(is)-1]+3)
	possible := 0
	paths := [][]int{}
	var arrange func(int, []int, []int)
	arrange = func(current int, remainder []int, path []int) {
		sort.Slice(remainder, func(i, j int) bool { return remainder[i] < remainder[j] })
		cp := make([]int, len(path))
		copy(cp, path)
		cp = append(cp, current)
		if len(remainder) == 0 {
			possible++
			paths = append(paths, cp)
			return
		}
		if remainder[0] > current+3 {
			return
		}
		for i, rem := range remainder {
			delta := int(math.Abs(float64(rem - current)))
			if delta <= 3 {
				// remove this one from the list
				c := make([]int, len(remainder))
				copy(c, remainder)
				c = c[i+1:]
				arrange(rem, c, cp)
			} else {
				//break
			}
		}

	}
	arrange(is[0], is[1:], []int{})
	return possible
}

func filter(is []int, x int) (out []int) {
	for _, i := range is {
		if i != x {
			out = append(out, i)
		}
	}
	return
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
