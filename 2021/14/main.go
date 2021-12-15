package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	//fmt.Printf("%v\n", solve())
	fmt.Printf("%v\n", solve2())
}

func getData() (start string, rules map[string]string) {
	in, _ := ioutil.ReadFile("./input.txt")

	lines := strings.Split(string(in), "\n")
	start = lines[0]

	rules = map[string]string{}
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

		line = strings.Replace(line, " ", "", -1)
		parts := strings.Split(line, "->")

		rules[parts[0]] = parts[1]
	}
	return
}

func solve() int {
	input, rules := getData()

	for i := 0; i < 10; i++ {
		input = applyRules(input, rules)
		fmt.Println(i)
	}

	min, max := minAndMax(input)

	return max - min
}

func getPairs(input string) []string {
	pairs := []string{}
	for i := 0; i < len(input)-1; i++ {
		pair := input[i : i+2]
		pairs = append(pairs, pair)
	}
	return pairs
}

func solve2() int {
	input, rules := getData()

	pairCount := map[string]int{}

	pairs := getPairs(input)

	for _, pair := range pairs {
		pairCount[pair]++
	}

	added := map[string]int{}
	// figur eout how to map the pairs to the most common & uncommon characters
	iterativeExpand(40, rules, pairCount, added)

	for _, char := range input {
		added[string(char)]++
	}

	fmt.Printf("%v\n", added)

	/*
		charCount := map[string]int{}

			for k, v := range added {
				charCount[k] = v
				newChar := rules[k]
				charCount[newChar] = v
			}
	*/

	min := int(10e12)
	max := 0

	for _, v := range added {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return max - min
}

func countFromPairCount(m map[string]int, rules map[string]string) map[string]int {
	out := map[string]int{}
	for k, v := range m {
		toAdd := rules[k]
		out[toAdd] += v
	}
	return out
}

func copyMap(m map[string]int) map[string]int {
	out := map[string]int{}
	for k, v := range m {
		out[k] = v
	}
	return out
}

// to avoid memoization stuff..
func iterativeExpand(levels int, rules map[string]string, paircount map[string]int, added map[string]int) {
	m := copyMap(paircount)
	for i := 0; i < levels; i++ {
		for pair, v := range paircount {
			m[pair] -= v
			left := string(pair[0]) + rules[pair]
			right := rules[pair] + string(pair[1])

			added[rules[pair]] += v

			m[left] += v
			m[right] += v
		}

		for k, v := range m {
			paircount[k] = v
		}
	}
}

func expand(pair string, level int, currentLevel int, rules map[string]string, paircount map[string]int) {
	if level == currentLevel {
		return
	}

	paircount[pair]++

	left := string(pair[0]) + rules[pair]
	right := rules[pair] + string(pair[1])

	expand(left, level, currentLevel+1, rules, paircount)
	expand(right, level, currentLevel+1, rules, paircount)
}

func createInput(pairs []string) string {
	output := string(pairs[0][0])

	for _, pair := range pairs {
		output += pair[1:]
	}
	return output
}

func minAndMax(input string) (min, max int) {

	m := map[string]int{}

	for _, char := range input {
		m[string(char)]++
	}

	min = int(10e3)
	max = 0

	for _, v := range m {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return
}

func applyRules(input string, rules map[string]string) string {
	output := string(input[0])

	for i := 0; i < len(input)-1; i++ {
		pair := input[i : i+2]
		output += rules[pair] + string(pair[1])
	}

	//output += string(input[len(input)-1])

	return output
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
