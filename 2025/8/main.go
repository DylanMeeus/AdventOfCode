package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Box struct {
	x, y, z int
}

type Node struct {
	start, end *Node
	value      Box
}

func distance(b1, b2 Box) int {
	x := float64((b2.x - b1.x) * (b2.x - b1.x))
	y := float64((b2.y - b1.y) * (b2.y - b1.y))
	z := float64((b2.z - b1.z) * (b2.z - b1.z))
	return int(math.Sqrt(x + y + z))
}

func main() {
	bs := parse(readInput())
	fmt.Println(solve1(bs))
}

func findClosest(b Box, bs []Box) Box {
	cd := int(10e6)
	var cb Box
	for _, o := range bs {
		if b == o {
			continue
		}
		d := distance(b, o)
		if d < cd {
			cd = d
			cb = o
		}
	}
	return cb
}

type Pair struct {
	b1, b2 Box
	dist   int
}

func (p Pair) equal(other Pair) bool {
	if p.dist != other.dist {
		return false
	}
	if p.b1 == other.b1 && p.b2 == other.b2 {
		return true
	}
	if p.b1 == other.b2 && p.b2 == other.b1 {
		return true
	}
	return false
}

func solve1(bs []Box) int {
	// create pairs of all boxes, with their distances, and then sort those by length..
	pairs := []Pair{}
	contains := func(needle Pair, ps []Pair) bool {
		for _, p := range ps {
			if p.equal(needle) {
				return true
			}
		}
		return false
	}
	for _, box := range bs {
		for _, box2 := range bs {
			if box == box2 {
				continue
			}
			pair := Pair{box, box2, distance(box, box2)}
			if !contains(pair, pairs) {
				pairs = append(pairs, pair)
			}
		}
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i].dist <= pairs[j].dist })

	boxesContain := func(b Box, bs []Box) bool {
		for _, x := range bs {
			if x == b {
				return true
			}
		}
		return false
	}

	chains := [][]Box{}
	taken := 0

outer:
	for taken != 10 {
		top := pairs[0]

		//fmt.Printf("(%v - %v) %v\n", top.b1, top.b2, top.dist)
		// connect them, bit figure out if either b1 or b2 already appears in a chain..

		added := false

		for i, chain := range chains {
			if boxesContain(top.b1, chain) && boxesContain(top.b2, chain) {
				pairs = pairs[1:]
				continue outer
			}
			if boxesContain(top.b1, chain) && !boxesContain(top.b2, chain) {
				taken++
				chains[i] = append(chain, top.b2)
				added = true
			}
			if boxesContain(top.b2, chain) && !boxesContain(top.b1, chain) {
				taken++
				chains[i] = append(chain, top.b1)
				added = true
			}
		}

		if !added {
			chains = append(chains, []Box{top.b1, top.b2})
			taken++
		}
		pairs = pairs[1:]
	}

	for _, chain := range chains {
		fmt.Printf("%v\n", len(chain))
	}

	return 0
}

func parse(lines []string) []Box {
	out := []Box{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		is := []int{}
		for _, part := range parts {
			i, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			is = append(is, i)
		}
		out = append(out, Box{is[0], is[1], is[2]})
	}
	return out
}

func readInput() []string {
	b, err := ioutil.ReadFile("test_input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}
