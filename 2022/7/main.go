package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const (
	TOTAL_MEM  = 70000000
	NEEDED_MEM = 30000000
)

func main() {
	fmt.Println(solve1())
	fmt.Println(solve2())
}

func buildTree() *node {
	data := getData()

	root := &node{name: "/", isDir: true, children: []*node{}}

	currentNode := root
	for _, line := range data {
		if line == "" {
			continue
		}
		fmt.Println(line)
		if string(line[0]) == "$" {
			// process a command
			parts := strings.Split(line, " ")
			if parts[1] == "cd" {
				if parts[2] == "/" {
					currentNode = root
				} else if parts[2] == ".." {
					currentNode = currentNode.parent

				} else {
					currentNode = currentNode.find(parts[2])
				}
			}
		} else {
			// process the output of a command
			parts := strings.Split(line, " ")
			if strings.HasPrefix(line, "dir") {
				currentNode.addDir(parts[1])
			} else {
				currentNode.addFile(parts[0], parts[1])
			}
		}
	}
	return root
}

func solve1() int {
	tree := buildTree()
	result := []int{}
	var inner func(n *node)
	inner = func(n *node) {
		if n.isDir {
			for _, file := range n.files() {
				n.size += file.size
			}
			for _, dir := range n.dirs() {
				inner(dir)
				n.size += dir.size
			}
			if n.size <= 100_000 {
				result = append(result, n.size)
			}
		}
	}
	inner(tree)
	sum := 0
	for _, i := range result {
		sum += i
	}
	return sum
}

func solve2() int {
	tree := buildTree()
	var inner func(n *node)
	dirSizes := []int{}
	inner = func(n *node) {
		if n.isDir {
			for _, file := range n.files() {
				n.size += file.size
			}
			for _, dir := range n.dirs() {
				inner(dir)
				n.size += dir.size
			}
			dirSizes = append(dirSizes, n.size)
		}
	}
	inner(tree)

	freeSpace := TOTAL_MEM - tree.size

	needed := NEEDED_MEM - freeSpace

	sort.Ints(dirSizes)

	for _, ds := range dirSizes {
		if ds >= needed {
			return ds
		}

	}

	return -1
}

type node struct {
	size     int
	name     string
	isDir    bool
	parent   *node
	children []*node
}

func (n *node) files() []*node {
	out := []*node{}
	for _, c := range n.children {
		if !c.isDir {
			out = append(out, c)
		}
	}
	return out
}

func (n *node) dirs() []*node {
	out := []*node{}
	for _, c := range n.children {
		if c.isDir {
			out = append(out, c)
		}
	}
	return out
}

func (n *node) find(name string) *node {
	for _, c := range n.children {
		if c.name == name {
			return c
		}
	}
	return nil
}

func (n *node) addDir(name string) {
	n.children = append(n.children, &node{name: name, isDir: true, children: []*node{}, parent: n})
}

func (n *node) addFile(size, name string) {
	i, err := strconv.Atoi(size)
	if err != nil {
		panic(err)
	}
	n.children = append(n.children, &node{size: i, name: name, isDir: false, children: nil, parent: n})
}

func getData() []string {
	f, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(f), "\n")
}
