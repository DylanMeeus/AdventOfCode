package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%v\n", solve1())
	fmt.Printf("%v\n", solve2())
}

func solve2() int {
	key := solve1()
	nums := getInput()

	// find contigious set which sums to this..

	for i := 0; i < len(nums); i++ {
		sum := 0
		is := []int{}
		for j := i; j < len(nums); j++ {
			is = append(is, nums[j])
			sum += nums[j]
			if sum == key && len(nums) > 2 {
				return min(is) + max(is)
			}
		}
	}

	return -1
}

func min(is []int) int {
	m := is[0]

	for _, i := range is {
		if i < m {
			m = i
		}
	}
	return m
}

func max(is []int) (out int) {
	for _, i := range is {
		if i > out {
			out = i
		}
	}
	return
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
