package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var (
	openClose = map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
		"<": ">",
	}

	closingSymbols = map[string]bool{
		")": true,
		"}": true,
		"]": true,
		">": true,
	}
)

func main() {
	fmt.Printf("%v\n", solve())
	fmt.Printf("%v\n", solve2())
}

func getData() []string {
	in, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(in), "\n")
	return lines
}

func solve() int {
	data := getData()

	pointMap := map[string]int{
		"":  0,
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	score := 0
	for _, line := range data {
		if line == "" {
			continue
		}
		firstError := firstSyntaxError(line)
		score += pointMap[firstError]

	}
	return score
}

func solve2() int {
	pointMap := map[string]int{
		"":  0,
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	data := getData()

	// remove lines with syntax errors
	incomplete := filter(data, func(s string) bool { return firstSyntaxError(s) == "" })

	scores := []int{}
	for _, line := range incomplete {
		if line == "" {
			continue
		}
		completion := complete(line)
		score := 0

		for _, char := range strings.Split(completion, "") {
			score *= 5
			score += pointMap[char]
		}
		//fmt.Printf("completion: %v for %v\nscore: %v\n", completion, line, score)
		scores = append(scores, score)
	}

	sort.Ints(scores)
	mid := len(scores) / 2

	return scores[mid]
}

// complete returns the string to complete the lines
func complete(line string) string {
	symbols := strings.Split(line, "")
	stack := Stack{}
	for _, symbol := range symbols {
		if !closingSymbols[symbol] {
			stack.Push(symbol)
		} else {
			_, err := stack.Pop()
			if err != nil {
				panic(err)
			}
		}
	}

	// now the stack contains the ones that do not have a closing paren..

	out := ""

	for !stack.Empty() {
		value, err := stack.Pop()
		if err != nil {
			panic(err)
		}
		out += openClose[value]
	}

	return out
}

func filter(hay []string, pred func(string) bool) []string {
	out := []string{}
	for _, s := range hay {
		if pred(s) {
			out = append(out, s)
		}
	}
	return out
}

// firstSyntaxError returns the first wrong closing symbol
func firstSyntaxError(str string) string {
	symbols := strings.Split(str, "")
	stack := Stack{}
	for _, symbol := range symbols {
		if !closingSymbols[symbol] {
			stack.Push(symbol)
		} else {
			opening, err := stack.Pop()
			if err != nil {
				return ""
			}
			if symbol != openClose[opening] {
				return symbol
			}
		}
	}

	return ""
}

type StackNode struct {
	value string
	next  *StackNode
}

type Stack struct {
	top *StackNode
}

func (s *Stack) Push(str string) {
	node := &StackNode{value: str, next: nil}

	if s.top == nil {
		s.top = node
	} else {
		node.next = s.top
		s.top = node
	}
}

func (s *Stack) Pop() (string, error) {
	if s.Empty() {
		return "", fmt.Errorf("popping empty stack")
	}

	returnValue := s.top.value
	s.top = s.top.next
	return returnValue, nil
}

func (s *Stack) Empty() bool {
	return s.top == nil
}
