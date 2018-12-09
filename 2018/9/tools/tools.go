package tools 

import (
    "strconv"
)

// make a doubly-linked list
// figure out the insertion algorithm


// doubly linked list
type List struct {
    start *Node
    Size int
}

type Node struct {
    Val int
    prev *Node
    next *Node
}

func (l *List) AddAt(Value, position int) {
    if position == 0 {
        l.AddHead(Value)
        return
    } 
    if position == l.Size {
        l.Add(Value)
        return
    }
    toReplace := l.start
    for i := 0; i < position; i++ {
        toReplace = toReplace.next
    }
    // now we have the node where we want to insert
    n := &Node{Value, nil, nil}
    before := toReplace.prev
    before.next = n
    n.prev = before
    n.next = toReplace
    toReplace.prev = n
    l.Size++
}

func (l *List) AddHead(Value int) {
    // push everything and add the head
    n := &Node{Value, l.start.prev, l.start}
    l.start = n
    l.Size++
}

func (l *List) Add(Value int) {
    if l.Size == 0 {
        l.start = &Node{Value, nil, nil}
        l.start.next = l.start
        l.start.prev = l.start
        l.Size++
    } else {
        n := l.start
        for i := 1; i < l.Size; i++{
            n = n.next
        }
        n.next = &Node{Value, n, l.start}
        l.Size++
    }
}

func (l *List) RemoveAt(index int) {
    n := l.Get(index)
    l.Remove(n)
}

func (l *List) Remove(n *Node) {
    n.prev.next = n.next
    n.next.prev = n.prev
    if n == l.start {
        l.start = n.next
    }
    l.Size--
    n = nil
}

func (l *List) Get(index int) *Node {
    n := l.start
    for i := 0; i < index; i++ {
        n = n.next
    }
    return n
}

func (l *List) FindRelative(rel, index int) *Node {
    n := l.Get(index)
    for i := 0; i < rel; i++ {
        n = n.prev
    }
    return n
}

func (l *List) String() string {
    var output string
    n := l.start
    for i := 0; i < l.Size; i++ {
        output += " " + strconv.Itoa(n.Val)
        n = n.next
    }
    return output
}
