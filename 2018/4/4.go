package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
    "sort"
    "time"
)

type naptime struct {
    sleep, wake int
}

type stats struct {
    total int
    naptimes []naptime
}

func main() {
    guardstats := parse(timesheet())
    fmt.Printf("%v\n", solve1(guardstats))
}

func parse(input []string) map[string]*stats {
    guardstats := make(map[string]*stats)
    var current string
    var sleep, wake int
    for _,s := range input {
        if strings.Contains(s, "Guard #") {
            //start := strings.(s, strings.Index(s,"#"))
            current = strings.Split(strings.SplitAfter(s, "#")[1], " ")[0]
            if guardstats[current] == nil {
                guardstats[current] = &stats{total: 0, naptimes: []naptime{}}
            }
        }
        if strings.Contains(s, "asleep") {
            sleep = minute(s)
        }
        if strings.Contains(s, "wakes") {
            wake = minute(s)
            timeAsleep := wake - sleep
            if guardstats[current] != nil {
                guardstats[current].total += timeAsleep
                guardstats[current].naptimes = append(guardstats[current].naptimes, naptime{sleep, wake})
            }
        }
    }
    return guardstats
}

func solve1(gs map[string]*stats) int {
    // most sleepy guard
    var sleepyGuard string
    var maxSleep int
    for g,s := range gs {
        if s.total > maxSleep {
            maxSleep = s.total
            sleepyGuard = g
        }
    }
    // for all ranges count minutes asleep
    guard := gs[sleepyGuard]
    minrange := make(map[int]int)
    for _, naptime := range guard.naptimes {
        for i := naptime.sleep; i < naptime.wake; i++ {
            minrange[i]++
        }
    }
    var maxM, maxV int
    for k, v := range minrange {
        if v > maxV {
            maxV = v 
            maxM = k
        }
    }
    id, _ := strconv.Atoi(sleepyGuard)
    fmt.Printf("%v at %v\n", id, maxM)
    return id * maxM 
}

func minute(input string) int {
    min := strings.Split(input, ":")[1][:2]
    m, err := strconv.Atoi(min)
    if err != nil {
        panic(err)
    }
    return m
}


func timesheet() []string {
    bytes, _ := ioutil.ReadFile("input.txt")
    // sort them
    parts := strings.Split(string(bytes), "\n") 
    sort.Slice(parts, func(i, j int) bool {
        extractDate(parts[i])
        return extractDate(parts[i]).Before(extractDate(parts[j]))
    })
    for _,p := range parts{
        fmt.Println(p)
    }
    return parts
}

func extractDate(input string) time.Time {
    part := strings.SplitAfter(input, "]")[0]
    trimmed := strings.TrimFunc(part, func(r rune) bool {
        return r == '[' || r == ']'
    })
    t, _ := time.Parse("2006-01-02 15:04", trimmed)
    return t
}
