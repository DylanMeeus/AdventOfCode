package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	// rock beats scissors
	// scissors beats paper
	// paper beats rock
	inputToPts = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,

		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	winFrom = map[string]string{
		"X": "C",
		"Y": "A",
		"Z": "B",
	}

	toWin = map[string]string{
		"A": "Y", // paper wins from rock
		"B": "Z", // scissors win from paper
		"C": "X", // rock wins from scissors
	}

	toLose = map[string]string{
		"A": "Z", // rock wins from scissors
		"B": "X", // paper wins from rick
		"C": "Y", // scissors win from paper
	}
)

func main() {
	fmt.Println(solve1())
	fmt.Println(solve2())
}

func solve2() int {
	data := getData()

	sum := 0
	for _, line := range data {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		opp := parts[0]
		stats := parts[1]

		if stats == "X" { // lose round
			sum += inputToPts[toLose[opp]]
		} else if stats == "Y" { // draw
			sum += inputToPts[opp] + 3
		} else if stats == "Z" { // win
			sum += inputToPts[toWin[opp]] + 6
		}

	}
	return sum

}

func solve1() int {
	data := getData()

	sum := 0
	for _, line := range data {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		opp := parts[0]
		me := parts[1]
		sum += inputToPts[me]
		if inputToPts[opp] == inputToPts[me] {
			sum += 3
		} else if winFrom[me] == opp {
			sum += 6
		}
	}
	return sum

}

func getData() []string {
	f, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(f), "\n")
}
