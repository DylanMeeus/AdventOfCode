package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "sort"
)

var previous []*node
var elves [5]*elf

type elf struct {
    task *node
    start int
}

type tree struct {
    nodes []*node
}

type node struct {
    id string
    next []*node
    pre []*node 
}

func main() {
    tree := prereqs()
    fmt.Println(findPath(tree))
    fmt.Println(solve2(tree))
}


func waittime(n *node) int {
    c := n.id[0]
    return int(c) - 64
}

func solve2(t *tree) int {
    basetime := 60 
    for i := range elves {
        elves[i] = &elf{task: nil, start:0}
    }
    nodes := t.nodes
    original := len(nodes)
    previous = make([]*node, 0)
    var second int
    for len(previous) != original {
        // perform cleanup
        for _,e := range elves {
            if e.task == nil {
                continue
            }
            if e.start + basetime + waittime(e.task) == second {
                previous = append(previous, e.task)
                e.task = nil 
                e.start = 0
            }
        }

        // check if an elf is free
        // assign work
        // check if work is over
        assignedDuringLoop := []*node{}
        // resort them
        sort.Slice(nodes, func(i, j int) bool {
            return nodes[i].id < nodes[j].id
        })
        for _,n := range nodes {
            if valid(n.pre, previous) && !contains(n, assignedDuringLoop) {
                for _, e := range elves {
                    if e.task == nil  && !contains(n, assignedDuringLoop){
                        assignedDuringLoop = append(assignedDuringLoop,n)
                        e.task = n
                        e.start = second
                    }
                }
            }
        }
        for _,n := range assignedDuringLoop {
            nodes = filter(nodes, n)
        }
        second++
    }

    return second-1 // the last worked second the work was actually done :)
}

func findPath(tree *tree) string {

    // create a fake root
    previous = make([]*node,0)
    findNext(tree.nodes)
    var out string
    for _,p := range previous {
        out += p.id
    }
    return out
}

func findNext(sorted []*node) {
    // sort nodes?
    sort.Slice(sorted, func(i, j int) bool {
        return sorted[i].id < sorted[j].id
    })

    if len(sorted) == 0 {
        return
    }
    for _,n := range sorted {
        if valid(n.pre, previous) {
            previous = append(previous, n)
            findNext(filter(sorted,n))
            return 
        }
    }
}

func filter(nodes []*node, target *node) []*node {
    f := make([]*node,0)
    for _,n := range nodes {
        if n.id != target.id {
            f = append(f, n)
        }
    }
    return f
}

func valid(pre []*node, seen []*node) bool {
    for _,n := range pre {
        if !contains(n, seen) {
            return false
        }
    }
    return true 
}

func contains(n *node, nodes []*node) bool {
    for _,o := range nodes {
        if n.id == o.id {
            return true
        }
    }
    return false
}


func prereqs() *tree {
    bytes,_ := ioutil.ReadFile("input.txt")
    nodemap := make(map[string]*node,0)
    for _,s := range strings.Split(string(bytes), "\n") {
        if s != "" {
            parts := strings.Split(s, " ")  
            id := parts[1]
            nextid := parts[7]
            currentNode := nodemap[id]
            nextNode := nodemap[nextid]
            if nextNode == nil {
                nextNode = &node{nextid, []*node{}, []*node{}}
                nodemap[nextid] = nextNode
            }
            if currentNode == nil {
                currentNode = &node{id, []*node{}, []*node{}}
                nodemap[id] = currentNode
            }
            nextNode.pre = append(nextNode.pre, currentNode)
            currentNode.next = append(currentNode.next, nextNode)
        }
    }
    // sort the nodemap and return them

    keys := make([]string, 0)
    for k,_ := range nodemap {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    sortedNodes := make([]*node, 0)
    for _,key := range keys {
        sortedNodes = append(sortedNodes, nodemap[key])
    }
    return &tree{sortedNodes}
}
