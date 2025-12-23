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
	//fmt.Println(solve1(parse(preprocess(readInput()))))
	fmt.Println(solve2(parse2(readInput())))
}

func solve1(problems iter.Seq[*Problem]) int {
	c := 0
	for p := range problems {
		c += p.result()
	}
	return c
}

func solve2(ps []Problem) int {
	c := 0
	for _, p := range ps {
		if len(p.values) == 0 {
			continue
		}
		c += p.result()
	}
	return c
}

// parser for the second problem
func parse2(lines []string) []Problem {
	max := 0

	ps := []Problem{}

	rowmaps := []map[int]string{}
	for _, line := range lines {
		if len(line) > max {
			max = len(line)
		}
		rowmap := map[int]string{}
		for i, char := range line {
			rowmap[i] = string(char)
		}
		rowmaps = append(rowmaps, rowmap)
	}

	lastOperator := ""
	numbers := []string{}
	for i := 0; i < max; i++ {
		num := ""
		for _, row := range rowmaps {
			// now we have a full number or a list of spaces..
			char := row[i]
			if char == "+" || char == "*" {
				lastOperator = char
			} else {
				if char != " " {
					num += row[i]
				}
			}
		}
		if num == "" {
			p := Problem{values: mapToInt(numbers), operator: lastOperator}
			ps = append(ps, p)
			numbers = []string{}
		} else {
			numbers = append(numbers, num)
		}
		//fmt.Println(num)
	}
	// add last in buffer
	ps = append(ps, Problem{values: mapToInt(numbers), operator: lastOperator})
	return ps
}

func mapToInt(ss []string) []int {
	out := make([]int, len(ss))

	for j, s := range ss {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		out[j] = i
	}

	return out
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
