package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Bounds struct {
	start, end int
}

func contains(b Bounds, i int) bool {
	return i >= b.start && i <= b.end
}

func main() {
	ri := readInput()
	m, is := parse(ri)
	fmt.Println(solve1(m, is))
	fmt.Println(solve2(m))
}

func solve1(bs []Bounds, is []int) int {
	c := 0

	sort.Slice(bs, func(i, j int) bool {
		return bs[i].start < bs[j].start
	})

outer:
	for _, i := range is {
		for _, b := range bs {
			if contains(b, i) {
				//fmt.Printf("%v contains %v\n", b, i)
				c++
				continue outer
			}
			if i < b.start {
				break
			}

		}
	}

	return c
}

func solve2(bs []Bounds) int {
	sort.Slice(bs, func(i, j int) bool {
		return bs[i].start < bs[j].start
	})

	acc := 0

	current_start := bs[0].start
	current_end := bs[0].end
	for _, b := range bs[1:] {

		// figure out if we completed a chain..
		if b.start <= current_end {
			// extend the chain if the end is further out..
			if b.end > current_end {
				current_end = b.end
			}
		} else {
			// we have a complete chain.. figure out it's length and add to acc
			//fmt.Printf("start: %v, end %v\n", current_start, current_end)
			acc += ((current_end - current_start) + 1)
			current_start, current_end = b.start, b.end
		}

	}

	// add the final chain and return acc
	return acc + ((current_end - current_start) + 1)
}

func parse(lines []string) ([]Bounds, []int) {
	bs := []Bounds{}
	out := []int{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, "-") {
			// parse it
			parts := strings.Split(line, "-")
			ls, rs := parts[0], parts[1]
			li, err := strconv.Atoi(ls)
			if err != nil {
				panic(err)
			}
			ri, err := strconv.Atoi(rs)
			if err != nil {
				panic(err)
			}
			bs = append(bs, Bounds{li, ri})
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			out = append(out, num)
		}
	}
	return bs, out
}

func readInput() []string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}
