package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%v\n", solve1())
	fmt.Printf("%v\n", solve2())
}

func solve2() int {
	lines := getInput()

	mask := ""
	memmap := map[int][36]int{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "=")
		if strings.Contains(line, "mem") {
			oloc := extractMemloc(parts[0])
			allLocs := applyQuantumMask(toMemory(oloc), mask)
			_ = allLocs
			//fmt.Printf("mem loc: %v\n", loc)
			svalue := strings.TrimSpace(parts[1])
			ival, _ := strconv.Atoi(svalue)
			// turn it into binary and write it
			for _, rloc := range allLocs {
				loc := toValue(rloc)
				memmap[loc] = toMemory(ival)
			}

		} else {
			// apply mask
			mask = strings.TrimSpace(parts[1])
		}
	}

	out := 0
	for _, reg := range memmap {
		out += toValue(reg)
	}

	return out
}

func solve1() int {
	lines := getInput()

	mask := ""
	memmap := map[int][36]int{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "=")
		if strings.Contains(line, "mem") {
			loc := extractMemloc(parts[0])
			svalue := strings.TrimSpace(parts[1])
			ival, _ := strconv.Atoi(svalue)
			memmap[loc] = applyMask(toMemory(ival), mask)

		} else {
			// apply mask
			mask = strings.TrimSpace(parts[1])
		}
	}

	out := 0
	for _, reg := range memmap {
		out += toValue(reg)
	}
	return out
}

func toValue(in [36]int) int {
	out := 0
	p := 0.
	for i := 35; i >= 0; i-- {
		out += in[i] * int(math.Pow(2., p))
		p++
	}
	return out
}

// toMemory turns the value into 36b memory
func toMemory(i int) [36]int {
	binString := strconv.FormatInt(int64(i), 2)
	out := [36]int{}

	bp := len(binString) - 1
	for i := 35; i >= 0; i-- {
		if bp >= 0 && string(binString[bp]) == "1" {
			out[i] = 1
		} else {
			out[i] = 0
		}
		bp--

	}

	return out
}

// each bit is in all possible states :-)
func applyQuantumMask(result [36]int, mask string) [][36]int {
	masks := []string{}
	var generate func(int, string, string)
	generate = func(position int, current, remainder string) {
		if len(remainder) == 0 {
			masks = append(masks, current)
			return
		}
		var head, tail string
		head = string(remainder[0])
		if len(remainder) > 1 {
			tail = remainder[1:]
		}
		// appending depends on head

		if head == "1" {
			generate(position+1, current+"1", tail)
		} else if head == "0" {
			generate(position+1, current+fmt.Sprintf("%v", result[position]), tail)
		} else {
			generate(position+1, current+"1", tail)
			generate(position+1, current+"0", tail)
		}
	}

	generate(0, "", mask)

	out := [][36]int{}
	for _, mask := range masks {
		x := [36]int{}
		//fmt.Printf("%v - %v\n", mask, len(mask))
		// turn the mask into an int slice
		for i, s := range mask {
			if string(s) == "1" {
				x[i] = 1
			} else {
				x[i] = 0
			}
		}
		out = append(out, x)

	}
	return out
}

func applyMask(result [36]int, mask string) [36]int {
	for i, value := range mask {
		if string(value) == "1" {
			result[i] = 1
		} else if string(value) == "0" {
			result[i] = 0
		}
	}
	return result
}

func extractMemloc(s string) int {
	valid := map[string]bool{"0": true, "1": true, "2": true, "3": true, "4": true, "5": true, "6": true,
		"7": true, "8": true, "9": true}
	loc := strings.Map(func(r rune) rune {
		if _, ok := valid[string(r)]; !ok {
			return rune(-1)
		}
		return r
	}, s)

	il, err := strconv.Atoi(loc)
	if err != nil {
		panic(err)
	}
	return il
}

func getInput() []string {
	in, _ := ioutil.ReadFile("input.txt")
	return strings.Split(string(in), "\n")
}
