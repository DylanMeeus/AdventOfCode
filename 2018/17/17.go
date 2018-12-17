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
    for wm.step(maxy) {
        // pass
    }
    // count the water tiles
    var sum int
    for k,v := range wm {
        if v == "~" {
            if k.y <= maxy {
                sum++
            }
        }
    }
    fmt.Printf("%v\n", sum)
}


// will mutate :D
func (w worldmap) step(maxy int) bool {
    water := w.waterTiles()
    if len(water) == 0 {
        // one below the spring
        w[point{500,1}] = "~"
        return true
    }

    var canGoDown bool
    for _, tile := range water {
       below := point{tile.x, tile.y+1}
       //up := point{tile.x, tile.y-1}
       if w[below] == "" && tile.y <= maxy {
            w[below] = "~"
            canGoDown = true
       } 
    }
   
    var leftOrRight bool
    if !canGoDown {
        for _, tile := range water {
            if tile.y >= maxy {
                break
            }
            left := point{tile.x-1, tile.y}
            right := point{tile.x+1, tile.y}
            below := point{tile.x, tile.y+1}
            if w[below] == "#" {
                if w[left] == "" {
                    w[left] = watermark
                    leftOrRight = true
                }
                if w[right] == "" {
                    w[right] = watermark
                    leftOrRight = true
                }
            }
        }
    }

    var flows = canGoDown || leftOrRight
    if !leftOrRight && !canGoDown {
        flows = false
        // find the lowest water bucket
        lowestPoint := point{500,1}
        for _, tile := range water {
            if tile.y > lowestPoint.y {
                left := point{tile.x-1, tile.y}
                right := point{tile.x+1, tile.y}
                //below := point{tile.x, tile.y+1}
                if w[left] == "" || w[right] == "" {
                    lowestPoint = tile
                }
            }
        }
        

        left := point{lowestPoint.x-1, lowestPoint.y}
        right := point{lowestPoint.x+1, lowestPoint.y}
        var moved int 
        for {
            if lowestPoint.y >= maxy {
                break
            }
            moved = 0
            if w[left] == "" {
                flows = true
                w[left] = watermark
            }
            if w[right] == "" {
                flows = true
                w[right] = watermark
            }
            if w[point{left.x -1, left.y}] != "#" && w[point{left.x - 1 ,left.y+1}] != ""{ 
                left = point{left.x - 1, left.y}
                moved++
            }
            if w[point{right.x + 1, right.x}] != "#" && w[point{right.x + 1, right.y+1}] != "" {
                right = point{right.x + 1, right.y}
                moved++
            }
            if moved == 0 {
                return flows
            }
        }
    }
    return flows
}

func  (w worldmap) waterTiles() []point {
    pnts := []point{}
    for k,v := range w {
        if v == "~" {
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
    bytes, _ := ioutil.ReadFile("input.txt")

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


