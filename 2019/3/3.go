package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type point struct {
	a, b int
}

type wire []string

func main() {
	fmt.Printf("%v\n", solve1())
}

func solve1() int {
	a, b := readData()
	as, bs := a.unroll(), b.unroll()
	// find collisions
	crosses := []point{}
	for _, x := range as {
		for _, y := range bs {
			if x == y {
				crosses = append(crosses, x)
			}
		}
	}
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
func (w wire) unroll() (out []point) {
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
			out = append(out, point{x, y})
		}

	}
	return
}

func readData() (wire, wire) {
	bs, _ := ioutil.ReadFile("input.txt")
	ss := string(bs)
	parts := strings.Split(ss, "\n")
	a := strings.Split(parts[0], ",")
	b := strings.Split(parts[1], ",")
	return wire(a), wire(b)
}
