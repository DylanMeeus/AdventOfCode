package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Group struct {
	Answer string
	Len    int
}

func main() {
	fmt.Printf("%v\n", solve1())
	fmt.Printf("%v\n", solve2())
}

func solve1() int {
	in := getInput()
	groups := reduce(in)

	out := 0
	for _, g := range groups {
		if g.Answer == "" {
			continue
		}
		out += countDistinct(g.Answer)
	}

	return out
}

func solve2() int {
	in := getInput()
	groups := reduce(in)

	out := 0
	for _, g := range groups {
		if g.Answer == "" {
			continue
		}
		freq := countFreq(g.Answer)
		for _, v := range freq {
			if v == g.Len {
				out++
			}
		}
	}

	return out
}

func countFreq(ss string) map[string]int {
	m := map[string]int{}
	for _, s := range ss {
		m[string(s)]++
	}
	return m
}

func countDistinct(ss string) (out int) {
	m := map[string]bool{}
	for _, s := range ss {
		m[string(s)] = true
	}
	return len(m)
}

func reduce(in []string) []Group {
	current := in[0]
	out := []Group{}
	size := 1
	for i := 1; i < len(in); i++ {
		if in[i] == "" {
			out = append(out, Group{current, size})
			size = 0
			current = ""
		} else {
			current += in[i]
			size++
		}
	}
	return out
}

func getInput() []string {
	in, _ := ioutil.ReadFile("input.txt")
	return strings.Split(string(in), "\n")

}
