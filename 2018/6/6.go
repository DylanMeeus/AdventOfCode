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


func main() {
    ps := points()
    l,t,r,b := boundingBox(ps)
    f := field(l,t,r,b)
    fmt.Printf("%v\n", solve(f,ps,l,t,r,b))
}

func solve(field []*point, positions []*point, left, top , right, bottom int) int {
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
        if valid(p, field, left,top,right,bottom) {
            count[p.id]++
        }
    }
    fmt.Printf("%v\n", count)
    var max int
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

func valid(p *point, field []*point, left, top, right, bottom int) bool {
    for _, fp := range field {
        if p.id == fp.id {
            if fp.y == left || fp.y == right || fp.x == top || fp.x == bottom {
                return false
            }
        }
    }
    return true
}


func distance(p1, p2 *point) int {
    return int(math.Abs(float64(p2.x - p1.x)) + math.Abs(float64(p2.y-p1.y)))
}

func field(left, top, right, bottom int) []*point {
    f := make([]*point, 0)
    for row := top; row <= bottom; row++ {
        for col := left; col <= right; col++ {
            f = append(f, &point{-1, row, col})
        }
    }
    return f
}


func boundingBox(points []*point) (left,top,right,bottom int) {
    left, top = 1000, 1000
    for _,p := range (points) {
        if p.x > bottom {
            bottom = p.x
        }
        if p.x < top {
            top = p.x
        }
        if p.y > right {
            right = p.y
        }
        if p.y < left {
            left = p.y
        }
    }
    return
}

func points() []*point{
    bytes, _ := ioutil.ReadFile("test.txt")
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
