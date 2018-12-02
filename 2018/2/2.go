package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)


func main() {
    fmt.Printf("%v\n", solve())
}

func solve() int {
    codes := codes()
    var doubles, triples int
    for _, code := range codes {
        charcount := make(map[string]int, len(code))
        for _, c := range strings.Split(code, "") {
            charcount[c]++
        }
        var hasDoubles bool
        var hasTriples bool
        for _,v := range charcount { 
            if v == 2 {
                hasDoubles = true
            }
            if v == 3 {
                hasTriples = true 
            }
        }
        if hasDoubles {
            doubles++
        }
        if hasTriples {
            triples++
        }
    }
    return doubles * triples 
}

func codes() []string {
    content,_ := ioutil.ReadFile("input.txt")
    return strings.Split(string(content), "\n")
}
