package main

import (
    "fmt"
    "strings"
    "io/ioutil"
)

type point struct {
    x, y int
}

// I could make this a 'game' and put the max point in here.
// But laziness won against writing something clean
type world map[point]string

const lumber, tree, empty = "#", "|", "."

func main() {
    fmt.Printf("%v\n", solve())
    fmt.Printf("%v\n", solve2())
}

func solve2() int {
    wm, p := parse()
    // find repetition in this set
    for i := 0; i <= 2000; i++ {
        wm = step(wm, p)
        fmt.Printf("%v counts %v\n", i, count(wm, p))
    }
    return count(wm, p)
}

func count(wm world, p *point) int {
    var trees, lumbers int
    for row := 0; row < p.y; row++ {
        for col := 0; col < p.x; col++ {
            currentp := point{col,row}
            if wm[currentp] == tree {
                trees++
            }
            if wm[currentp] == lumber {
                lumbers++
            }
        }
    }

    return trees * lumbers
}

func solve() int {
    wm, p := parse()
    for i := 0; i <= 10; i++ {
       wm = step(wm, p)
    }
    return count(wm, p)
}

func step(w world, max *point) (world) {
    // rules:
    // open -> trees if 3+ are trees
    // tree -> lumberyard if 3+ are lumberyards
    // lumberyard -> lumberyard if next to 1+ trees, 1+ lumberyard
    neighbours := func(wm world, current point, lookup string) (count int) {
        for row := current.y - 1; row <= current.y + 1; row++ {
            for col := current.x - 1; col <= current.x + 1; col++ {
                check := point{col,row}
                if check != current && wm[check] == lookup {
                    count++
                }
            }
        }
        return
    }


    huxley := make(world)
    for row := 0; row < max.y; row++ {
        for col := 0; col < max.x; col++ {
            currentp := point{row,col}
            trees := neighbours(w, currentp, tree) 
            lumbers := neighbours(w, currentp, lumber)
            if w[currentp] == empty {
               if trees >= 3 {
                   huxley[currentp] = tree
               } else {
                   huxley[currentp] = w[currentp]
               }
            } else if w[currentp] == tree {
                if lumbers >= 3 {
                    huxley[currentp] = lumber
                } else {
                    huxley[currentp] = w[currentp]
                }
            } else if w[currentp] == lumber {
                if trees >= 1 && lumbers >= 1 {
                    huxley[currentp] = lumber // stay lumber 
                } else {
                    huxley[currentp] = empty
                }
            }
        }
    }
    return huxley
}

func (w world) stdout(max *point) {
    for row := 0; row < max.y; row++ {
        for col := 0; col < max.x; col++ {
            fmt.Printf("%v", w[point{col,row}])
        }
        fmt.Println()
    }
}


func parse() (world, *point) {
    bytes,_ := ioutil.ReadFile("input.txt")

    world := make(map[point]string)
    lines := strings.Split(string(bytes), "\n")
    var maxy, maxx int
    maxy = len(lines)
    for y, line := range lines {
        if line == "" {
            continue
        }
        chars := strings.Split(line, "")
        maxx = len(chars)
        for x,char := range chars {
            world[point{x,y}] = char
        }
    }
    max := point{maxx, maxy}
    return world, &max
}
