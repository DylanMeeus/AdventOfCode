package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
    "./opcodes"
)


var _ = opcodes.Functions

type instruction struct {
    opcode string
    data [3]int
}
func main(){
    instructions := parseInstructions()
    fmt.Printf("%v\n", solve(instructions))
    fmt.Printf("%v\n", solve2(instructions))
}

func solve2(instructions []instruction) int {
    reg := [6]int{}
    reg[0] = 1
    ip := 4
    ci := 0
    for ci < len(instructions){
        instruction := instructions[ci]
        reg[ip] = ci
        in1 := instruction.data[0]
        in2 := instruction.data[1]
        out := instruction.data[2]
        f := opcodes.Operators[instruction.opcode]
        //fmt.Printf("lookup %v: %v\n", instruction.opcode, f)
        f(in1, in2, out, &reg)
        // fetch next instruction
        ci = reg[ip]
        ci++
    }
    return reg[0]
}

func solve(instructions []instruction) int {
    reg := [6]int{}
    ip := 4
    ci := 0
    for ci < len(instructions){
        instruction := instructions[ci]
        reg[ip] = ci
        in1 := instruction.data[0]
        in2 := instruction.data[1]
        out := instruction.data[2]
        f := opcodes.Operators[instruction.opcode]
        f(in1, in2, out, &reg)
        // fetch next instruction
        ci = reg[ip]
        ci++
    }
    return reg[0]
}


func parseInstructions() []instruction {
    bytes,_ := ioutil.ReadFile("input.txt")
    ops := make([]instruction,0)
    for _, line := range strings.Split(string(bytes), "\n") {
        if line == "" {
            continue
        }   
        inst := instruction{strings.Split(line, " ")[0], Stoa(line)}
        ops = append(ops, inst)
    }
    return ops
        
}

func Stoa(in string) [3]int {
    // replace possible [], with ''
    numbers := [3]int{}
    var index int
    nums := strings.Split(in, " ")
    for _,n := range nums[1:] {
        if n == "" {
            continue
        }
        i, err := strconv.Atoi(n)
        if err != nil {
            panic(err)
        }
        numbers[index] = i
        index++
    }
    return numbers 
}
