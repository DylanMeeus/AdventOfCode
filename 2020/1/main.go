package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%v\n", solve1())
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
			delta := 2020 - (is[i] + is[j])
			if binSearch(delta, is[j:]) {
				return is[i] * is[j] * delta
			}
		}
	}
	return -1
}

func binSearch(target int, nums []int) bool {
	lo, hi := 0, len(nums)
	for lo <= hi {
		mid := lo + ((hi - lo) / 2)
		if target > nums[mid] {
			lo = mid + 1
		} else if target < nums[mid] {
			hi = mid - 1
		} else {
			return true
		}
	}
	return false
}

func solve1() int {
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
			return delta * i
		}
		nums[i] = true
	}

	return -1
}
