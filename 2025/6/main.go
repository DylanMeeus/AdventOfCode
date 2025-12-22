package main

import (
	"fmt"
	"io/ioutil"
	"iter"
	"maps"
	"regexp"
	"strconv"
	"strings"
)

type calcfunc func(int, int) int

var (
	reducer = map[string]calcfunc{
		"+": calcfunc(func(a, b int) int { return a + b }),
		"*": calcfunc(func(a, b int) int { return a * b }),
	}
)

type Problem struct {
	values   []int
	operator string
}

func (p *Problem) result() int {
	acc := p.values[0]
	for _, value := range p.values[1:] {
		acc = reducer[p.operator](acc, value)
	}

	return acc
}

func main() {
	problems := parse(preprocess(readInput()))
	fmt.Println(solve1(problems))
}

func solve1(problems iter.Seq[*Problem]) int {
	c := 0
	for p := range problems {
		c += p.result()
	}
	return c
}

func parse(lines []string) iter.Seq[*Problem] {
	m := map[int]*Problem{}
	for linecount, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		for i, part := range parts {
			if part == "" {
				continue
			}
			if m[i] == nil {
				m[i] = &Problem{values: []int{}, operator: ""}
			}
			if linecount == len(lines)-2 {
				m[i].operator = part
			} else {
				val, err := strconv.Atoi(part)
				if err != nil {
					panic(err)
				}
				m[i].values = append(m[i].values, val)
			}
		}
	}
	// return all values
	return maps.Values(m)
}

func preprocess(lines []string) []string {
	// remove the superfluous spaces and replace with a single space
	out := []string{}
	rgx, err := regexp.Compile(" +")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		newline := rgx.ReplaceAllString(line, " ")
		out = append(out, newline)
	}
	return out
}

func readInput() []string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}
