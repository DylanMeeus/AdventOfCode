package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

const (
    prefixOn = "turn on"
    prefixOff = "turn off"
    toggle = "toggle"
)

type point struct {
    x, y int
}

type transition int

const (
    ON transition = 1
    OFF transition = 2
    TOGGLE transition = 3
)


func NewPoint(in string) point {
    parts := strings.Split(in, ",")
    a, err := strconv.Atoi(parts[0])
    if err != nil {
        panic(err)
    }
    b, err := strconv.Atoi(parts[1])
    if err != nil {
        panic(err)
    }
    return point{a,b}
}

func main() {
    grid := make([][]bool,1000)
    for i,_ := range grid {
        grid[i] = make([]bool, 1000)
    }
    b, err := ioutil.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    instructions := string(b)
    lines := strings.Split(instructions, "\n")
    for _, line := range lines[:len(lines)-1] {
        trans, start, end := parse(line)
        for row := start.x; row <= end.x; row++ {
            for col := start.y; col <= end.y; col++ {
                switch trans {
                case ON:
                    grid[row][col] = true
                case OFF:
                    grid[row][col] = false
                case TOGGLE:
                    grid[row][col] = !grid[row][col]
                }
            }
        }
    }
    fmt.Println(count(grid))
}

func count(g [][]bool) int {
    c := 0
    for _,row := range g {
        for _, col := range row {
            if col {
                c++
            }
        }
    }
    return c
}


func parse(line string) (transition, point, point) {
    state := strings.HasPrefix(line, prefixOn)
    var t transition
    if state {
        t = ON
    } else {
        t = OFF
    }
    parts := strings.Split(line, " ")
    var startS, endS string
    if strings.HasPrefix(line, toggle) {
        // parse toggle
        t = TOGGLE
        startS = parts[1]
        endS = parts[3]
    } else {
        startS = parts[2]
        endS = parts[4]
    }
    start := NewPoint(startS)
    end := NewPoint(endS)
    return t, start, end
}
