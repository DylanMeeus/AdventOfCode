package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instruction struct {
	amount int
	from   int
	to     int
}

func main() {
	fmt.Println(solve1())
	fmt.Println(solve2())
}

func solve1() string {
	stacks, instructions := getData()

	for _, instr := range instructions {
		for i := 0; i < instr.amount; i++ {
			val := stacks[instr.from-1].pop()
			stacks[instr.to-1].push(val)
		}
	}

	out := ""
	for i := 0; i < 9; i++ {
		out += stacks[i].pop()
	}
	return out
}

func solve2() string {
	stacks, instructions := getData()

	for _, instr := range instructions {
		values := []string{}
		for i := 0; i < instr.amount; i++ {
			val := stacks[instr.from-1].pop()
			values = append(values, val)
		}
		for i := len(values) - 1; i >= 0; i-- {
			stacks[instr.to-1].push(values[i])
		}
	}

	out := ""
	for i := 0; i < 9; i++ {
		out += stacks[i].pop()
	}
	return out
}

func getData() ([]*stack, []instruction) {

	stacks := make([]*stack, 9)
	for i := 0; i < len(stacks); i++ {
		stacks[i] = &stack{}
	}

	lines := readFile()
	idxMap := map[int]int{}
	for i, ch := range lines[8] {
		if string(ch) != " " {
			num, err := strconv.Atoi(string(ch))
			if err != nil {
				panic(err)
			}
			idxMap[i] = num
		}
	}

	for i := 7; i >= 0; i-- {
		for i, ch := range lines[i] {
			if idx, ok := idxMap[i]; ok && string(ch) != " " {
				stacks[idx-1].push(string(ch))
			}
		}
	}

	instructions := []instruction{}

	for i := 9; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			continue
		}
		line = strings.Replace(line, "move ", "", -1)
		line = strings.Replace(line, " from ", ",", -1)
		line = strings.Replace(line, " to ", ",", -1)
		parts := strings.Split(line, ",")

		amount, err := strconv.Atoi(parts[0])
		handleError(err)
		from, err := strconv.Atoi(parts[1])
		handleError(err)
		to, err := strconv.Atoi(parts[2])
		handleError(err)

		instructions = append(instructions, instruction{
			amount, from, to,
		})
	}

	return stacks, instructions
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile() []string {
	f, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(f), "\n")
}

type stack struct {
	top *stackNode
}

func (s *stack) push(input string) {
	if s.top == nil {
		s.top = &stackNode{value: input}
	} else {
		node := &stackNode{value: input, next: s.top}
		s.top = node
	}
}

func (s *stack) pop() string {
	if s.empty() {
		panic("can not pop empty stack")
	}

	n := s.top
	s.top = s.top.next
	return n.value
}

func (s *stack) empty() bool {
	return s.top == nil
}

func (s *stack) print() {
	for node := s.top; node != nil; node = node.next {
		fmt.Println(node.value)
	}
}

type stackNode struct {
	value string
	next  *stackNode
}
