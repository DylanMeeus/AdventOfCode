package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type InputLine struct {
	Signals []string
	Outputs []string
}

func main() {
	fmt.Printf("%v\n", solve())
}

func getData() []InputLine {
	in, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(in), "\n")
	result := make([]InputLine, len(lines))
	for i, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "|")
		signals := strings.Split(parts[0], " ")
		outputs := strings.Split(parts[1], " ")
		result[i] = InputLine{signals, outputs}
	}
	return result
}

func solve() int {
	data := getData()
	fmt.Printf("%v\n", data)

	validLength := map[int]bool{
		2: true, // represents a 1
		3: true, // represents a 7
		4: true, // represents a 4
		7: true, // represents a 8
	}

	sum := 0

	for _, inputLine := range data {
		for _, output := range inputLine.Outputs {
			if validLength[len(output)] {
				sum++
			}
		}
	}

	return sum
}
