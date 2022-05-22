package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type gateType int

const (
	AND gateType = iota
	OR
	NOT
	RSHIFT
	LSHIFT
)

type wire struct {
	name  string
	value int16
}

type gate struct {
	input        []*wire
	output       *wire
	operation    gateType
	initialValue int16 // for rshift / lshift modiifiers
}

type node struct {
	gate     *gate
	wire     *wire
	children []*node
}

// we have to construct a bunch of wires and gates together...

func main() {
	solve()
}

func solve() {
	wires, gates := parseToWireAndGate(getInput())
	printWires(wires)
	_ = gates
}

func getInput() []string {
	bts, err := ioutil.ReadFile("./test_input.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(bts), "\n")
}

func getWire(ws []*wire, name string) *wire {
	for _, w := range ws {
		if w.name == name {
			return w
		}
	}
	return nil
}

// for debugging stuff
func printWires(ws []*wire) {
	for _, w := range ws {
		fmt.Printf("%v,", w.name)
	}
	fmt.Println()
}

func createWires(lines []string) []*wire {
	wires := []*wire{}
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		isInput := len(parts) == 3
		if isInput {
			value, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			wires = append(wires, &wire{name: parts[2], value: int16(value)})
		} else {
			for _, part := range parts {
				isOperator := part == "NOT" || part == "RSHIFT" || part == "LSHIFT" || part == "AND" || part == "OR"
				isArrow := part == "->"
				isDigit := func() bool {
					_, err := strconv.Atoi(part)
					return err == nil
				}()
				isWire := !isOperator && !isArrow && !isDigit
				if isWire {
					if getWire(wires, part) == nil {
						wires = append(wires, &wire{name: part, value: 0})
					}
				}
			}
		}
	}
	return wires
}

func parseToWireAndGate(lines []string) (wires []*wire, gates []gate) {
	wires = createWires(lines)
	return
	gates = []gate{}
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		isInput := len(parts) == 3
		if isInput {
			value, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			wires = append(wires, &wire{name: parts[2], value: int16(value)})
		} else if len(parts) == 5 {
			// this is an "operator" mapping
			OP := parts[1]
			switch OP {
			case "OR":
				fallthrough
			case "AND":
				break
				/*
					nodeName := parts[1] + parts[0] + parts[2]
					graph[parts[0]] = append(graph[parts[0]], nodeName)
					graph[parts[2]] = append(graph[parts[2]], nodeName)
					graph[nodeName] = []string{parts[4]}
				*/
			case "LSHIFT":
				fallthrough
			case "RSHIFT":
				break
				/*
					nodeName := parts[1] + parts[0] + parts[2]
					graph[parts[0]] = append(graph[parts[0]], nodeName)
					graph[parts[2]] = append(graph[parts[2]], nodeName)
					graph[nodeName] = []string{parts[3]}
				*/
			}
		}
	}
	return
}

func parseInput(in []string) (graph map[string][]string, inputs map[string]int16) {
	graph = map[string][]string{}
	inputs = map[string]int16{}
	for _, line := range in {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		isInput := len(parts) == 3
		if isInput {
			value, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			inputs[parts[2]] = int16(value)
		} else if len(parts) == 5 {
			// this is an "operator" mapping
			OP := parts[1]
			switch OP {
			case "OR":
				fallthrough
			case "AND":
				nodeName := parts[1] + parts[0] + parts[2]
				graph[parts[0]] = append(graph[parts[0]], nodeName)
				graph[parts[2]] = append(graph[parts[2]], nodeName)
				graph[nodeName] = []string{parts[4]}
			case "LSHIFT":
				fallthrough
			case "RSHIFT":
				nodeName := parts[1] + parts[0] + parts[2]
				graph[parts[0]] = append(graph[parts[0]], nodeName)
				graph[parts[2]] = append(graph[parts[2]], nodeName)
				graph[nodeName] = []string{parts[3]}
			}

		}
	}

	return
}
