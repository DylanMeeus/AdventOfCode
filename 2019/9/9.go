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
	readFunc := func() int { return 5 }
	var relativeBase int
	for i := 0; i < len(input); {
		codeparam := strconv.Itoa(input[i])
		var opcode string
		var mode1, mode2 string
		mode1, mode2 = "0", "0"
		if len(codeparam) != 1 {
			codeparam = "00" + codeparam
			opcode = string(codeparam[len(codeparam)-2]) + string(codeparam[len(codeparam)-1])
			mode1 = string(codeparam[len(codeparam)-3])
			mode2 = string(codeparam[len(codeparam)-4])
		} else {
			opcode = "0" + codeparam
		}
		switch opcode {
		case "99":
			return input
		case "01":
			ind1, ind2, store := input[i+1], input[i+2], input[i+3]
			a := ind1
			b := ind2
			if mode1 == "0" {
				a = input[ind1]
			}
			if mode2 == "0" {
				b = input[ind2]
			}
			input[store] = a + b
			i += 4
		case "02":
			ind1, ind2, store := input[i+1], input[i+2], input[i+3]
			a := ind1
			b := ind2
			if mode1 == "0" {
				a = input[ind1]
			}
			if mode2 == "0" {
				b = input[ind2]
			}
			input[store] = a * b
			i += 4
		case "03":
			store := input[i+1]
			input[store] = readFunc()
			i += 2
		case "04":
			store := input[i+1]
			fmt.Printf("%v\n", input[store])
			i += 2
		case "05":
			ind1, ind2 := input[i+1], input[i+2]
			a := ind1
			b := ind2
			if mode1 == "0" {
				a = input[ind1]
			}
			if mode2 == "0" {
				b = input[ind2]
			}
			if a != 0 {
				i = b
			} else {
				i += 3
			}
		case "06":
			ind1, ind2 := input[i+1], input[i+2]
			a := ind1
			b := ind2
			if mode1 == "0" {
				a = input[ind1]
			}
			if mode2 == "0" {
				b = input[ind2]
			}
			if a == 0 {
				i = b
			} else {
				i += 3
			}
		case "07":
			// less than
			ind1, ind2, store := input[i+1], input[i+2], input[i+3]
			a := ind1
			b := ind2
			if mode1 == "0" {
				a = input[ind1]
			}
			if mode2 == "0" {
				b = input[ind2]
			}
			if a < b {
				input[store] = 1
			} else {
				input[store] = 0
			}
			i += 4
		case "08":
			// equals
			ind1, ind2, store := input[i+1], input[i+2], input[i+3]
			a := ind1
			b := ind2
			if mode1 == "0" {
				a = input[ind1]
			}
			if mode2 == "0" {
				b = input[ind2]
			}
			if a == b {
				input[store] = 1
			} else {
				input[store] = 0
			}
			i += 4
		case "09":
			ind1 := input[i+1]
			relativeBase += ind1
			i += 2
		default:
			i++
		}
	}
	return input
}
