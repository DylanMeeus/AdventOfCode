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

func solve1() int {
	arrival, busses := getInput()

	// find the bus closes to arrival

	min := int(10e9)
	minID := 0
	for _, bus := range busses {
		if bus == 0 {
			continue
		}
		nextBus := (bus - arrival%bus)
		if nextBus < min {
			min = nextBus
			minID = bus
		}
	}

	return minID * min
}

func solve2() int {
	_, busses := getInput()

	start := busses[0]
	soFar := []int{busses[0]}
	for _, bus := range busses[1:] {
		time := findAllignmentTime(soFar, bus, start)
		soFar = append(soFar, bus)
		start = time
	}
	return start

}

func findAllignmentTime(busses []int, bus, start int) int {
	incr := lcm(busses)
	c := make([]int, len(busses))
	copy(c, busses)
	c = append(c, bus)
	for {
		if checkBusses(c, start) {
			return start
		}
		start += incr
	}

}

func lcm(busses []int) int {
	out := 1
	for _, bus := range busses {
		if bus == 0 {
			continue
		}
		out *= bus
	}
	return out
}

func checkBusses(busses []int, start int) bool {
	for i, bus := range busses {
		if bus == 0 {
			continue
		}
		expected := start + i
		if expected%bus != 0 {
			return false
		}
	}
	return true
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
			bus = append(bus, 0)
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
