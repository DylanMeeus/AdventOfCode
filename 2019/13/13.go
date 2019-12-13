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

type tile struct {
	x, y, block int
}

func solve1() {
	data := readData()
	output := calculate(data)
	tiles := []tile{}
	for i := 0; i < len(output)-4; i++ {
		x := output[i]
		i++
		y := output[i]
		i++
		block := output[i]
		tiles = append(tiles, tile{x, y, block})
	}
	const blockTile = 2
	var out int
	for _, t := range tiles {
		if t.block == blockTile {
			out++
		}
	}
	fmt.Printf("%v\n", out)
}

func calculate(input []int) []int {
	input = append(input, make([]int, 3000)...)
	readFunc := func() int { return 2 }
	var relativeBase int
	relativeBase = 0
	out := []int{}
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
			return out
		case "01":
			ind1, ind2, store := input[i+1], input[i+2], input[i+3]
			a := parseMode(mode1, relativeBase, ind1, input)
			b := parseMode(mode2, relativeBase, ind2, input)
			input = storeMode(mode3, a+b, relativeBase, store, input)
			i += 4
		case "02":
			ind1, ind2, store := input[i+1], input[i+2], input[i+3]
			a := parseMode(mode1, relativeBase, ind1, input)
			b := parseMode(mode2, relativeBase, ind2, input)
			input = storeMode(mode3, a*b, relativeBase, store, input)
			i += 4
		case "03":
			ind := input[i+1]
			if mode1 == "2" {
				fmt.Printf("mode 2\n")
				input[relativeBase+ind] = readFunc()
			} else if mode1 == "1" {
				fmt.Printf("here..\n")
			} else {
				fmt.Printf("mode 0\n")
				input[ind] = readFunc()
			}
			i += 2
		case "04":
			store := input[i+1]
			a := store
			var output int
			if mode1 == "2" {
				output = input[relativeBase+a]
			} else if mode1 == "1" {
				output = store
			} else {
				// mode 0
				output = input[a]
			}
			out = append(out, output)
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
			if a < b {
				input = storeMode(mode3, 1, relativeBase, store, input)
			} else {
				input = storeMode(mode3, 0, relativeBase, store, input)
			}
			i += 4
		case "08":
			// equals
			ind1, ind2, store := input[i+1], input[i+2], input[i+3]
			a := parseMode(mode1, relativeBase, ind1, input)
			b := parseMode(mode2, relativeBase, ind2, input)
			//store = parseMode(mode3, relativeBase, store, input)
			if a == b {
				input = storeMode(mode3, 1, relativeBase, store, input)
			} else {
				input = storeMode(mode3, 0, relativeBase, store, input)
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

func storeMode(mode string, value, relbase, storeLocation int, input []int) []int {
	switch mode {
	case "0":
		input[storeLocation] = value
		return input
	case "1":
		input[input[storeLocation]] = value
		return input
	case "2":
		input[relbase+storeLocation] = value
		return input
	default:
		panic("fubar")
	}
}
