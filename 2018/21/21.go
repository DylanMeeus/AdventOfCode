package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
    "./opcodes"
)


type instruction struct {
    opcode string
    data [3]int
}
func main(){
    instructions := parseInstructions()
    //fmt.Printf("%v\n", solve(instructions))
    fmt.Printf("%v\n", solve2(instructions))
}

func solve2(instructions []instruction) int {
    i := 0 
    seen := make(map[int]bool)
    outer:
    for {
        //fmt.Printf("starting outer with %v\n", i)
        reg := [6]int{}
        reg[0] = i
        ip := 5
        ci := 0
        completed := 0
        for ci < len(instructions){
            instruction := instructions[ci]
            reg[ip] = ci
            lookingFor := ci == 28 
            in1 := instruction.data[0]
            in2 := instruction.data[1]
            out := instruction.data[2]
            if lookingFor {
                // check if we already contain this value
                if seen[completed] && len(seen) > 1 {
                    fmt.Printf("%v\n", completed)
                    return i 
                }
                seen[completed] = true
                i++
                goto outer 
            }
            f := opcodes.Operators[instruction.opcode]
            f(in1, in2, out, &reg)
        
            ci = reg[ip]
            completed++
            ci++
        }
    }
    return -1 // something went wrong! 
}
func solve(instructions []instruction) int {
    reg := [6]int{}
    ip := 5
    ci := 0
    reg[0] = 6132824
    for ci < len(instructions){
        instruction := instructions[ci]
        reg[ip] = ci
        lookingFor := ci == 28 
        in1 := instruction.data[0]
        in2 := instruction.data[1]
        out := instruction.data[2]
        f := opcodes.Operators[instruction.opcode]
        f(in1, in2, out, &reg)
    
        if lookingFor {
            // if the eqrr instruction is this value, it will increment the IP and exit
            // We just inspect what this value is the first  time we reach it
            // Then we can assume that if 0 starts with this value, it exists immediatly
            return reg[3]
        }
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
