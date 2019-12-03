package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type point struct {
	a, b  int
	steps int
}

type wire []string

func main() {
	fmt.Printf("%v\n", solve2())
}

func solve2() int {
	a, b := readData()
	as, bs := a.unroll(), b.unroll()
	// find collisions
	var min *int
start:
	for _, x := range as {
		for _, y := range bs {
			if x.a == y.a && x.b == y.b {
				fmt.Printf("another cross.. %v\n", x.steps+y.steps)
				st := x.steps + y.steps
				if min == nil || st < *min {
					min = &st
				}
				continue start
			}
		}
	}
	return *min
}
func solve1() int {
	a, b := readData()
	as, bs := a.unroll(), b.unroll()
	// find collisions
	crosses := []point{}
start:
	for _, x := range as {
		for _, y := range bs {
			if x.a == y.a && x.b == y.b {
				fmt.Printf("another cross.. %v\n", x.steps+y.steps)
				crosses = append(crosses, x)
				continue start
			}
		}
	}
	fmt.Printf("done: %v\n", crosses)
	// calculate manhattan distance
	var max *int
	for _, c := range crosses {
		d := int(math.Sqrt(math.Pow(float64(c.a), 2))) + int(math.Sqrt(math.Pow(float64(c.b), 2)))
		if max == nil || d < *max {
			max = &d
		}
	}
	return *max
}

// return all points on the wire..
func (w wire) unroll() []point {
	out := map[point]struct{}{}
	var x, y int
	for _, p := range w {
		char := string(p[0])
		pos, err := strconv.Atoi(p[1:])
		if err != nil {
			panic(err)
		}
		for i := 0; i < pos; i++ {
			switch char {
			case "R":
				x++
			case "L":
				x--
			case "U":
				y--
			case "D":
				y++
			}
			po := point{x, y, len(out) + 1}
			if _, ok := out[po]; !ok {
				out[po] = struct{}{}
			}
		}

	}
	m := make([]point, len(out))
	var i int
	for k, _ := range out {
		m[i] = k
		i++
	}
	return m
}

func readData() (wire, wire) {
	bs, _ := ioutil.ReadFile("input.txt")
	ss := string(bs)
	parts := strings.Split(ss, "\n")
	a := strings.Split(parts[0], ",")
	b := strings.Split(parts[1], ",")
	return wire(a), wire(b)
}
