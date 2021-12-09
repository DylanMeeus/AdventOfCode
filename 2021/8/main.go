package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

var (
	digitToIndices = map[int][]int{
		0: []int{0, 1, 2, 4, 5, 6},
		1: []int{2, 5},
		2: []int{0, 2, 3, 4, 6},
		3: []int{0, 2, 3, 5, 6},
		4: []int{1, 3, 2, 5},
		5: []int{0, 1, 3, 5, 6},
		6: []int{0, 1, 3, 4, 5, 6},
		7: []int{0, 2, 5},
		8: []int{0, 1, 2, 3, 4, 5, 6, 7},
		9: []int{0, 1, 2, 3, 5, 6, 7},
	}
)

type InputLine struct {
	Signals []string
	Outputs []string
}

func main() {
	fmt.Printf("%v\n", solve())
	fmt.Printf("%v\n", solve2())
}

func getData() []InputLine {
	in, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(in), "\n")
	result := make([]InputLine, len(lines))
	for i, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "|")
		signals := strings.Split(parts[0], " ")
		outputs := strings.Split(parts[1], " ")
		result[i] = InputLine{signals, outputs}
	}
	return result
}

func solve() int {
	data := getData()

	validLength := map[int]bool{
		2: true, // represents a 1
		3: true, // represents a 7
		4: true, // represents a 4
		7: true, // represents a 8
	}

	sum := 0

	for _, inputLine := range data {
		for _, output := range inputLine.Outputs {
			if validLength[len(output)] {
				sum++
			}
		}
	}

	return sum
}

func solve2() int {
	data := getData()
	fmt.Printf("%v\n", data[0])

	/*

		segment looks like:

		    0
		1	2
		    3
		4	5
		    6

	*/

	sum := 0
	for _, line := range data {
		fmt.Printf("line: %v\nmappin", line)
		if len(line.Signals) == 0 {
			continue
		}
		mapping := createMap(line.Signals)
		value := 0
		for i, output := range line.Outputs {
			strOut := strings.Split(output, "")
			sort.Strings(strOut)
			ss := strings.Join(strOut, "")
			value += mapping[ss] * int(math.Pow(10, 4.-float64(i)))
		}
		sum += value
	}
	return sum
}

// createMap maps a string ('segment code') to a value
func createMap(signals []string) map[string]int {
	// we can find all the numbers by exclusion

	length := func(i int) func(s string) bool {
		return func(s string) bool { return len(s) == i }
	}

	ones := filter(signals, length(2))
	fours := filter(signals, length(4))
	sevens := filter(signals, length(3))
	eights := filter(signals, length(7))

	threes := filter(signals, func(s string) bool {
		if len(s) != 5 {
			return false
		}
		// 3 is superset of 7, so it needs to contain all of the 7s
		return isSuperSet(s, sevens)
	})

	nines := filter(signals, func(s string) bool {
		if len(s) != 6 {
			return false
		}

		return isSuperSet(s, threes)
	})

	sixes := filter(signals, func(s string) bool {
		if len(s) != 6 {
			return false
		}

		for _, one := range ones {
			if !containsAll(s, strings.Split(one, "")) {
				return true
			}
		}
		return false
	})

	zeroes := filter(signals, func(s string) bool {
		if len(s) != 6 {
			return false
		}

		thisString := func(that string) bool { return that == s }
		return len(filter(sixes, thisString)) == 0 && len(filter(nines, thisString)) == 0
	})

	fives := filter(signals, func(s string) bool {
		if len(s) != 5 {
			return false
		}
		for _, six := range sixes {
			if !isSuperSet(six, strings.Split(s, "")) {
				return false
			}
		}
		return true
	})

	twos := filter(signals, func(s string) bool {
		if len(s) != 5 {
			return false
		}

		thisString := func(that string) bool { return that == s }
		return len(filter(threes, thisString)) == 0 && len(filter(fives, thisString)) == 0
	})
	/*

		fmt.Printf("%v\n", ones)
		fmt.Printf("%v\n", fours)
		fmt.Printf("%v\n", sevens)
		fmt.Printf("threes %v\n", threes)
		fmt.Printf("%v\n", eights)
		fmt.Printf("%v\n", nines)
		fmt.Printf("%v\n", sixes)
		fmt.Printf("%v\n", zeroes)
		fmt.Printf("%v\n", fives)
		fmt.Printf("%v\n", twos)
	*/

	badMap := map[string]int{
		zeroes[0]: 0,
		ones[0]:   1,
		twos[0]:   2,
		threes[0]: 3,
		fours[0]:  4,
		fives[0]:  5,
		sixes[0]:  6,
		sevens[0]: 7,
		eights[0]: 8,
		nines[0]:  9,
	}

	sortedMap := map[string]int{}
	for k, v := range badMap {
		ss := strings.Split(k, "")
		sort.Strings(ss)
		sortedK := strings.Join(ss, "")
		sortedMap[sortedK] = v
	}
	return sortedMap
}

// checks for a string superset
func isSuperSet(super string, subs []string) bool {
	for _, sub := range subs {
		if !containsAll(super, strings.Split(sub, "")) {
			return false
		}
	}
	return true
}

func containsAll(stack string, needles []string) bool {
	for _, needle := range needles {
		if !strings.Contains(stack, needle) {
			return false
		}
	}
	return true
}

func containsAny(stack string, needles []string) bool {
	for _, needle := range needles {
		if strings.Contains(stack, needle) {
			return true
		}
	}
	return false
}
func reducePossibilities(possibleAtIndex map[int][]string, possibleForDigit map[int][]string) map[int][]string {

	return nil
}

func filter(stack []string, predicate func(string) bool) []string {
	out := []string{}
	for _, s := range stack {
		if predicate(s) {
			out = append(out, s)
		}
	}
	return out
}

func contains(stack []string, needle string) bool {
	for _, s := range stack {
		if s == needle {
			return true
		}
	}
	return false
}
