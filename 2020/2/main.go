package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%v\n", solve1())
	fmt.Printf("%v\n", solve2())
}

func solve1() (out int) {
	input := getInput()

	for _, in := range input {
		parts := strings.Split(in, ":")
		if len(parts) != 2 {
			continue
		}
		rule, pass := parts[0], parts[1]

		ruleParts := strings.Split(rule, " ")
		bounds, char := ruleParts[0], ruleParts[1]

		boundParts := strings.Split(bounds, "-")
		start, end := boundParts[0], boundParts[1]
		si, e1 := strconv.Atoi(start)
		ei, e2 := strconv.Atoi(end)
		if e1 != nil || e2 != nil {
			continue
		}
		c := count(char, pass)
		if c >= si && c <= ei {
			out++
		}
	}
	return out
}

func solve2() (out int) {
	input := getInput()

	for _, in := range input {
		parts := strings.Split(in, ":")
		if len(parts) != 2 {
			continue
		}
		rule, pass := parts[0], parts[1]
		pass = strings.TrimSpace(pass)

		ruleParts := strings.Split(rule, " ")
		bounds, char := ruleParts[0], ruleParts[1]

		boundParts := strings.Split(bounds, "-")
		start, end := boundParts[0], boundParts[1]
		si, e1 := strconv.Atoi(start)
		ei, e2 := strconv.Atoi(end)
		if e1 != nil || e2 != nil {
			continue
		}

		si--
		ei--

		a, b := string(pass[si]), string(pass[ei])
		if string(a) == char && string(b) == char {
			continue
		}
		if string(a) == char || string(b) == char {
			out++
		}

	}
	return out
}

func count(c, s string) (out int) {
	for _, char := range s {
		if string(char) == c {
			out++
		}
	}
	return
}

func getInput() []string {
	input, _ := ioutil.ReadFile("input.txt")
	return strings.Split(string(input), "\n")
}
