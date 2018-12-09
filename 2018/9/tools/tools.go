package tools 

import (
    "strconv"
    "fmt"
)

// make a doubly-linked list
// figure out the insertion algorithm


// doubly linked list
type List struct {
    start *Node
    Size int
}

type Node struct {
    val int
    prev *Node
    next *Node
}

func (l *List) AddAt(value, position int) {
    if position == 0 {
        l.AddHead(value)
        return
    } 
    if position == l.Size {
        l.Add(value)
        return
    }
    toReplace := l.start
    for i := 0; i < position; i++ {
        toReplace = toReplace.next
    }
    fmt.Printf("replacing: %v\n", toReplace.val)
    // now we have the node where we want to insert
    n := &Node{value, nil, nil}
    before := toReplace.prev
    before.next = n
    n.prev = before
    n.next = toReplace
    toReplace.prev = n
    l.Size++
}

func (l *List) AddHead(value int) {
    // push everything and add the head
    n := &Node{value, l.start.prev, l.start}
    l.start = n
    l.Size++
}

func (l *List) Add(value int) {
    if l.Size == 0 {
        l.start = &Node{value, nil, nil}
        l.start.next = l.start
        l.start.prev = l.start
        l.Size++
    } else {
        n := l.start
        for i := 1; i < l.Size; i++{
            n = n.next
        }
        n.next = &Node{value, n, l.start}
        l.Size++
    }
}

func (l *List) String() string {
    var output string
    n := l.start
    for i := 0; i < l.Size; i++ {
        output += " " + strconv.Itoa(n.val)
        n = n.next
    }
    return output
}
