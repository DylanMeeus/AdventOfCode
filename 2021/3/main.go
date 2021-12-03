package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type position int
type bitcounter map[string]int

func main() {
	fmt.Printf("%v\n", solve())
	fmt.Printf("%v\n", bitsToNum([]bool{true, true, false, false}))
}

func getData() []string {
	in, _ := ioutil.ReadFile("./input.txt")
	return strings.Split(string(in), "\n")
}

func solve() int {
	data := getData()

	m := map[position]bitcounter{}

	for _, line := range data {
		for i, char := range strings.Split(line, "") {
			if m[position(i)] == nil {
				m[position(i)] = bitcounter{}
			}
			m[position(i)][char]++
		}
	}

	gamma := make([]bool, len(m))
	epsilon := make([]bool, len(m))

	for i := 0; i < len(m); i++ {
		gamma[i] = m[position(i)]["1"] > m[position(i)]["0"]
		epsilon[i] = !gamma[i]
	}
	// epsilon is the opposite of gamma

	return bitsToNum(gamma) * bitsToNum(epsilon)
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

func solve2() {
}
