package main

import (
    "fmt"
    "strconv"
)

// make a doubly-linked list
// figure out the insertion algorithm


// doubly linked list
type list struct {
    start *node
    size int
}

type node struct {
    val int
    prev *node
    next *node
}

func main(){
    l := &list{}
    l.add(0)
    l.add(1)
    l.add(2)
    l.addAt(3,3)
    fmt.Printf("%v\n",l)
}

func (l *list) addAt(value, position int) {
    if position == 0 {
        l.addHead(value)
        return
    } 
    if position == l.size {
        l.add(value)
        return
    }
    toReplace := l.start
    for i := 0; i < position; i++ {
        toReplace = toReplace.next
    }
    // now we have the node where we want to insert
    n := &node{value, toReplace.prev, toReplace}
    toReplace.prev.next = n
    l.size++
}

func (l *list) addHead(value int) {
    // push everything and add the head
    n := &node{value, l.start.prev, l.start}
    l.start = n
    l.size++
}

func (l *list) add(value int) {
    if l.size == 0 {
        l.start = &node{value, nil, nil}
        l.start.next = l.start
        l.start.prev = l.start
        l.size++
    } else {
        n := l.start
        for i := 1; i < l.size; i++{
            n = n.next
        }
        n.next = &node{value, n, l.start}
        l.size++
    }
}

func (l *list) String() string {
    var output string
    n := l.start
    for i := 0; i < l.size; i++ {
        output += " " + strconv.Itoa(n.val)
        n = n.next
    }
    return output
}
