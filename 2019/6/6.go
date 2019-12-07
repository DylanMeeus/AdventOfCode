package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type node struct {
	name     string
	children []*node
}

func (n node) String() string {
	return n.name
}

type tree struct {
	root *node
}

func main() {
	m, t := readData()
	total := solve(m, t)
	fmt.Printf("%v\n", total)
	steps := solve2(t)
	fmt.Printf("%v\n", steps)
}

func solve2(t tree) int {
	// find first common ancestor
	younode := findNode("YOU", t.root)
	sannode := findNode("SAN", t.root)

	// trace path to root
	youPath := []string{}
	for p := findParent(t.root, younode); p.name != t.root.name; {
		youPath = append(youPath, p.name)
		p = findParent(t.root, p)
	}
	sanPath := []string{}
	for p := findParent(t.root, sannode); p.name != t.root.name; {
		sanPath = append(sanPath, p.name)
		p = findParent(t.root, p)
	}

	// find common ancestor
	for i, y := range youPath {
		for j, s := range sanPath {
			if y == s {
				return i + j
			}
		}
	}
	return 0
}

func solve(m map[string][]string, t tree) int {
	var sum int
	collapsed := map[string]struct{}{}
	for k, v := range m {
		collapsed[k] = struct{}{}
		for _, s := range v {
			collapsed[s] = struct{}{}
		}
	}
	for k, _ := range collapsed {
		orbits := 0 // one is direct
		n := findNode(k, t.root)
		if n == t.root {
			continue
		}
		for p := findParent(t.root, n); p.name != t.root.name; {
			orbits++
			p = findParent(t.root, p)
		}
		sum += orbits + 1
	}
	return sum // remove COM
}

func toRoot(n *node, steps int) int {
	if n.name == "COM" {
		return steps
	}
	return 0

}

func findParent(current, target *node) *node {
	for _, n := range current.children {
		if n.name == target.name {
			return current
		}
		lookup := findParent(n, target)
		if lookup != nil {
			return lookup
		}
	}
	return nil
}

func printy(n *node) {
	fmt.Printf("%v %v\n", n, n.children)
	for _, r := range n.children {
		printy(r)
	}
}

func readData() (map[string][]string, tree) {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	parts := strings.Split(string(data), "\n")
	t := tree{
		root: &node{
			"COM",
			[]*node{},
		},
	}
	m := map[string][]string{}
	for _, p := range parts {
		if p == "" {
			continue
		}
		pp := strings.Split(p, ")")
		first, second := pp[0], pp[1]
		m[first] = append(m[first], second)
	}
	mapToTree(&t, "COM", m)
	return m, t
}

func mapToTree(t *tree, cur string, m map[string][]string) {
	n := findNode(cur, t.root)
	if n == nil {
		n = &node{name: cur}
	}
	for _, s := range m[cur] {
		newnode := &node{
			name: s, children: []*node{},
		}
		if n.children == nil {
			n.children = []*node{newnode}
		} else {
			n.children = append(n.children, newnode)
		}
		mapToTree(t, s, m)
	}
}

func findNode(value string, current *node) *node {
	if current.name == value {
		return current
	}
	for _, c := range current.children {
		n := findNode(value, c)
		if n != nil {
			return n
		}
	}
	return nil
}
