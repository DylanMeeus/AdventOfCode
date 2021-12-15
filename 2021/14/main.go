package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Printf("%v\n", solve())
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

	}

	min, max := minAndMax(input)

	return max - min
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
