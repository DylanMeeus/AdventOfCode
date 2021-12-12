package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Printf("%v\n", solve())
	fmt.Printf("%v\n", solve2())
}

// build the adjacency map
func getData() map[string][]string {

	in, _ := ioutil.ReadFile("./input.txt")

	m := map[string][]string{}
	for _, line := range strings.Split(string(in), "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "-")
		from, to := parts[0], parts[1]

		m[from] = append(m[from], to)
		m[to] = append(m[to], from)
	}

	return m
}

func solve() int {
	data := getData()
	state := [][]string{}
	trace("start", data, []string{}, &state)
	return len(state)
}

func solve2() int {
	data := getData()
	state := [][]string{}
	trace2("start", data, []string{}, &state, map[string]int{})

	for _, path := range state {
		fmt.Printf("%v\n", path)
	}

	return len(state)
}

func trace2(node string, adj map[string][]string, currentPath []string,
	state *[][]string,
	visitedCount map[string]int) {
	currentPath = append(currentPath, node)
	visitedCount[node]++

	if node == "end" {
		*state = append(*state, currentPath)
		return
	}

	for _, connection := range adj[node] {
		if isValidMove(visitedCount, connection) {
			copyPath := make([]string, len(currentPath))
			copy(copyPath, currentPath)
			trace2(connection, adj, copyPath, state, copyMap(visitedCount))
		}
	}
}

func copyMap(origin map[string]int) map[string]int {
	clone := map[string]int{}

	for k, v := range origin {
		clone[k] = v
	}
	return clone
}

func isValidMove(m map[string]int, target string) bool {
	if target == "start" {
		return false
	}

	if strings.ToUpper(target) == target {
		return true
	}

	count := m[target]
	if count == 0 {
		return true
	}

	if count >= 2 {
		return false
	}

	// else it's only valid if this has less counts than other small caves
	// we know there's only one count now

	if count != 1 {
		// sanity check "assert style"
		panic(count)
	}

	for k, v := range m {
		if strings.ToUpper(k) == k {
			continue
		}
		if v > 1 && k != target {
			return false
		}
	}
	return true
}

func trace(node string, adj map[string][]string, currentPath []string, state *[][]string) {
	currentPath = append(currentPath, node)
	if node == "end" {
		*state = append(*state, currentPath)
	}

	for _, connection := range adj[node] {
		if !contains(currentPath, connection) || strings.ToUpper(connection) == connection {
			copyPath := make([]string, len(currentPath))
			copy(copyPath, currentPath)
			trace(connection, adj, copyPath, state)
		}
	}
}

func contains(hay []string, needle string) bool {
	for _, s := range hay {
		if s == needle {
			return true
		}
	}
	return false
}
