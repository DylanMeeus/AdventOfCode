package main

import (
    "fmt"
    "strings"
    "io/ioutil"
)

// nice strings contain:
// - 3 vowels
// - 1 letter that appears twice in a row
// does not contain (ab,cd,pq,xy)

var (
    vowels = map[string]struct{}{
        "a":struct{}{},
        "e":struct{}{},
        "i":struct{}{},
        "o":struct{}{},
        "u":struct{}{}}
    blacklist = []string{"ab","cd","pq","xy"}
)

func main() {
    b, err := ioutil.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    input := string(b)
    lines := strings.Split(input, "\n")
    nice := 0
    for _, word := range lines {
        parts := strings.Split(word, "")
        if _3vowels(parts) && excludes(word) && twice(parts) {
            nice++
        }
    }
    fmt.Printf("%v\n",nice)
}


func twice(xs []string) bool {
    for i := 0; i < len(xs) - 1; i++ {
        if xs[i] == xs[i+1] {
            return true
        }
    }
    return false
}

// min 3 vowels
func _3vowels(xs []string) bool {
    count := 0
    for _,x := range xs {
        if _, ok := vowels[x]; ok {
            count++
        }
    }
    return count >= 3 
}

// string does not contain (ab, cd, pq, xy)
func excludes(s string) bool {
    for _,x := range blacklist {
        if strings.Contains(s, x) {
            return false
        }
    }
    return true
}
