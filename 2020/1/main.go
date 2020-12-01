package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%v\n", solve2())
}

func getNumbers() []int {
	in, _ := ioutil.ReadFile("input.txt")
	nums := []int{}
	parts := strings.Split(string(in), "\n")
	for _, part := range parts {
		i, err := strconv.Atoi(part)
		if err != nil {
			continue
		}
		nums = append(nums, i)
	}
	return nums
}

func solve2() int {
	is := getNumbers()
	sort.Slice(is, func(i, j int) bool { return is[i] < is[j] })

	for i := 0; i < len(is); i++ {
		for j := i + 1; j < len(is); j++ {
			for z := j + 1; z < len(is); z++ {
				a, b, c := is[i], is[j], is[z]
				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}
	return -1
}

func solve1() {
	in, _ := ioutil.ReadFile("input.txt")
	nums := map[int]bool{}
	parts := strings.Split(string(in), "\n")
	for _, part := range parts {
		i, err := strconv.Atoi(part)
		if err != nil {
			continue
		}
		delta := 2020 - i
		if _, ok := nums[delta]; ok {
			result := delta * i
			fmt.Printf("%v\n", result)
		}
		nums[i] = true
	}

	fmt.Println("done")
}
