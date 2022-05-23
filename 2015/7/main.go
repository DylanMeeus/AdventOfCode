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
	NOOP // not really a gate, but we have "passthrough" values here
)

var (
	gateTypeMapping = map[string]gateType{
		"AND":    AND,
		"OR":     OR,
		"NOT":    NOT,
		"RSHIFT": RSHIFT,
		"LSHIFT": LSHIFT,
		"->":     NOOP,
	}

	constantWireCounter = 0
	constantWirePrefix  = "C_"
)

type valueProvider interface {
	getValue() uint16
}

// signalEmitter always emits a constant signal...
type signalEmitter struct {
	signal uint16
}

type wire struct {
	name  string
	value uint16
}

type gate struct {
	input        []*wire
	output       *wire
	operation    gateType
	initialValue uint16 // for rshift / lshift modifiers, so this is optional
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
	_ = wires
	printWires(wires)
	//printGates(gates)

	// solve for the gates..

	for _, gate := range gates {
		// try to resolve it
		fmt.Printf("%v %v\n", *gate.input[1], *gate.output)
		if gate.output.value == 0 {
			switch gate.operation {
			case AND:
				gate.output.value = gate.input[0].value & gate.input[1].value
			case OR:
				gate.output.value = gate.input[0].value | gate.input[1].value
			case RSHIFT:
				gate.output.value = gate.input[0].value >> gate.initialValue
			case LSHIFT:
				gate.output.value = gate.input[0].value << gate.initialValue
			case NOT:
				gate.output.value = ^gate.input[0].value // XOR with FF to get bitwise NOT
			case NOOP:
				gate.output.value = gate.input[0].value
			}
		}
	}

	printWires(wires)
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
		fmt.Printf("(%v %v) ", w.name, w.value)
	}
	fmt.Println()
}

func printGates(gs []*gate) {
	for _, g := range gs {
		fmt.Printf("BEGIN\n")
		fmt.Print("(")
		fmt.Printf("input wires: ")
		printWires(g.input)
		fmt.Printf("output wires: ")
		printWires([]*wire{g.output})
		fmt.Printf("operator: %v\n", g.operation)
		fmt.Printf("init: %v\n", g.initialValue)
		fmt.Printf("END \n\n")
	}
}

// hacky way to determine if input represents a constant..
func isConstant(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

func createConstantWire(s string) *wire {
	value, err := strconv.Atoi(s)
	panicIfErr(err)

	name := constantWirePrefix + strconv.Itoa(constantWireCounter)
	constantWireCounter++
	return &wire{name: name, value: uint16(value)}
}

func createWires(lines []string) []*wire {
	wires := []*wire{}
	newLines := []string{}
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		isInput := len(parts) == 3
		if isInput {
			wireIn := parts[0]
			if isConstant(wireIn) {
				wires = append(wires, createConstantWire(wireIn))
			} else {
				if w := getWire(wires, wireIn); w == nil {
					wires = append(wires, &wire{name: wireIn, value: 0})
				}
			}
			wireOut := parts[2]
			if isConstant(wireOut) {
				wires = append(wires, createConstantWire(wireOut))
			} else {
				if w := getWire(wires, wireOut); w == nil {
					wires = append(wires, &wire{name: wireOut, value: 0})
				}
			}

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
				} else if isConstant(part) {
					wires = append(wires, createConstantWire(part))
				}
			}
		}
	}
	return wires
}

// createGates constructs the actual logic gates. This expects an exhaustive list of all wires to be
// passed in.
func createGates(lines []string, wires []*wire) []*gate {
	gates := []*gate{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) == 3 {
			continue
		} else if len(parts) == 5 {
			// this is an "operator" mapping
			OP := parts[1]
			switch OP {
			case "OR":
				fallthrough
			case "AND":
				op := gateTypeMapping[OP]
				wireIn1 := getWire(wires, parts[0])
				if wireIn1 == nil {
					panic(fmt.Sprintf("no wire found for %v\nline:%v\n", parts[0], line))
				}
				wireIn2 := getWire(wires, parts[2])
				wireOut := getWire(wires, parts[4])
				gates = append(gates, &gate{
					input:        []*wire{wireIn1, wireIn2},
					output:       wireOut,
					operation:    op,
					initialValue: 0,
				})
			case "LSHIFT":
				fallthrough
			case "RSHIFT":
				op := gateTypeMapping[OP]
				wireIn1 := getWire(wires, parts[0])
				wireOut := getWire(wires, parts[4])
				initialValue, err := strconv.Atoi(parts[2])
				panicIfErr(err)
				gates = append(gates, &gate{
					input:        []*wire{wireIn1},
					output:       wireOut,
					operation:    op,
					initialValue: uint16(initialValue),
				})
			}
		} else {
			op := gateTypeMapping[parts[0]]
			wireIn1 := getWire(wires, parts[1])
			wireOut := getWire(wires, parts[3])
			gates = append(gates, &gate{
				input:        []*wire{wireIn1},
				output:       wireOut,
				operation:    op,
				initialValue: 0,
			})

		}
	}
	return gates
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func parseToWireAndGate(lines []string) (wires []*wire, gates []*gate) {
	wires = createWires(lines)
	gates = createGates(lines, wires)
	return
}

func parseInput(in []string) (graph map[string][]string, inputs map[string]uint16) {
	graph = map[string][]string{}
	inputs = map[string]uint16{}
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
			inputs[parts[2]] = uint16(value)
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
