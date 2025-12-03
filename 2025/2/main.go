package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	numbers := parseInput(readInput())
	//fmt.Println(solve1(numbers))

	//numbers = []int{1010}
	fmt.Println(solve2(numbers))
}

func solve1(is []int) int {
	c := 0

	for _, number := range is {
		strNum := strconv.Itoa(number)
		if len(strNum)%2 != 0 {
			// no need to check this because a repetition is impossible
			continue
		}
		a := strNum[0 : len(strNum)/2]
		b := strNum[len(strNum)/2 : len(strNum)]
		if a == b {
			fmt.Println(number)
			c += number
		}
	}

	return c
}

func solve2(is []int) int {
	c := 0

outer:
	for _, number := range is {
		strNum := strconv.Itoa(number)
		// find all possible sequences and see if they repeat..
		for i := 0; i < (len(strNum) / 2); i++ {
			for j := i + 1; j < len(strNum); j++ {
				sequence := strNum[i:j]
				re := regexp.MustCompile(fmt.Sprintf("(%s){%d}", sequence, len(strNum)/len(sequence)))
				matches := re.FindString(strNum) == strNum
				//fmt.Printf("checking sequence: %v against %v\n", strNum, re)
				if matches {
					c += number
					continue outer
				}
			}
		}
	}

	return c
}

func parseInput(s string) []int {
	out := []int{}
	ranges := strings.Split(s, ",")

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		from, to := parts[0], parts[1]
		fi, err := strconv.Atoi(from)
		if err != nil {
			panic(err)
		}
		ti, err := strconv.Atoi(to)
		if err != nil {
			panic(err)
		}

		for i := fi; i <= ti; i++ {
			out = append(out, i)
		}
	}

	return out
}

func readInput() string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")[0]
}
