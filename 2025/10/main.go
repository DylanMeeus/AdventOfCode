package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Machine struct {
	desiredState map[int]bool // desiredState as a map of [index]bool
	buttons      [][]int
	joltage      []int
}

func copyMap(m map[int]bool) map[int]bool {
	out := map[int]bool{}
	for k, v := range m {
		out[k] = v
	}
	return out
}

func applyButton(buttons []int, state map[int]bool) map[int]bool {
	for _, b := range buttons {
		state[b] = !state[b]
	}
	return state
}

// minMoves determines min moves to reach the desired state, starting from all turned off..
func (m Machine) minMoves() int {
	// can I short it heuristically?
	start := map[int]bool{}

	// 100 is a heuristic lol
	minMoves := len(m.buttons) * 3
	var recurse func(map[int]bool, int)
	recurse = func(current map[int]bool, moves int) {
		if moves > minMoves || moves > minMoves {
			return
		}
		if stateEquals(current, m.desiredState) {
			minMoves = moves
		}
		for _, buttons := range m.buttons {
			recurse(applyButton(buttons, copyMap(current)), moves+1)
		}
	}
	recurse(start, 0)
	return minMoves

}

// stateEquals for all the values that are TRUE
func stateEquals(m1, m2 map[int]bool) bool {
	for k, v := range m1 {
		if v == true && m2[k] != v {
			return false
		}
	}
	for k, v := range m2 {
		if v == true && m1[k] != v {
			return false
		}
	}
	return true
}

func main() {
	machines := parse(readInput())
	fmt.Println(solve1(machines))
}

func solve1(machines []Machine) int {
	min := 0
	LEN := len(machines)
	for i, m := range machines {
		fmt.Printf("processed: %v of %v\n", i, LEN)
		min += m.minMoves()
	}
	return min
}

func parse(lines []string) []Machine {

	parseDesiredState := func(in string) map[int]bool {
		out := map[int]bool{}
		for idx, s := range in[1 : len(in)-1] {
			out[idx] = s == '#'
		}
		return out
	}

	parseJoltage := func(in string) []int {
		return nil
	}

	parseButtons := func(in string) []int {
		trimmed := in[1 : len(in)-1]
		out := []int{}
		for _, p := range strings.Split(trimmed, ",") {
			num, err := strconv.Atoi(p)
			if err != nil {
				panic(err)
			}
			out = append(out, num)
		}

		return out
	}

	machines := []Machine{}
	for _, line := range lines {
		m := Machine{buttons: [][]int{}}
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		m.desiredState = parseDesiredState(parts[0])
		m.joltage = parseJoltage(parts[len(parts)-1])
		for i := 1; i < len(parts)-1; i++ {
			m.buttons = append(m.buttons, parseButtons(parts[i]))
		}
		machines = append(machines, m)
	}
	return machines
}
func readInput() []string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}
