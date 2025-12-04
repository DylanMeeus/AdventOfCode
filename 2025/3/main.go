package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines := readInput()
	fmt.Println(solve1(lines))

	f := (findMax12("987654321111111"))
	fmt.Println(f)
	fmt.Println(len(strconv.Itoa(f)))
	f = (findMax12("234234234234278"))
	fmt.Println(f)
	fmt.Println(len(strconv.Itoa(f)))
	n := lineToTree("1234")
	printTree(n, "")
	fmt.Println(solve2(lines))
}

// find the max of 2 batteries combined
func findMax(line string) int {
	max := 0

	for i := 0; i < len(line); i++ {
		for j := i + 1; j < len(line); j++ {
			combined := string(line[i]) + string(line[j])
			num, err := strconv.Atoi(combined)
			if err != nil {
				panic(err)
			}
			if num > max {
				max = num
			}
		}
	}
	return max
}

type Node struct {
	Value          int
	DepthRemaining int
	Children       []*Node
}

func findMax12(line string) int {
	max := 0
	root := lineToTree(line)
	// now find the path with max value, using max 12 nodes..
	var nodeMax func(*Node, int, int)
	nodeMax = func(start *Node, currentValue, remainder int) {
		if start == nil {
			return
		}
		if start.DepthRemaining < remainder {
			return
		}
		if remainder == 0 {
			nodeValue := start.Value * int(math.Pow(10., float64(remainder)))
			currentValue += nodeValue
			if currentValue > max {
				max = currentValue
			}
			return
		}
		nodeValue := start.Value * int(math.Pow(10., float64(remainder)))
		// take each of the children
		for _, child := range start.Children {
			nodeMax(child, currentValue+nodeValue, remainder-1)
			if start.DepthRemaining > 11 {
				nodeMax(child, 0, 11) // also start a search from the child onwwards
			}
		}
	}
	nodeMax(root, 0, 11)
	return max
}

func lineToTree(line string) *Node {
	if line == "" {
		return nil
	}
	val, err := strconv.Atoi(string(line[0]))
	if err != nil {
		panic(err)
	}
	node := &Node{Value: val, DepthRemaining: len(line) - 1}
	for i := 1; i < len(line); i++ {
		node.Children = append(node.Children, lineToTree(line[i:]))
	}
	return node
}

func solve1(lines []string) int {
	t := 0

	for _, line := range lines {
		t += findMax(line)
	}

	return t
}

func solve2(lines []string) int {
	t := 0

	for i, line := range lines {
		t += findMax12(line)
		fmt.Printf("parsed %v of %v lines\n", i, len(lines))
	}

	return t
}
func readInput() []string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}

func printTree(n *Node, start string) {

	fmt.Printf("%v node value: %v, D rem: %v\n", start, n.Value, n.DepthRemaining)
	for _, c := range n.Children {
		printTree(c, start+"-")
	}

}
