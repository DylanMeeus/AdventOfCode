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
	fmt.Println("vim-go")
	d := readData()
	fmt.Printf("%v\n", d)
	solve(d)
}

func solve(t tree) {
	fmt.Printf("%v\n", t.root.children)
	printy(t.root)
}

func printy(n *node) {
	fmt.Printf("%v %v\n", n, n.children)
	for _, r := range n.children {
		printy(r)
	}
}

func readData() tree {
	data, err := ioutil.ReadFile("test.txt")
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
	return t
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
