package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const (
	NOOP = "noop"
	ADD  = "addv"
)

type instruction struct {
	operation string
	parameter int
}

func main() {
	fmt.Println(solve1())
	solve2()
}

func solve1() int {

	data := getData()

	acc := 1
	cycle := 0

	signal := 0

	incrSignal := func() {
		if cycle == 20 || (cycle > 40 && (cycle-20)%40 == 0) {
			fmt.Printf("cycle: %v with acc %v = %v\n", cycle, acc, cycle*acc)
			signal += cycle * acc
		}
	}

	for _, in := range data {
		if in.operation == NOOP {
			cycle++
			incrSignal()
			continue
		} else {
			tmp := cycle
			for cycle < tmp+2 {
				cycle++
				incrSignal()
			}
			acc += in.parameter
		}

	}

	return signal
}

func solve2() {

	data := getData()

	acc := 1
	cycle := 0

	output := [][]bool{}
	for i := 0; i < 6; i++ {
		output = append(output, make([]bool, 40))
	}

	row, col := 0, 0
	draw := func() {
		fmt.Printf("%v - %v \n", col, acc)
		if math.Abs(float64(col-acc)) < 2 {
			output[row][col] = true
		}
		col++
		if col > 39 {
			row++
			col = 0
		}
	}

	for _, in := range data {
		if in.operation == NOOP {
			cycle++
			draw()
			continue
		} else {
			tmp := cycle
			for cycle < tmp+2 {
				cycle++
				draw()
			}
			acc += in.parameter
		}

	}

	print(output)
}

func print(input [][]bool) {

	for row := 0; row < 6; row++ {
		for col := 0; col < 40; col++ {
			if input[row][col] {
				fmt.Print(" # ")
			} else {
				fmt.Print(" . ")
			}
		}
		fmt.Println()
	}

}

func getData() []instruction {
	f, err := ioutil.ReadFile("./input.txt")
	handleError(err)

	instructions := []instruction{}
	lines := strings.Split(string(f), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		if line == "noop" {
			instructions = append(instructions, instruction{operation: NOOP})
		} else {
			parts := strings.Split(line, " ")
			num, err := strconv.Atoi(parts[1])
			handleError(err)
			instructions = append(instructions, instruction{operation: ADD, parameter: num})
		}
	}
	return instructions

}

func handleError(err error) {
	if err != nil {
		panic(err)
	}

}
