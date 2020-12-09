package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%v\n", solve1())
}

func solve1() int {
	nums := getInput()

	for i := 25; i < len(nums); i++ {
		if !valid(nums[i-25:i], nums[i]) {
			return nums[i]
		}
	}

	return -1
}

func valid(nums []int, current int) bool {
	m := map[int]bool{}

	for _, num := range nums {
		delta := current - num
		if _, ok := m[delta]; ok {
			return true
		}
		m[num] = true
	}

	return false
}

func getInput() []int {
	in, _ := ioutil.ReadFile("input.txt")
	out := []int{}
	for _, line := range strings.Split(string(in), "\n") {
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		out = append(out, num)
	}
	return out
}
