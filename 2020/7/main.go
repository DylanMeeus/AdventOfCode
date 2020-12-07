package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Node struct {
	ID    string
	Value int
}

func main() {
	fmt.Printf("%v\n", solve1())
	fmt.Printf("%v\n", solve2())
}

func solve1() int {
	target := "shinygold"
	input := getInput()
	adj := map[string][]Node{}
	for _, line := range input {
		if line == "" {
			continue
		}
		parse(line, adj)
	}

	out := 0

	for k, _ := range adj {
		if k != target {
			if canReach(k, target, adj) {
				out++
			}
		}
	}

	return out
}

func solve2() int {
	target := "shinygold"
	input := getInput()
	adj := map[string][]Node{}
	for _, line := range input {
		if line == "" {
			continue
		}
		parse(line, adj)
	}

	out := traverse(target, adj)

	return out
}

// how many bags can be reduced to this one
func traverse(node string, adj map[string][]Node) int {
	if adj[node] == nil {
		return 0
	}
	cost := 0
	for _, child := range adj[node] {
		childcost := child.Value
		inner := traverse(child.ID, adj)
		if inner != 0 {
			cost += childcost * inner
		}
		cost += childcost
	}
	return cost

}

func canReach(node, target string, adj map[string][]Node) bool {
	if adj[node] == nil {
		return false
	}
	if node == target {
		return true
	}

	for _, child := range adj[node] {
		if canReach(child.ID, target, adj) {
			return true
		}
	}

	return false
}

func parse(s string, adj map[string][]Node) {
	words := strings.Split(s, " ")
	parent := words[0] + words[1]
	remainder := words[2:]
	// see them in groups of two
	child := ""
	count := 0
	children := []Node{}
	for _, rem := range remainder {
		if strings.ContainsAny(rem, "0123456789") {
			i, _ := strconv.Atoi(rem)
			count = i
		}
		if !strings.ContainsAny(rem, "0123456789,") && !strings.Contains(rem, "contain") && !strings.Contains(rem, "bag") &&
			!strings.Contains(rem, "bags") {
			if child == "" {
				child += rem
			} else {
				child += rem
				node := Node{child, count}
				children = append(children, node)
				child = ""
			}

		}
	}
	for _, child := range children {
		adj[parent] = append(adj[parent], child)
	}
}

func getInput() []string {
	in, _ := ioutil.ReadFile("input.txt")
	return strings.Split(string(in), "\n")
}
