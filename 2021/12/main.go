package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Printf("%v\n", solve())
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
