package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instruction struct {
	op    string
	value int
}

func main() {
	fmt.Printf("%v\n", solve1())
	fmt.Printf("%v\n", solve2())
}

func solve1() int {
	reg := 0
	iptr := 0
	seen := map[int]bool{}
	instructions := getInput()

	for iptr < len(instructions) {
		if _, ok := seen[iptr]; ok {
			return reg
		}
		seen[iptr] = true
		instr := instructions[iptr]
		switch instr.op {
		case "acc":
			reg += instr.value
			iptr++
		case "jmp":
			iptr += instr.value
		case "nop":
			iptr++
		}
	}
	return reg
}

func solve2() int {
	instructions := getInput()
	for i, inst := range instructions {
		original := inst
		if inst.op == "jmp" {
			instructions[i].op = "nop"
		} else if inst.op == "nop" {
			instructions[i].op = "jmp"
		}
		value, loop := loops(instructions)
		if !loop {
			fmt.Printf("%v\n", value)
			return value
		}
		instructions[i] = original
	}

	return -1
}

func loops(instructions []instruction) (int, bool) {
	iptr := 0
	reg := 0
	seen := map[int]bool{}
	for iptr < len(instructions) {
		if _, ok := seen[iptr]; ok {
			return -1, true
		}
		seen[iptr] = true
		instr := instructions[iptr]
		switch instr.op {
		case "acc":
			reg += instr.value
			iptr++
		case "jmp":
			iptr += instr.value
		case "nop":
			iptr++
		}
	}
	return reg, false
}

func getInput() []instruction {
	in, _ := ioutil.ReadFile("input.txt")

	instructions := []instruction{}
	for _, line := range strings.Split(string(in), "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		op, val := parts[0], parts[1]
		ival, _ := strconv.Atoi(val)
		instructions = append(instructions, instruction{op, ival})
	}
	return instructions

}
