package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	test = []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
)

func main() {
	solve1()
	res := solve2()
	fmt.Printf("%v\n", res)
}

func readData() (out []int) {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	ss := string(data)
	for _, s := range strings.Split(strings.ReplaceAll(ss, "\n", ""), ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		out = append(out, i)
	}
	return
}

func solve1() {
	data := readData()
	data[1] = 12
	data[2] = 2
	data = calculate(data)
	fmt.Printf("%v\n", data[0])
}

func solve2() int {
	target := 19690720
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			data := readData()
			data[1] = a
			data[2] = b
			data = calculate(data)
			if data[0] == target {
				return (100 * a) + b
			}
		}
	}
	return 0
}

func calculate(input []int) []int {
	for i := 0; i < len(input); i += 4 {
		opcode := input[i]
		ind1, ind2, store := input[i+1], input[i+2], input[i+3]
		switch opcode {
		case 99:
			return input
		case 1:
			input[store] = input[ind1] + input[ind2]
		case 2:
			input[store] = input[ind1] * input[ind2]
		}
	}
	return input
}
