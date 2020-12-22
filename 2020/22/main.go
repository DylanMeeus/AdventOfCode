package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Config struct {
	deck1, deck2 string
}

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

func (q *Queue) Copy(n int) *Queue {
	ca := make([]int, n)
	for i := 0; i < n; i++ {
		ca[i] = q.a[i]
	}
	return &Queue{a: ca}
}

func (q *Queue) ToString() string {
	s := ""
	for _, x := range q.a {
		s += strconv.Itoa(x)
	}
	return s
}

func main() {
	fmt.Printf("%v\n", solve1())
	fmt.Printf("%v\n", solve2())
}

func solve1() int {
	p1, p2 := getInput()

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

func solve2() int {
	p1, p2 := getInput()
	p1won := play(p1, p2, map[Config]bool{}, map[Config]bool{})

	winner := p1
	if !p1won {
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

func conf(p1, p2 *Queue) Config {
	return Config{
		deck1: p1.ToString(),
		deck2: p2.ToString(),
	}

}

func play(p1, p2 *Queue, cache, memo map[Config]bool) (p1won bool) {
	// if we have seen this config before, player 1 wins instantly
	// now we check the cache to see if we have seen this configuration before..
	startconfig := conf(p1, p2)
	if p1, ok := memo[startconfig]; ok {
		return p1
	}
	for !p1.Empty() && !p2.Empty() {
		// check here if we have seen the config before
		if cache[conf(p1, p2)] {
			return true
		} else {
			cache[conf(p1, p2)] = true
		}

		top1, top2 := p1.Pop(), p2.Pop()
		// determine if we have to play a subgame

		if p1.Len() >= top1 && p2.Len() >= top2 {
			// play sub game!
			// copy the queues and press play, essentially
			if play(p1.Copy(top1), p2.Copy(top2), map[Config]bool{}, memo) {
				// player 1 won the cards
				p1.Push(top1)
				p1.Push(top2)
			} else {
				p2.Push(top2)
				p2.Push(top1)
			}
		} else {
			if top1 > top2 {
				p1.Push(top1)
				p1.Push(top2)
			} else {
				p2.Push(top2)
				p2.Push(top1)
			}
		}
	}

	memo[startconfig] = !p1.Empty()
	// determine the winner
	return memo[startconfig]
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
