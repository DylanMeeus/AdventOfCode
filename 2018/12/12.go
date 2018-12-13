package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)


type rule struct {
    rule map[int]bool
    outcome bool
}

func main(){
    initial, rules := input()
    solve2(initial,rules)
}

type plantstate struct {
    data map[int64]bool
}

func solve2(initial map[int64]bool, rules []*rule) {
    state := initial 
    pad := func(state map[int64]bool) {
        // find the minimum
        var min, max int64 
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

    sum := func(state map[int64]bool) int64 {
        var s int64
        for k,v := range state {
            if v {
                s += int64(k)
            }
        }
        return s
    }

    pad(state)
    // now we create a new state based on this one and the rules
    for i := int64(0); i < int64(50000000000); i++ {
        newmap := make(map[int64]bool)
        for k,v := range state {
           newmap[k] = calculateState(k,v, state, rules) 
        }
        delta := sum(newmap) - sum(state)
        if delta == 75 { // we recurse here (by trial and error) find out our generation
            currentScore := sum(state)
            revisionsRemaining := 50000000000 - i
            currentScore += (75 * revisionsRemaining)
            fmt.Printf("%v\n", currentScore)
            return
        }
        state = newmap
        pad(state)
    }
}

func solve(initial map[int64]bool, rules []*rule) {
    state := initial 
    pad := func(state map[int64]bool) {
        // find the minimum
        var min, max int64 
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
    for i := 0; i < 20; i++ {
        newmap := make(map[int64]bool)
        for k,v := range state {
           newmap[k] = calculateState(k,v, state, rules) 
        }
        state = newmap
        pad(state)
    }
    var sum int64
    for k,v := range state {
        if v {
            sum += k
        }
        //fmt.Printf(" %v", newmap[i])
    }
    fmt.Printf("%v\n", sum)
}

func calculateState(index int64, alive bool, states map[int64]bool, rules []*rule) bool {
    // get all rules that might apply to current one by checking the center
    for _,r := range rules {
        match := true 
        for k,v := range r.rule {
            if states[index+int64(k)] != v {
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

func input() (state map[int64]bool, rules []*rule) {
    state = make(map[int64]bool)
    bytes,_ := ioutil.ReadFile("input.txt")
    parts := strings.Split(string(bytes), "\n")
    originalState := strings.Split(parts[0], " ")[2]
    for i,r := range originalState {
        if r == rune('#') {
            state[int64(i)] = true
        } else {
            // put it explicitely to track it easier?
            state[int64(i)] = false
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
