package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type irange struct {
	start int
	end   int
}

func (i irange) containsAll(other irange) bool {
	return other.start >= i.start && other.end <= i.end
}

func (i irange) containsAny(other irange) bool {

	for j := other.start; j <= other.end; j++ {
		if j >= i.start && j <= i.end {
			return true
		}
	}
	return false

}

type pair struct {
	a irange
	b irange
}

func main() {
	fmt.Println(solve1())
	fmt.Println(solve2())
}

func solve1() int {

	pairs := getData()

	sum := 0
	for _, p := range pairs {
		if p.a.containsAll(p.b) || p.b.containsAll(p.a) {
			sum++
		}
	}

	return sum
}

func solve2() int {

	pairs := getData()

	sum := 0
	for _, p := range pairs {
		if p.a.containsAny(p.b) || p.b.containsAny(p.a) {
			sum++
		}
	}

	return sum
}

func getData() []pair {
	f, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(f), "\n")
	out := []pair{}

	getRange := func(s string) irange {
		parts := strings.Split(s, "-")
		a, err := strconv.Atoi(parts[0])
		handleError(err)
		b, err := strconv.Atoi(parts[1])
		handleError(err)
		return irange{
			start: a, end: b,
		}
	}

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		out = append(out, pair{getRange(parts[0]), getRange(parts[1])})

	}
	return out
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
