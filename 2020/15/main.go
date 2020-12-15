package main

import (
	"fmt"
)

var m map[int][]int

func main() {
	m = map[int][]int{} // map of key -> indices
	fmt.Printf("%v\n", solve2())
}

func next(in []int) int {
	last := in[len(in)-1]

	if indices, ok := m[last]; !ok || len(indices) == 1 {
		return 0
	} else {
		i1, i2 := indices[len(indices)-1], indices[len(indices)-2]
		return i1 - i2
	}
}

func solve2() int {
	numbers := getInput()
	for i, n := range numbers {
		m[n] = append(m[n], i)
	}
	var n int
	for i := len(numbers); i != 30000000; i++ {
		n = next(numbers)
		numbers = append(numbers, n)
		m[n] = append(m[n], i)
	}
	return n
}

func getInput() []int {
	return []int{2, 0, 6, 12, 1, 3}
}
