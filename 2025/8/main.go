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
	value Box
	next  *Node
}

func (n *Node) Print() {
	s := n
	for s != nil {
		fmt.Printf(" %v -", s.value)
		s = s.next
	}
	fmt.Println()
}

func (n *Node) Len() int {
	i := 0
	s := n
	for s != nil {
		i++
		s = s.next
	}
	return i
}

func addTail(n *Node, b Box) {
	node := n

	for node.next != nil {
		node = node.next
	}
	// now we found the last node
	node.next = &Node{
		value: b,
	}

}

func addTailNode(n, t *Node) {
	node := n

	for node.next != nil {
		node = node.next
	}
	// now we found the last node
	node.next = t

}

func containsBox(b Box, n *Node) bool {
	node := n
	for node != nil {
		if node.value == b {
			return true
		}
		node = node.next
	}
	return false
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
	fmt.Println(solve2(bs))
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
	pairMap := map[Pair]bool{}

	contains := func(needle Pair, ps []Pair) bool {
		for _, p := range ps {
			if p.equal(needle) {
				return true
			}
		}
		return false
	}
	_ = contains
	for _, box := range bs {
		for _, box2 := range bs {
			if box == box2 {
				continue
			}
			var pair Pair
			if box.x < box2.x {
				pair = Pair{box, box2, distance(box, box2)}
			} else {
				pair = Pair{box2, box, distance(box, box2)}
			}
			if _, ok := pairMap[pair]; ok {
				//fmt.Printf("map contains %v\n", pair)
			} else {
				pairs = append(pairs, pair)
			}
			pairMap[pair] = true
		}
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i].dist <= pairs[j].dist })

	// sorted pairs, turn these into nodes??
	chains := []*Node{}

	taken := 0

	for taken < 10 {
		top := pairs[0]

		a, b := top.b1, top.b2

		var firstChain, secondChain *Node

		for _, chain := range chains {
			if containsBox(a, chain) {
				firstChain = chain
			}
			if containsBox(b, chain) {
				secondChain = chain
			}
		}

		// if both are nil, we create a new chain
		if firstChain == nil && secondChain == nil {
			newChain := &Node{
				value: a,
				next: &Node{
					value: b,
					next:  nil,
				},
			}
			chains = append(chains, newChain)
			taken++
		}

		if firstChain != nil && secondChain == nil {
			// we found a chain for A, but not for B
			addTail(firstChain, b)
			taken++
		}

		if firstChain == nil && secondChain != nil {
			// we found a chain for B but not for A
			addTail(secondChain, a)
			taken++
		}

		if firstChain != nil && secondChain != nil {
			// we only need to connect them if they are not the same chain..

			if firstChain.value != secondChain.value {
				addTailNode(firstChain, secondChain)
				newChains := []*Node{}
				for _, c := range chains {
					if c.value != secondChain.value {
						newChains = append(newChains, c)
					}
				}
				chains = newChains
				taken++
			} else {
				// already in same chain, so it's a no-op
				//fmt.Printf("%v in chain %v\n", firstChain, secondChain)
				taken++
			}
		}

		pairs = pairs[1:]
	}

	sort.Slice(chains, func(i, j int) bool { return chains[i].Len() > chains[j].Len() })

	return chains[0].Len() * chains[1].Len() * chains[2].Len()

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
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}

func solve2(bs []Box) int {
	// create pairs of all boxes, with their distances, and then sort those by length..
	pairs := []Pair{}
	pairMap := map[Pair]bool{}

	contains := func(needle Pair, ps []Pair) bool {
		for _, p := range ps {
			if p.equal(needle) {
				return true
			}
		}
		return false
	}
	_ = contains
	for _, box := range bs {
		for _, box2 := range bs {
			if box == box2 {
				continue
			}
			var pair Pair
			if box.x < box2.x {
				pair = Pair{box, box2, distance(box, box2)}
			} else {
				pair = Pair{box2, box, distance(box, box2)}
			}
			if _, ok := pairMap[pair]; ok {
				//fmt.Printf("map contains %v\n", pair)
			} else {
				pairs = append(pairs, pair)
			}
			pairMap[pair] = true
		}
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i].dist <= pairs[j].dist })
	// sorted pairs, turn these into nodes??
	chains := []*Node{}

	for {
		top := pairs[0]

		a, b := top.b1, top.b2

		var firstChain, secondChain *Node

		for _, chain := range chains {
			if containsBox(a, chain) {
				firstChain = chain
			}
			if containsBox(b, chain) {
				secondChain = chain
			}
		}

		// if both are nil, we create a new chain
		if firstChain == nil && secondChain == nil {
			newChain := &Node{
				value: a,
				next: &Node{
					value: b,
					next:  nil,
				},
			}
			chains = append(chains, newChain)
		}

		if firstChain != nil && secondChain == nil {
			// we found a chain for A, but not for B
			addTail(firstChain, b)
			if firstChain.Len() == len(bs) {
				// we have a full match
				return a.x * b.x
			}
		}

		if firstChain == nil && secondChain != nil {
			// we found a chain for B but not for A
			addTail(secondChain, a)
			if secondChain.Len() == len(bs) {
				// we have a full match
				return a.x * b.x
			}
		}

		if firstChain != nil && secondChain != nil {
			// we only need to connect them if they are not the same chain..

			if firstChain.value != secondChain.value {
				// figure out when the last chain is merged
				addTailNode(firstChain, secondChain)
				newChains := []*Node{}
				for _, c := range chains {
					if c.value != secondChain.value {
						newChains = append(newChains, c)
					}
				}
				chains = newChains
				if firstChain.Len() == len(bs) {
					// we have a full match
					return a.x * b.x
				}
			} else {
				// already in same chain, so it's a no-op
				//fmt.Printf("%v in chain %v\n", firstChain, secondChain)
			}
		}

		pairs = pairs[1:]
	}

	sort.Slice(chains, func(i, j int) bool { return chains[i].Len() > chains[j].Len() })

	return chains[0].Len() * chains[1].Len() * chains[2].Len()

}
