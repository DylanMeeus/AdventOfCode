package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
    "math"
)

type point struct{
    id, x, y int
}

type boundary struct {
    left,right,top,bottom int
}


func main() {
    ps := points()
    bounds := boundingBox(ps)
    f := field(bounds)
    fmt.Printf("%v\n", solve(f,ps, bounds))
    fmt.Printf("%v\n", solve2(ps, bounds))
}

func solve2(positions []*point, bounds *boundary) int {
    // we grow outside of our bounding box, because we know the max distance from it is 10.000
    // we can use maxdistance (h) from top instead of bottom, due to connecting to all points
    // and actually we can use the diff from top-bottom, but this is bruteforce anyway :)
    left, right, top, bottom := bounds.left, bounds.right, bounds.top, bounds.bottom
    maxdistance := 10000 
    cpos := len(positions)
    left -= (maxdistance / cpos)
    top -= maxdistance / cpos 
    right += maxdistance / cpos
    bottom += maxdistance / cpos
    var size int
    for i := left; i <= right; i++{
        for j := top; j <= bottom; j++ {
            var total int
            for _, cor := range positions {
                total += distance(&point{-1,j,i}, cor)
                if total >= maxdistance {
                   break  
                }
            }
            if total < maxdistance {
                size++
            }
        }
    }
    return size
}


func solve(field []*point, positions []*point, bounds *boundary) (max int) {
    // foreach point in the field, find the closest coordinate
    for _,point := range field {
        dist := 1000
        for _,cor := range positions {
            d := distance(point, cor)
            if d < dist {
                dist = d
                point.id = cor.id
            } else if d == dist {
                point.id = -1
            }
        }
    }

    count := make(map[int]int, 0)
    for _,p := range field {
        // make sure the p is valid
        if valid(p, field, bounds) {
            count[p.id]++
        }
    }
    fmt.Printf("%v\n", count)
    for k, v := range count {
        if k == -1 {
            continue
        }
        if v > max {
            max = v
        }
    }
    return max
}

func valid(p *point, field []*point, bounds *boundary) bool {
    for _, fp := range field {
        if p.id == fp.id {
            if fp.y == bounds.left || fp.y == bounds.right || fp.x == bounds.top || fp.x == bounds.bottom {
                return false
            }
        }
    }
    return true
}


func distance(p1, p2 *point) int {
    return int(math.Abs(float64(p2.x - p1.x)) + math.Abs(float64(p2.y-p1.y)))
}

func field(bounds *boundary) []*point {
    f := make([]*point, 0)
    for row := bounds.top; row <= bounds.bottom; row++ {
        for col := bounds.left; col <= bounds.right; col++ {
            f = append(f, &point{-1, row, col})
        }
    }
    return f
}


func boundingBox(points []*point) *boundary {
    bounds := &boundary{}
    bounds.left, bounds.top = 1000, 1000
    for _,p := range (points) {
        if p.x > bounds.bottom {
            bounds.bottom = p.x
        }
        if p.x < bounds.top {
            bounds.top = p.x
        }
        if p.y > bounds.right {
            bounds.right = p.y
        }
        if p.y < bounds.left {
            bounds.left = p.y
        }
    }
    return bounds
}

func points() []*point{
    bytes, _ := ioutil.ReadFile("input.txt")
    ps := make([]*point, 0)
    for i, part := range strings.Split(string(bytes), "\n") {
        if part == "" {
            continue
        }
        xy := strings.Split(part, ",")
        x,_ := strconv.Atoi(xy[0])
        y,_ := strconv.Atoi(strings.TrimSpace(xy[1]))
        ps = append(ps, &point{i,x,y})
    }
    return ps 
}
