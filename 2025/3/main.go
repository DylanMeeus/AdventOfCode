package main

import (
	"cmp"
	"fmt"
	"io/ioutil"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	/*
		n := lineToTree("1234", 0)
		printTree(n, "")
		fmt.Printf("%v\n", treeMap)
		fmt.Printf("%v\n", treeMap[0])
		fmt.Printf("%v\n", treeMap[0].Children[0])
		fmt.Printf("%v\n", treeMap[1])
		sorted := sortNodes(n.Children, 1)
		for _, s := range sorted {
			fmt.Println(s.Value)
		}
		return
	*/
	lines := readInput()
	fmt.Println(solve1(lines))
	fmt.Println("solving 2")
	m := solve2(lines)
	fmt.Println(len(strconv.Itoa(m)))
	fmt.Println(m)
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

var (
	speedyMap = map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
)

// todo: figure out the early exit path
// probably something to do with characters at each.
func findMax12(line string) int {
	if line == "" {
		return 0
	}

	first12 := line[0:12]
	num, err := strconv.Atoi(first12)
	if err != nil {
		panic(err)
	}

	max := num
	treeMap = map[int]*Node{}
	root := lineToTree(line, 0)
	// now find the path with max value, using max 12 nodes..
	shouldExit := func(a, b, remainder int) bool {
		as := strconv.Itoa(a)[0 : 12-remainder]
		bs := strconv.Itoa(b)[0 : 12-remainder]

		for i := 0; i < len(as); i++ {
			if speedyMap[as[i]] < speedyMap[bs[i]] {
				return true
			} else if speedyMap[as[i]] == speedyMap[bs[i]] {
				continue
			} else {
				return false
			}

		}

		return false
	}
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
		if shouldExit(currentValue+nodeValue, max, remainder) {
			//fmt.Printf("%v is smaller than %v with only %v remaining", currentValue, max, remainder)
			return
		}
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

var (
	// map an index to a node??
	treeMap = map[int]*Node{}
)

func lineToTree(line string, idx int) *Node {
	if line == "" {
		return nil
	}
	val, err := strconv.Atoi(string(line[0]))
	if err != nil {
		panic(err)
	}
	var node *Node
	if n, ok := treeMap[idx]; ok {
		return n
	}
	node = &Node{Value: val, DepthRemaining: len(line) - 1}
	if node.Children == nil || len(node.Children) == 0 {
		for i := 1; i < len(line); i++ {
			node.Children = append(node.Children, lineToTree(line[i:], idx+i))
		}
	}
	treeMap[idx] = node
	return node
}

func sortNodes(children []*Node, depth int) []*Node {
	// sortNodes from best child to worst child (based on value and depth remaining)

	out := []*Node{}

	for _, child := range children {
		if child.DepthRemaining >= depth {
			out = append(out, child)
		}
	}

	slices.SortFunc(out, func(i, j *Node) int {
		return cmp.Compare(i.DepthRemaining, j.DepthRemaining)
	})

	return out
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

	return 0
}
func readInput() []string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("done")
	return strings.Split(string(b), "\n")
}

func printTree(n *Node, start string) {

	fmt.Printf("%v node value: %v, D rem: %v\n", start, n.Value, n.DepthRemaining)
	for _, c := range n.Children {
		printTree(c, start+"-")
	}

}
