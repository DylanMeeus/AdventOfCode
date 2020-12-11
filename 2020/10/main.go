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
}

func solve2() int {
	is := append([]int{0}, getInput()...)
	sort.Slice(is, func(i, j int) bool { return is[i] < is[j] })
	is = append(is, is[len(is)-1]+3)
	var arrange func(int, []int, map[int]int) int
	arrange = func(current int, remainder []int, memo map[int]int) int {
		if val, ok := memo[current]; ok {
			return val
		}
		if len(remainder) == 0 {
			return 1
		}
		if remainder[0] > current+3 {
			return 0
		}
		total := 0
		for i, rem := range remainder {
			delta := int(math.Abs(float64(rem - current)))
			if delta <= 3 {
				// remove this one from the list
				c := make([]int, len(remainder))
				copy(c, remainder)
				c = c[i+1:]
				ways := arrange(rem, c, memo)
				total += ways
				memo[rem] = ways
			} else {
				break
			}
		}
		if _, ok := memo[current]; !ok {
			memo[current] = total
		}
		return memo[current]
	}
	result := arrange(is[0], is[1:], map[int]int{})
	return result
}

func solve1() int {
	is := append([]int{0}, getInput()...)
	sort.Slice(is, func(i, j int) bool { return is[i] < is[j] })

	// find all the differences
	j1, j3 := 0, 1

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
