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
	arrival, busses := getInput()

	// find the bus closes to arrival

	min := int(10e9)
	minID := 0
	for _, bus := range busses {
		nextBus := (bus - arrival%bus)
		if nextBus < min {
			min = nextBus
			minID = bus
		}
	}

	return minID * min
}

func getInput() (arrival int, bus []int) {
	in, _ := ioutil.ReadFile("input.txt")
	parts := strings.Split(string(in), "\n")
	as, bs := parts[0], parts[1]
	is, err := strconv.Atoi(as)
	if err != nil {
		panic(err)
	}
	arrival = is

	for _, p := range strings.Split(bs, ",") {
		if p == "x" {
			continue
		}
		bt, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}
		bus = append(bus, bt)
	}

	return
}
