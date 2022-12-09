package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solve1())
}

func solve1() int {
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
	return 0

}

type node struct {
	size     int
	name     string
	isDir    bool
	parent   *node
	children []*node
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
