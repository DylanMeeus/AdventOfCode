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
    dat := parse()
    fmt.Printf("%v\n", solve(dat))
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

func parse() []*data {
    bytes, _ := ioutil.ReadFile("input.txt")
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
