package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "sort"
)


type rule struct {
    rule map[int]bool
    outcome bool
}

func main(){
    initial, rules := input()
    solve(initial,rules)
}


func solve(initial map[int]bool, rules []*rule) {
    state := initial 
    pad := func(state map[int]bool) {
        // find the minimum
        var min, max int 
        for k,v := range state {
            if k < min && v {
                min = k
            }
            if k > max && v {
                max = k
            }
        }
        for i := min - 5; i < max + 5;i++ {
            state[i] = state[i] && true
        }
    }
    pad(state)
    // now we create a new state based on this one and the rules
    for i := 0; i < 50000000000; i++ {
        newmap := make(map[uint]bool)
        for k,v := range state {
           newmap[k] = calculateState(k,v, state, rules) 
        }
        state = newmap
        pad(state)
    }
    var sum uint64
    for k,v := range state {
        if v {
            sum += k
        }
        //fmt.Printf(" %v", newmap[i])
    }
    fmt.Printf("%v\n", sum)
}


func printfield(state map[int]bool) {
    // first sort them
    keys := []int{}
    for k,_ := range state {
        keys = append(keys,k)
    }
    sort.Ints(keys)
    for _,k := range keys {
        fmt.Printf(" %v", func(b bool) string {
            if b {
                return "#"
            }
            return "."
        }(state[k]))
    }
    fmt.Println()
}

func calculateState(index int, alive bool, states map[int]bool, rules []*rule) bool {
    // get all rules that might apply to current one by checking the center
    for _,r := range rules {
        match := true 
        for k,v := range r.rule {
            if states[index+k] != v {
               match = false 
            }
        }
        if match == true {
            return r.outcome
        }
    }
    // else no match found, so don't change
    return false 
}

func input() (state map[int]bool, rules []*rule) {
    state = make(map[int]bool)
    bytes,_ := ioutil.ReadFile("input.txt")
    parts := strings.Split(string(bytes), "\n")
    originalState := strings.Split(parts[0], " ")[2]
    for i,r := range originalState {
        if r == rune('#') {
            state[i] = true
        } else {
            // put it explicitely to track it easier?
            state[i] = false
        }
    }

    rules = make([]*rule,0)
    for i := 1; i < len(parts); i++ {
        if parts[i] == "" {
            continue
        }
        rps := strings.Split(parts[i], " ")
        r, outcome := rps[0], rps[2]
        rulemap := make(map[int]bool,0)
        j := -2
        rl := &rule{}
        for _,rp := range r {
            if rp == rune('#') {
                rulemap[j] = true
            } else {
                rulemap[j] = false 
            }
            j++
        }
        rl.rule = rulemap
        if []rune(outcome)[0] == rune('#') {
            rl.outcome = true 
        } else {
            rl.outcome = false
        }
        rules = append(rules,rl)
    }
    return state, rules
}
