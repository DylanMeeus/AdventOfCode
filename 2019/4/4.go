package main

import (
	"fmt"
	"strconv"
)

func main () {
	fmt.Printf("%v\n", solve1())
}

func solve1() (sum int){
	min, max := 136760, 595730
	for min <= max {
		str := strconv.Itoa(min)
		if increasing(str) && containsDouble(str) {
			sum++
		}
		min++
	}
	return
}

func increasing(ss string) bool {
	for i := 0; i < len(ss) - 1; i++ {
		a, _:= strconv.Atoi(string(ss[i]))
		b, _:= strconv.Atoi(string(ss[i+1]))
		if a > b {
			return false
		}
	}
	return true
}

func containsDouble(ss string) bool {
	for i := 0; i < len(ss) - 1; i++ {
		if ss[i] == ss[i+1] {
			return true
		}
	}
	return false
}
