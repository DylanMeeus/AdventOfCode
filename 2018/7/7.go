package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "sort"
)

type tree struct {
    nodes []*node
}

type node struct {
    id string
    next []*node
    pre []*node //prerequisites
}

func main() {
    tree := prereqs()
    fmt.Println(findPath(tree))
}

var previous []*node
func findPath(tree *tree) string {

    // create a fake root
    previous = make([]*node,0)
    findNext(tree.nodes)
    for _,p := range previous {
        fmt.Printf("%v", p.id)
    }
    return ""
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
    //findNext(sorted)
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
    /*
    for _,v := range nodemap {
        fmt.Println("prerequisites for: " + v.id)
        for _,n := range v.pre {
            if n == nil {
                continue
            }
            fmt.Printf("%v\n", n.id)
        }
    }
    */
    // find the root (element without parent children)
    //return &tree{findRoot(nodemap)}
}

func findRoot(nodes map[string]*node) []*node {
    seen := make(map[string]bool, len(nodes))
    for _,v := range nodes {
        for _,c := range v.next {
            seen[c.id] = true
        }
    }
    possibleRoots := []string{}
    for k,_ := range nodes {
        if seen[k] == false {
            possibleRoots = append(possibleRoots,k)
        }
    }
    sort.Strings(possibleRoots)
    rootnodes := []*node{}
    for _,s := range possibleRoots {
        rootnodes = append(rootnodes, nodes[s])
    }
    return rootnodes 
}
