package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%v\n", solve1())
	fmt.Printf("%v\n", solve2())
}

func solve1() (sum int) {
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

func solve2() (sum int) {
	min, max := 136760, 595730
	for min <= max {
		str := strconv.Itoa(min)
		if increasing(str) && exactly2(str) {
			sum++
		}
		min++
	}
	return
}

// exactly 2 repeating digits
func exactly2(ss string) bool {
	ss = "a" + ss + "a"
	for i := 1; i < len(ss)-2; i++ {
		if ss[i-1] != ss[i] && ss[i] == ss[i+1] && ss[i+2] != ss[i] {
			return true
		}
	}
	return false
}

func increasing(ss string) bool {
	for i := 0; i < len(ss)-1; i++ {
		a, _ := strconv.Atoi(string(ss[i]))
		b, _ := strconv.Atoi(string(ss[i+1]))
		if a > b {
			return false
		}
	}
	return true
}

func containsDouble(ss string) bool {
	for i := 0; i < len(ss)-1; i++ {
		if ss[i] == ss[i+1] {
			return true
		}
	}
	return false
}
