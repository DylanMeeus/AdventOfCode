package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
    "./opcodes"
)

type data struct {
    before, operation, after [4]int
}

func main(){
    instructions := parseInstructions()
    fmt.Printf("%v\n", solve(instructions))
}


func solve(instructions [][6]int) {
    ip := 4
}


func parseInstructions() [][6]int {
    bytes,_ := ioutil.ReadFile("program.txt")
    ops := make([][6]int,0)
    for _, line := range strings.Split(string(bytes), "\n") {
        if line == "" {
            continue
        }   
        ops = append(ops, Stoa(line))
    }
    return ops
        
}

func Stoa(in string) [6]int {
    // replace possible [], with ''
    sane := strings.Map(func(r rune) rune {
        if r == rune('[') || r == rune(']') || r == rune(',') {
            return -1
        }
        return r
    }, in)
    numbers := [6]int{}
    var index int
    nums := strings.Split(sane, " ")
    for _,n := range nums {
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
