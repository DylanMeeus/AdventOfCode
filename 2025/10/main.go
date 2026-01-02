package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type LinearProblem struct {
	Inputs [][][]int
	Output [][]int
	N      int
}

func (l LinearProblem) solver() int {
	// minimize the multipliers?

	return 0
}

func print2d(is [][]int) {
	N := len(is)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf(" %v ", is[i][j])
		}
		fmt.Println()
	}
}

func toMatrix(input []int, N int) [][]int {
	output := make([][]int, N)
	for i := 0; i < N; i++ {
		row := make([]int, N)
		for j := 0; j < N; j++ {
			if i == j {
				row[j] = input[j]
			} else {
				row[j] = 0
			}
		}
		output[i] = row
	}
	return output
}

func contains(i int, input []int) bool {
	for _, in := range input {
		if i == in {
			return true
		}
	}
	return false
}

func toInputMatrix(input []int, N int) [][]int {
	output := make([][]int, N)
	for i := 0; i < N; i++ {
		row := make([]int, N)
		for j := 0; j < N; j++ {
			if i == j && contains(j, input) {
				row[j] = 1
			} else {
				row[j] = 0
			}
		}
		output[i] = row
	}
	return output
}

func (m Machine) toLP() LinearProblem {
	N := len(m.joltage)
	output := toMatrix(m.joltage, N)

	inputs := [][][]int{}
	for _, input := range m.buttons {
		fmt.Println(input)
		im := toInputMatrix(input, N)
		print2d(im)
		inputs = append(inputs, im)
	}

	print2d(output)

	return LinearProblem{
		Inputs: inputs,
		Output: output,
		N:      N,
	}
}

type Machine struct {
	desiredState map[int]bool // desiredState as a map of [index]bool
	buttons      [][]int
	magicButtons []map[int][]int
	joltage      []int
}

func (machine *Machine) initMagicButtons() {
	for _, buttonGroup := range machine.buttons {
		// find all the indexes this button matches and how high they go in the voltage..

		multiplier := buttonGroup[0]
		for _, button := range buttonGroup[1:] {
			j := machine.joltage[button]
			if j < multiplier {
				multiplier = j
			}
		}
		// now we create all multiples up to J?
		m := map[int][]int{}
		for i := 2; i < multiplier; i++ {
			m[i] = func(input []int) []int {
				out := make([]int, len(input))
				for idx, num := range input {
					out[idx] = num * i
				}
				return out
			}(buttonGroup)
		}

		if len(m) != 0 {
			machine.magicButtons = append(machine.magicButtons, m)
		}
	}
}

func copyMap(m map[int]bool) map[int]bool {
	out := map[int]bool{}
	for k, v := range m {
		out[k] = v
	}
	return out
}

func copySlice(in []int) []int {
	out := make([]int, len(in))
	for idx, i := range in {
		out[idx] = i
	}
	return out
}

func applyButton(buttons []int, state map[int]bool) map[int]bool {
	for _, b := range buttons {
		state[b] = !state[b]
	}
	return state
}

func applyJoltage(buttons []int, currentJoltage []int) []int {
	out := make([]int, len(currentJoltage))
	for i, j := range currentJoltage {
		out[i] = j
	}
	for _, idx := range buttons {
		out[idx] = out[idx] + 1

	}
	return out
}

func (m Machine) joltageSum() int {
	out := 0
	for _, j := range m.joltage {
		out += j
	}
	return out
}

// minmoves to reach joltage
func (m Machine) solve2() int {
	// start with all 0
	start := make([]int, len(m.joltage))

	// test magic buttons
	m.initMagicButtons()
	fmt.Println(m.joltage)
	fmt.Println(m.magicButtons)
	return 0

	// 100 is a heuristic lol
	minMoves := m.joltageSum()
	maxMoves := m.joltageSum()

	earlyExit := func(joltages []int) bool {
		for i, _ := range joltages {
			if joltages[i] > m.joltage[i] {
				return true
			}
		}
		return false
	}

	var recurse func([]int, int)
	recurse = func(current []int, moves int) {
		if moves > minMoves || moves > maxMoves {
			return
		}
		if earlyExit(current) {
			return
		}
		if joltageEquals(current, m.joltage) {
			fmt.Println("hit")
			minMoves = moves
		} else {
			//fmt.Printf("%v %v\n", current, m.joltage)
		}
		for _, buttons := range m.buttons {
			recurse(applyJoltage(buttons, copySlice(current)), moves+1)
		}
	}

	recurse(start, 0)
	return minMoves

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

func joltageEquals(j1, j2 []int) bool {

	if len(j1) != len(j2) {
		return false
	}
	for i := 0; i < len(j1); i++ {
		if j1[i] != j2[i] {
			return false
		}
	}
	return true
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
	for _, m := range machines {
		m.toLP().solver()
		return
	}
	return
	//fmt.Println(solve1(machines))
	fmt.Println(solve2(machines))
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

func solve2(machines []Machine) int {
	min := 0
	LEN := len(machines)
	for i, m := range machines {
		fmt.Printf("processed: %v of %v\n", i, LEN)
		min += m.solve2()
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

	parseJoltage := func(in string) []int {
		return parseButtons(in)
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
	b, err := ioutil.ReadFile("test_input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}
