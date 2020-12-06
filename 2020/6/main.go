package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Printf("%v\n", solve1())
}

func solve1() int {
	in := getInput()
	groups := reduce(in)

	out := 0
	for _, g := range groups {
		if g == "" {
			continue
		}
		out += countDistinct(g)
	}

	return out
}

func countDistinct(ss string) (out int) {
	m := map[string]bool{}
	for _, s := range ss {
		m[string(s)] = true
	}
	return len(m)
}

func reduce(in []string) []string {
	current := in[0]
	out := []string{}
	for i := 1; i < len(in); i++ {
		if in[i] == "" {
			out = append(out, current)
			current = ""
		} else {
			current += in[i]
		}
	}
	return out
}

func getInput() []string {
	in, _ := ioutil.ReadFile("input.txt")
	return strings.Split(string(in), "\n")

}
