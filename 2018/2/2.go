package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)


func main() {
    fmt.Printf("%v\n", solve())
    fmt.Printf("%v\n", solve2())
}


func solve2() string {
    codes := codes()
    for _,code := range codes {
        for _,other := range codes {
            diff, common := diffs(code, other)
            if diff == 1 {
                return common
            }
        }
    }
    return ""
}

func diffs(one, two string) (diff int, common string) {
    xs := strings.Split(one, "")
    ys := strings.Split(two, "")
    if len(xs) != len(ys) {
        return 0, ""
    }
    for i,x := range xs {
        if x != ys[i] {
            diff++
        } else {
            common += x
        }
    }
    return diff, common
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
