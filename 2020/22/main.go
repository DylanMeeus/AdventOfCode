package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Queue struct {
	a []int
}

func (q *Queue) Empty() bool {
	return len(q.a) == 0
}

func (q *Queue) Push(i int) {
	q.a = append(q.a, i)
}

func (q *Queue) Pop() int {
	if q.Empty() {
		return -1
	}
	tmp := q.a[0]
	if len(q.a) > 1 {
		q.a = q.a[1:]
	} else {
		q.a = []int{}
	}
	return tmp
}

func (q *Queue) Len() int {
	return len(q.a)
}

func main() {
	fmt.Printf("%v\n", solve1())
}

func solve1() int {
	p1, p2 := getInput()
	fmt.Printf("p1: %v\np2: %v\n", p1, p2)

	for !p1.Empty() && !p2.Empty() {
		top1, top2 := p1.Pop(), p2.Pop()
		if top1 > top2 {
			p1.Push(top1)
			p1.Push(top2)
		} else {
			p2.Push(top2)
			p2.Push(top1)
		}
	}

	winner := p1
	if p1.Empty() {
		winner = p2
	}

	N := winner.Len()

	res := 0
	for !winner.Empty() {
		val := winner.Pop()
		res += (val * N)
		N--
	}

	return res
}

// getInput returns the cards for player 1 / player 2
func getInput() (*Queue, *Queue) {
	in, _ := ioutil.ReadFile("input.txt")

	player1 := &Queue{}
	player2 := &Queue{}
	p1 := true
	for _, line := range strings.Split(string(in), "\n") {
		if line == "" {
			continue
		}

		if line == "#" {
			p1 = false
			continue
		}

		i, _ := strconv.Atoi(line)
		if p1 {
			player1.Push(i)
		} else {
			player2.Push(i)
		}
	}

	return player1, player2
}
