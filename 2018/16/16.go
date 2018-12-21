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
    dat := parse("input.txt")
    fmt.Printf("%v\n", solve(dat))
    fmt.Printf("%v\n", solve2(dat))
}



func solve2(dat []*data) int {
    possibleMatches := make(map[int][]*opcodes.Op)
    for i := 0; i < 16; i++ {
        possibleMatches[i] = []*opcodes.Op{}
    }
    for _,d := range dat {
        code := d.operation[0]
        in1 := d.operation[1]
        in2 := d.operation[2]
        out := d.operation[3]
        for i := 0; i < len(opcodes.Functions); i++ {
            f := opcodes.Functions[i]
            cop := [4]int{d.before[0], d.before[1], d.before[2], d.before[3]}
            f(in1,in2,out,&cop)
            if cop == d.after {
                var contains bool
                for _,v := range possibleMatches[code] {
                    if v == &opcodes.Functions[i] {
                        contains = true 
                    }
                }
                if !contains{
                    possibleMatches[code] = append(possibleMatches[code], &opcodes.Functions[i])
                }
            }
        }
    }
    mappedOpcodes := filter(possibleMatches)
    instructions := parseInstructions()
    var reg [4]int
    for _, instruction := range instructions {
        code := instruction[0]
        in1 := instruction[1]
        in2 := instruction[2]
        out := instruction[3]
        f := *mappedOpcodes[code]
        f(in1, in2, out, &reg)
    }
    fmt.Printf("%v\n", reg)

    return len(mappedOpcodes) 
}

func filter(possible map[int][]*opcodes.Op) (map[int]*opcodes.Op){
    certain := make(map[int]*opcodes.Op)

    for len(certain) < 16 {
        for k,v := range possible {
            if len(v) == 1 {
                certain[k] = v[0]
                // remove this one from all others
                for other,vother := range possible {
                    if other != k {
                        options := []*opcodes.Op{}
                        for _,option := range vother {
                            if option != v[0] {
                                options = append(options, option)
                            }
                        }
                        possible[other] = options
                    }
                }
            }
        }
    }
    fmt.Printf("%v\n", certain)
    return certain
}

func solve(dat []*data) int {
    var count int
    for _,d := range dat {
        in1 := d.operation[1]
        in2 := d.operation[2]
        out := d.operation[3]
        var opcount int
        for _,f := range opcodes.Functions {
            cop := [4]int{d.before[0], d.before[1], d.before[2], d.before[3]}
            f(in1,in2,out,&cop)
            if cop == d.after {
                opcount++
            }
        }
        if opcount >= 3 {
            count++
        }
    }
    return count
}

func parseInstructions() [][4]int {
    bytes,_ := ioutil.ReadFile("program.txt")
    ops := make([][4]int,0)
    for _, line := range strings.Split(string(bytes), "\n") {
        if line == "" {
            continue
        }   
        ops = append(ops, Stoa(line))
    }
    return ops
        
}

func parse(file string) []*data {
    bytes, _ := ioutil.ReadFile(file)
    parts := strings.Split(string(bytes), "\n") 
    var i int
    exs := make([]*data,0)
    for i < len(parts) {
        before := parts[i]
        i++
        operation := parts[i]
        i++
        after := parts[i]
        i += 2
        be := Stoa(before[7:])
        op := Stoa(operation)
        af := Stoa(after[6:])
        exs = append(exs, &data{be,op,af})
    }
    return exs
}


func Stoa(in string) [4]int {
    // replace possible [], with ''
    sane := strings.Map(func(r rune) rune {
        if r == rune('[') || r == rune(']') || r == rune(',') {
            return -1
        }
        return r
    }, in)
    numbers := [4]int{}
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
