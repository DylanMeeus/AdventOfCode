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
			//for i := li; i <= ri; i++ {
			//m[i] = true
			//}
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
