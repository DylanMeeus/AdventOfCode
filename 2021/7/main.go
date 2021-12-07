package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%v\n", solve())
}

func getData() []int {
	in, _ := ioutil.ReadFile("./input.txt")

	body := strings.Replace(string(in), "\n", "", -1)
	parts := strings.Split(body, ",")

	out := make([]int, len(parts))
	var err error
	for i, p := range parts {
		out[i], err = strconv.Atoi(p)
		if err != nil {
			panic(err)
		}
	}
	return out
}

func solve() int {
	crabs := getData()
	lower, higher := min(crabs), max(crabs)

	minCost := int(10e7)
	for ; lower < higher; lower++ {
		cost := getCost(crabs, lower)
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost

}

func getCost(crabs []int, position int) int {
	totalCost := 0
	for _, c := range crabs {
		totalCost += int(math.Abs(float64(position - c)))
	}
	return totalCost
}

func min(crabs []int) int {
	m := crabs[0]
	for _, c := range crabs {
		if c < m {
			m = c
		}
	}
	return m
}

func max(crabs []int) (m int) {
	for _, c := range crabs {
		if c > m {
			m = c
		}
	}
	return
}
