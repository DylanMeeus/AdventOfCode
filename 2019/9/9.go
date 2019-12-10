package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	solve1()
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
	calculate(data)
}

func calculate(input []int) []int {
	input = append(input, make([]int, 34915192)...)
	readFunc := func() int { return 1 }
	var relativeBase int
	relativeBase = 0
	for i := 0; i < len(input); {
		codeparam := strconv.Itoa(input[i])
		var opcode string
		var mode1, mode2, mode3 string
		mode1, mode2, mode3 = "0", "0", "0"
		_ = mode3
		if len(codeparam) != 1 {
			codeparam = "000" + codeparam
			opcode = string(codeparam[len(codeparam)-2]) + string(codeparam[len(codeparam)-1])
			mode1 = string(codeparam[len(codeparam)-3])
			mode2 = string(codeparam[len(codeparam)-4])
			mode3 = string(codeparam[len(codeparam)-5])
		} else {
			opcode = "0" + codeparam
		}
		switch opcode {
		case "99":
			return input
		case "01":
			ind1, ind2, store := input[i+1], input[i+2], input[i+3]
			a := parseMode(mode1, relativeBase, ind1, input)
			b := parseMode(mode2, relativeBase, ind2, input)
			//store = parseMode(mode3, relativeBase, store, input)
			input[store] = a + b
			i += 4
		case "02":
			ind1, ind2, store := input[i+1], input[i+2], input[i+3]
			a := parseMode(mode1, relativeBase, ind1, input)
			b := parseMode(mode2, relativeBase, ind2, input)
			store = parseMode(mode3, relativeBase, store, input)
			input[store] = a * b
			i += 4
		case "03":
			ind := input[i+1]
			if mode1 == "2" {
				fmt.Println("woop")
				fmt.Println(ind)
				fmt.Println(input[ind])
				input[relativeBase+input[ind]] = readFunc()
			} else if mode1 == "1" {
				fmt.Println("fuck you")
			} else {
				fmt.Println("woopwoop")
				input[ind] = readFunc()
			}
			i += 2
		case "04":
			store := input[i+1]
			a := store
			if mode1 == "2" {
				fmt.Printf("%v\n", input[relativeBase+a])
			} else {
				fmt.Printf("%v\n", input[a])
			}
			i += 2
		case "05":
			ind1, ind2 := input[i+1], input[i+2]
			a := parseMode(mode1, relativeBase, ind1, input)
			b := parseMode(mode2, relativeBase, ind2, input)
			if a != 0 {
				i = b
			} else {
				i += 3
			}
		case "06":
			ind1, ind2 := input[i+1], input[i+2]
			a := parseMode(mode1, relativeBase, ind1, input)
			b := parseMode(mode2, relativeBase, ind2, input)
			if a == 0 {
				i = b
			} else {
				i += 3
			}
		case "07":
			// less than
			ind1, ind2, store := input[i+1], input[i+2], input[i+3]
			a := parseMode(mode1, relativeBase, ind1, input)
			b := parseMode(mode2, relativeBase, ind2, input)
			//store = parseMode(mode3, relativeBase, store, input)
			if a < b {
				input[store] = 1
			} else {
				input[store] = 0
			}
			i += 4
		case "08":
			// equals
			ind1, ind2, store := input[i+1], input[i+2], input[i+3]
			a := parseMode(mode1, relativeBase, ind1, input)
			b := parseMode(mode2, relativeBase, ind2, input)
			//store = parseMode(mode3, relativeBase, store, input)
			if a == b {
				input[store] = 1
			} else {
				input[store] = 0
			}
			i += 4
		case "09":
			alt := parseMode(mode1, relativeBase, input[i+1], input)
			relativeBase += alt
			i += 2
		default:
			i++
		}
	}
	return input
}

// return the location of the blahblabhlah?
func parseMode(mode string, relbase, value int, input []int) int {
	switch mode {
	case "0":
		return input[value]
	case "1":
		return value
	case "2":
		return input[relbase+value]
	default:
		panic("fubar")
	}
}
