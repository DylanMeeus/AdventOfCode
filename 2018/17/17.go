package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

type point struct {
    x, y int
}

type worldmap map[point]string

var watermark = "~"
var sandmark = ""
var wall = "#"

func main() {
    wm := parse()
    maxy := wm.getMaxy()

    for i := 0; i < 20; i++ {
        wm.step(maxy)
    }   
    wm.stdout()
    // count the water tiles
    var sum int
    for k,v := range wm {
        if v == "~" {
            if k.y <= maxy {
                sum++
            }
        }
    }
//    fmt.Printf("%v\n", sum)
}


// will mutate :D
func (w worldmap) step(maxy int) bool {
    flows := false
    water := w.waterTiles()
    if len(water) == 0 {
        // one below the spring
        w[point{500,1}] = "|"
        return true
    }
    
    for _,t := range water {
        left := point{t.x-1, t.y}
        right := point{t.x+1, t.y}
        below := point{t.x, t.y+1}
        if w[t] == "~" {
            if w[below] == "" {
                w[t] = "|"
            }
        }

        if w[t] == "|" {
            if w[point{t.x, t.y+1}] == "" {
                w[point{t.x, t.y+1}] = "|"
            } else if w[below] == "#" {
                w[t] = "~"
            } else if w[below] == "~" {
                // check if it is bound
                var leftbound, rightbound bool
                for i := 0; i < 10000; i++ {
                    if w[point{t.x - i, t.y}] == "#" {
                        leftbound = true
                    }
                    if w[point{t.x + i, t.y}] == "#" {
                        rightbound = true
                    }
                }
                if leftbound || rightbound {
                    w[t] = "~"
                }
            }
        }


        if w[t] == "~" {
            f w[left] == "" {
                w[left] = "~"
            } 
            if w[right] == "" {
                w[right] = "~"
            }
        }

    }

    return flows
}

// check bounds
func (w worldmap) inContainer(p point) bool {
    return false
}

func  (w worldmap) waterTiles() []point {
    pnts := []point{}
    for k,v := range w {
        if v == "~" || v == "|" {
            pnts = append(pnts, k)
        }
    }
    return pnts
}

func (w worldmap) getMaxy() int {
    var maxx, minx, maxy, miny int
    miny, minx = 10000, 10000
    for p,_ := range w {
        if p.x > maxx {
            maxx = p.x
        }
        if p.y > maxy {
            maxy = p.y
        }
        if p.x < minx {
            minx = p.x
        }
        if p.y < miny {
            miny = p.y
        }
    }
    return maxy
}

// print the world map
func (w worldmap) stdout() {
    var maxx, minx, maxy, miny int
    miny, minx = 10000, 10000
    for p,_ := range w {
        if p.x > maxx {
            maxx = p.x
        }
        if p.y > maxy {
            maxy = p.y
        }
        if p.x < minx {
            minx = p.x
        }
        if p.y < miny {
            miny = p.y
        }
    }
    for row := miny - 2; row < maxy + 2; row++ {
        for col := minx - 5; col < maxx + 3; col++ {
            char := w[point{col,row}]
            if char == "" {
                char = "."
            }
            fmt.Print(char)
        }
        fmt.Println()
    }
}



func getNumber(s string) int {
    num, err := strconv.Atoi(s)
    if err != nil {
        panic(err)
    }
    return num
}

func parse() worldmap {
    world := make(map[point]string)
    world[point{500,0}] = "+"
    bytes, _ := ioutil.ReadFile("test.txt")

    toXy := func(x, y string) (xi, yi []int) {
        xi, yi = []int{}, []int{}
        if strings.Contains(x, "..") {
            // parse list
            parts := strings.Split(x, "..") 
            start := getNumber(parts[0])
            end := getNumber(parts[1])
            for i := start; i <= end; i++ {
                xi = append(xi, i)
            }
        } else {
            num := getNumber(x)
            xi = append(xi, num)
        }

        if strings.Contains(y, "..") {
            parts := strings.Split(y, "..") 
            start := getNumber(parts[0])
            end := getNumber(parts[1])
            for i := start; i <= end; i++ {
                yi = append(yi, i)
            }
        } else {
            num := getNumber(y)
            yi = append(yi, num)
        }
        return 
    }

    for _,line := range strings.Split(string(bytes), "\n") {
        if line == "" {
            continue
        }
        parts := strings.Split(line, ",")
        fst, snd := parts[0], parts[1]
        var x, y []int
        if fst[:1] == "x" {
            x,y = toXy(fst[2:], snd[2:])
        } else {
            x,y = toXy(snd[2:],fst[2:])
        }
        for _,xs := range x {
            for _,ys := range y {
                world[point{xs,ys}] = "#"
            }
        }
    }
    return world
}


