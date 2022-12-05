package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println(solve1())
}

func solve1() int {
	data := getData()

	sum := 0
	for _, line := range data {
		if line == "" {
			continue
		}
		a := line[0 : len(line)/2]
		b := line[len(line)/2:]

		char := getMatchingChar(a, b)
		var priority int
		if char > 90 {
			priority = int(char - 96)
		} else {
			priority = int(26 + (char - 64))
		}
		sum += priority
	}

	return sum
}

func getMatchingChar(a, b string) rune {
	for _, x := range a {
		for _, y := range b {
			if x == y {
				return x
			}
		}
	}
	panic("no similar char")
}

func getData() []string {
	f, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(f), "\n")
}
