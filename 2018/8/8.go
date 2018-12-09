package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

type node struct {
    children []*node
    metadata []int
}

func main(){
    fmt.Println("starting")
    fmt.Printf("%v\n", solve(parse()))
}

func solve(n *node) int {
    var sum int
    for _,c := range n.children {
        sum += solve(c)
    }
    for _,m := range n.metadata {
        sum += m 
    }
    return sum
}



func parse() *node {
    bytes,_ := ioutil.ReadFile("input.txt")
    input := strings.Split(string(bytes[:len(bytes)-1]), " ")
    ints := []int{}
    for _,in := range input {
        if in == "" {
            continue
        }
        i, err := strconv.Atoi(in)
        if err != nil {
            panic(err)
        }
        ints = append(ints, i)
    }
    root,_ := create(ints)
    return root
}

func create(input []int) (*node, []int) {
    if len(input) < 2 {
        return nil, input
    }
    childCount := input[0]
    metadataCount := input[1]
    current := &node{children: []*node{}, metadata: []int{}}
    input = input[2:]
    for i := 0; i < childCount; i++ {
        child,out := create(input)
        input = out 
        current.children = append(current.children, child)
    }
    for i := 0; i < metadataCount; i++ {
        current.metadata = append(current.metadata, input[0])
        input = input[1:]
    }
    return current, input
}
