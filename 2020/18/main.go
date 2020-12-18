package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type StackNode struct {
	Next *StackNode
	Val  string
}

type Stack struct {
	Top *StackNode
}

func (s *Stack) Push(in string) {
	sn := &StackNode{Val: in}
	sn.Next = s.Top
	s.Top = sn
}

func (s *Stack) Pop() string {
	if s.Top == nil {
		return ""
	}
	tmp := s.Top.Val
	s.Top = s.Top.Next
	return tmp
}

func (s *Stack) Empty() bool {
	return s.Top == nil
}

func (s *Stack) Debug() {
	for node := s.Top; node != nil; node = node.Next {
		fmt.Printf("%v\n", node.Val)
	}
}

// reduce to a number yo
func (s *Stack) Reduce() int {
	if s.Empty() {
		return 0
	}

	s = reverse(s)
	for !s.Empty() {
		a := s.Pop()
		op := s.Pop()
		b := s.Pop()

		ia, _ := strconv.Atoi(a)
		ib, _ := strconv.Atoi(b)

		if op == "" {
			return ia
		}
		if op == "*" {
			s.Push(strconv.Itoa(ia * ib))
		}
		if op == "+" {
			s.Push(strconv.Itoa(ia + ib))
		}
	}

	return -1
}

func reverse(s *Stack) *Stack {
	news := &Stack{}

	for !s.Empty() {
		news.Push(s.Pop())
	}
	return news

}

type tokens []string

func main() {
	fmt.Printf("solve 1: %v\n", solve1())
}

func solve1() int {
	input := getInput()

	out := 0

	for _, in := range input {
		out += solveEq(in)
	}
	return out
}

func solveEq(eq tokens) int {
	// how to deal with the parens??
	fmt.Printf("solving: %v\n", eq)

	i := 0

	s := Stack{}
	for i < len(eq) {
		if eq[i] == "(" {
			// find closing parens
			closing := findClosingIndex(i, eq)
			solution := solveEq(eq[i+1 : closing])
			isol := strconv.Itoa(solution)
			s.Push(isol)
			i = closing + 1
		} else if eq[i] != ")" {
			s.Push(eq[i])
			i++
		}
	}

	fmt.Println("DEBUG")
	s.Debug()
	fmt.Println("END DEBUG")

	// solve the stack stuff

	return s.Reduce()
}

func findClosingIndex(start int, eq tokens) int {
	parens := Stack{}
	parens.Push(eq[start])
	for i := start + 1; i < len(eq); i++ {
		if eq[i] == "(" {
			parens.Push(eq[i])
		}
		if eq[i] == ")" {
			parens.Pop()
			if parens.Empty() {
				return i
			}
		}
	}
	return -1

}

func getInput() []tokens {
	in, _ := ioutil.ReadFile("input.txt")

	filter := func(in []string) (out []string) {

		for _, i := range in {
			if i != " " {
				out = append(out, i)
			}
		}
		return

	}

	t := []tokens{}
	for _, line := range strings.Split(string(in), "\n") {
		if line == "" {
			continue
		}
		t = append(t, filter(strings.Split(line, "")))
	}

	return t
}
