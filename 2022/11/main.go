package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Monkey struct {
	ID        int
	Items     *Queue
	Operation func(i int) int
	TestDiv   int
	Ytarget   int
	Ntarget   int
}

func main() {
	fmt.Println(solve1())
}

func solve1() int {
	getData()
	return 0
}

func getData() {
	f, err := ioutil.ReadFile("./input.txt")
	handleError(err)
	lines := strings.Split(string(f), "\n")

	for i := 0; i < len(lines)-5; i += 6 {
		parseMonkey(lines[i : i+6])
	}
}

func parseMonkey(lines []string) Monkey {
	fmt.Println(lines)
	return Monkey{}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

type Queue struct {
	inner []int
}

func (q *Queue) Empty() bool {
	return len(q.inner) == 0
}

func (q *Queue) Push(i int) {
	q.inner = append(q.inner, i)
}

func (q *Queue) Pop() (int, bool) {
	if q.Empty() {
		return 0, false
	}
	first := q.inner[0]
	q.inner = q.inner[1:]
	return first, true
}
