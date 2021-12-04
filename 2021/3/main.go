package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type position int
type bitcounter map[string]int

func main() {
	fmt.Printf("%v\n", solve())
	fmt.Printf("%v\n", solve2())
}

func getData() []string {
	in, _ := ioutil.ReadFile("./input.txt")
	return strings.Split(string(in), "\n")
}

func solve() int {
	m := getBitCounts(getData())
	gamma := make([]bool, len(m))
	epsilon := make([]bool, len(m))

	for i := 0; i < len(m); i++ {
		gamma[i] = m[position(i)]["1"] > m[position(i)]["0"]
		epsilon[i] = !gamma[i]
	}
	// epsilon is the opposite of gamma

	return bitsToNum(gamma) * bitsToNum(epsilon)
}

func solve2() int64 {
	data := getData()

	oxygenTarget, scrubberTarget := getTargets(getData())

	oxygenCandidates := make([]string, len(data))
	scrubberCandidates := make([]string, len(data))

	copy(oxygenCandidates, data)
	copy(scrubberCandidates, data)

	var oxygenValue, scrubberValue int64
	var err error
	for col := 0; col < len(data[0]); col++ {
		oxygenCandidates = filter(oxygenCandidates, func(s string) bool {
			return string(s[col]) == oxygenTarget[col]
		})

		//fmt.Printf("oxygen %v\n", oxygenCandidates)
		if len(oxygenCandidates) == 1 {
			oxygenValue, err = strconv.ParseInt(oxygenCandidates[0], 2, 64)
			break
		}
		if err != nil {
			panic(err)
		}
		oxygenTarget, scrubberTarget = getTargets(oxygenCandidates)
	}

	for col := 0; col < len(data[0]); col++ {
		scrubberCandidates = filter(scrubberCandidates, func(s string) bool {
			return string(s[col]) == scrubberTarget[col]
		})

		if len(scrubberCandidates) == 1 {
			scrubberValue, err = strconv.ParseInt(scrubberCandidates[0], 2, 64)
			break
		}
		if err != nil {
			panic(err)
		}
		oxygenTarget, scrubberTarget = getTargets(scrubberCandidates)

	}

	fmt.Printf("oxygen target: %v\n", oxygenTarget)
	fmt.Printf("%v\n", scrubberTarget)

	fmt.Printf("oxygen: %v\nscrubber: %v\n", oxygenValue, scrubberValue)

	return oxygenValue * scrubberValue
}

func getTargets(data []string) (oxygenTarget []string, scrubberTarget []string) {
	m := getBitCounts(data)

	oxygenTarget = make([]string, len(m))
	scrubberTarget = make([]string, len(m))

	for i := 0; i < len(m); i++ {
		if m[position(i)]["1"] >= m[position(i)]["0"] {
			oxygenTarget[i] = "1"
		} else {
			oxygenTarget[i] = "0"
		}

		if m[position(i)]["0"] <= m[position(i)]["1"] {
			scrubberTarget[i] = "0"
		} else {
			scrubberTarget[i] = "1"
		}
	}

	return oxygenTarget, scrubberTarget
}

func filter(input []string, pred func(s string) bool) []string {
	out := []string{}
	for _, s := range input {
		if s == "" {
			continue
		}
		if pred(s) {
			out = append(out, s)
		}
	}
	return out
}

func getBitCounts(data []string) map[position]bitcounter {
	m := map[position]bitcounter{}

	for _, line := range data {
		for i, char := range strings.Split(line, "") {
			if m[position(i)] == nil {
				m[position(i)] = bitcounter{}
			}
			m[position(i)][char]++
		}
	}

	return m
}

func bitsToNum(bits []bool) int {
	result := 0
	n := len(bits) - 1
	for i := 0; i < len(bits); i++ {
		if bits[i] {
			result += int(math.Pow(2, float64(n)))
		}
		n--
	}
	return result
}
