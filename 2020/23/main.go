package main

import (
	"fmt"
	"strconv"
	"strings"
)

const MAX_CUP = 9

var (
	testInput = []int{3, 8, 9, 1, 2, 5, 4, 6, 7} // test input
	input     = []int{3, 8, 9, 5, 4, 7, 6, 1, 2} // real input
)

type ListNode struct {
	Next *ListNode
	Val  int
}

func (l *ListNode) Debug() {
	seen := map[int]bool{}
	out := []string{}
	for node := l; node != nil; node = node.Next {
		if seen[node.Val] {
			break
		}
		seen[node.Val] = true
		out = append(out, strconv.Itoa(node.Val))
	}
	fmt.Printf("%v\n", strings.Join(out, "-"))

}

func main() {
	fmt.Printf("%v\n", solve1())
	//fmt.Printf("%v\n", solve2())
}

func atIndex(m map[int]int, idx int) int {

	for k, v := range m {
		if v == idx {
			return k

		}
	}

	fmt.Printf("%v not found in %v\n", idx, m)

	return -1

}

func solve2() int {
	m := getInput2()

	previousCup := -1
	for i := 0; i < 100; i++ {
		fmt.Printf("move :%v\n", i+1)
		var take int
		if previousCup == -1 {
			take = atIndex(m, 0)
			previousCup = take
			//fmt.Printf("\ttook (%v)\n", take)
		} else {
			// at the right of the previous cup
			next := m[previousCup]
			next += 1
			if next >= MAX_CUP {
				next = 0
			}
			take = atIndex(m, next)
			previousCup = take
			//fmt.Printf("\ttook (%v)\n", take)
		}
		nextThree := take3(m, m[take])
		//fmt.Printf("\tpick up: %v\n", nextThree)
		dest := getDestination(m, take, nextThree)
		//fmt.Printf("\tdestination: %v\n", atIndex(m, dest))
		m = move(m, nextThree, dest)

	}

	// find the labels on the cups after cup 1

	fmt.Printf("%v\n", m)
	// let's play the game

	return 0
}
func solve1() int {
	m := getInput()

	fmt.Printf("%v\n", m)
	previousCup := -1
	for i := 0; i < 100; i++ {
		//fmt.Printf("move :%v\n", i+1)
		var take int
		if previousCup == -1 {
			take = atIndex(m, 0)
			previousCup = take
			//fmt.Printf("\ttook (%v)\n", take)
		} else {
			// at the right of the previous cup
			next := m[previousCup]
			next += 1
			if next >= MAX_CUP {
				next = 0
			}
			take = atIndex(m, next)
			previousCup = take
			//fmt.Printf("\ttook (%v)\n", take)
		}
		nextThree := take3(m, m[take])
		//fmt.Printf("\tpick up: %v\n", nextThree)
		dest := getDestination(m, take, nextThree)
		//fmt.Printf("\tdestination: %v\n", atIndex(m, dest))
		m = move(m, nextThree, dest)

	}

	// find the labels on the cups after cup 1

	fmt.Printf("%v\n", m[1])

	fmt.Printf("%v\n", m)
	// let's play the game

	for i := 0; i < 8; i++ {
		fmt.Printf("%v - ", atIndex(m, i))
	}
	fmt.Println()

	return 0
}

func move(m map[int]int, nextThree []int, destinationIdx int) map[int]int {

	// remove "nextThree" from the map and re-index?
	// get the current ordering minus the new ones?

	idxs := []int{}

	for i := 0; i < len(m); i++ {
		if !contains(nextThree, atIndex(m, i)) {
			idxs = append(idxs, atIndex(m, i))
		}
	}

	// now insert after destination... ?

	mod := []int{}

	for i := range idxs {
		if idxs[i] == atIndex(m, destinationIdx) {
			mod = append(mod, idxs[i])
			mod = append(mod, nextThree...)
			continue
		}
		mod = append(mod, idxs[i])
	}

	//fmt.Printf("\tresult: %v\n", mod)
	out := map[int]int{}
	for i, v := range mod {
		out[v] = i
	}

	return out
}

func contains(is []int, needle int) bool {
	for _, i := range is {
		if i == needle {
			return true
		}
	}
	return false
}

func getDestination(m map[int]int, took int, pickedUp []int) int {
	i := took - 1
	for {
		if contains(pickedUp, i) || i <= 0 {
			i--

			if i <= 0 {
				i = MAX_CUP
			}
			continue
		}

		// now it doesn't contain it
		return m[i]
	}

}

func findNodeByID(root *ListNode, id int) *ListNode {
	for node := root; node != nil; node = node.Next {
		if node.Val == id {
			return node
		}
	}
	return nil
}

func cupID(root *ListNode, idx int) int {
	i := 0

	node := root
	for i < idx {
		node = node.Next
		i++
	}

	return node.Val
}

// something like this?
func take3(m map[int]int, idx int) []int {
	a, b, c := (idx+1)%len(m), (idx+2)%len(m), (idx+3)%len(m)
	//fmt.Printf("a %v b %v c %v\n", a, b, c)
	return []int{atIndex(m, a), atIndex(m, b), atIndex(m, c)}
}

func copy(input *ListNode) *ListNode {
	new := &ListNode{Val: input.Val}

	root := new
	for node := input.Next; node != nil; node = node.Next {
		new.Next = &ListNode{Val: node.Val}
	}
	return root
}

func getInput() map[int]int {
	m := map[int]int{}

	for i, v := range input {
		m[v] = i
	}

	return m
}

func getInput2() map[int]int {
	m := map[int]int{}

	lastIdx := 0
	for i, v := range input {
		m[v] = i
		lastIdx = i
	}

	lastIdx++
	for i := 10; i < int(10e6); i++ {
		m[i] = lastIdx
		lastIdx++
	}

	return m
}
