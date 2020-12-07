package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Printf("%v\n", solve1())
}

func solve1() int {
	target := "shinygold"
	input := getInput()
	adj := map[string][]string{}
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

// how many bags can be reduced to this one

func canReach(node, target string, adj map[string][]string) bool {
	if adj[node] == nil {
		return false
	}
	if node == target {
		return true
	}

	for _, child := range adj[node] {
		if canReach(child, target, adj) {
			return true
		}
	}

	return false
}

func parse(s string, adj map[string][]string) {
	words := strings.Split(s, " ")
	parent := words[0] + words[1]
	remainder := words[2:]
	// see them in groups of two
	child := ""
	children := []string{}
	for _, rem := range remainder {
		if !strings.ContainsAny(rem, "0123456789,") && !strings.Contains(rem, "contain") && !strings.Contains(rem, "bag") &&
			!strings.Contains(rem, "bags") {
			if child == "" {
				child += rem
			} else {
				child += rem
				children = append(children, child)
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
