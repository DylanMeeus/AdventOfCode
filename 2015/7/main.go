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
	TRANSFER // not really a gate, but we have "passthrough" values here
)

var (
	gateTypeMapping = map[string]gateType{
		"AND":    AND,
		"OR":     OR,
		"NOT":    NOT,
		"RSHIFT": RSHIFT,
		"LSHIFT": LSHIFT,
		"->":     TRANSFER, // copy value from one wire to another as-is
	}

	constantWireCounter = 0
	constantWirePrefix  = "C_"
)

type wire struct {
	name  string
	value uint16
	isSet bool
}

type gate struct {
	input        []*wire
	output       *wire
	operation    gateType
	initialValue uint16 // for rshift / lshift modifiers, so this is optional
}

func main() {
	solve()
}

func solve() {
	wires, gates := parseToWireAndGate(getInput())
	// solve for the gates..

	wireA := getWire(wires, "a")

	for !wireA.isSet {

		for _, gate := range gates {
			// try to resolve it

			if gate.output.isSet || !gate.input[0].isSet || (len(gate.input) == 2 && !gate.input[1].isSet) {
				continue
			}

			switch gate.operation {
			case AND:
				gate.output.value = gate.input[0].value & gate.input[1].value
			case OR:
				gate.output.value = gate.input[0].value | gate.input[1].value
			case RSHIFT:
				gate.output.value = gate.input[0].value >> gate.input[1].value
			case LSHIFT:
				gate.output.value = gate.input[0].value << gate.input[1].value
			case NOT:
				gate.output.value = ^gate.input[0].value // XOR with FF to get bitwise NOT
			case TRANSFER:
				gate.output.value = gate.input[0].value
			}

			gate.output.isSet = true
		}
	}

	printWires(wires)
	fmt.Printf("a: %v\n", wireA.value)
}

func getInput() []string {
	bts, err := ioutil.ReadFile("./input2.txt")

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
	return &wire{name: name, value: uint16(value), isSet: true}
}

func createWires(lines []string) ([]*wire, []string) {
	wires := []*wire{}
	newLines := []string{}
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		newparts := []string{}
		for _, part := range parts {
			isOperator := part == "NOT" || part == "RSHIFT" || part == "LSHIFT" || part == "AND" || part == "OR"
			isArrow := part == "->"
			isWire := !isOperator && !isArrow && !isConstant(part)
			if isWire {
				if getWire(wires, part) == nil {
					wires = append(wires, &wire{name: part, value: 0})
				}
				newparts = append(newparts, part)
			} else if isConstant(part) {
				cw := createConstantWire(part)
				wires = append(wires, cw)
				newparts = append(newparts, cw.name)
			} else {
				newparts = append(newparts, part)
			}
		}
		newLines = append(newLines, strings.Join(newparts, " "))
	}
	return wires, newLines
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
			// transfer stuff
			wireIn1 := getWire(wires, parts[0])
			wireOut := getWire(wires, parts[2])
			gates = append(gates, &gate{input: []*wire{wireIn1}, output: wireOut, operation: TRANSFER})
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
				wireIn2 := getWire(wires, parts[2])
				wireOut := getWire(wires, parts[4])

				gates = append(gates, &gate{
					input:        []*wire{wireIn1, wireIn2},
					output:       wireOut,
					operation:    op,
					initialValue: 0,
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

func parseToWireAndGate(lines []string) ([]*wire, []*gate) {
	wires, newLines := createWires(lines)
	gates := createGates(newLines, wires)
	return wires, gates
}
