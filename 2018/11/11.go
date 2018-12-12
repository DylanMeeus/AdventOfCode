package main

import (
    "fmt"
    "math"
)

//const serial = 1718
const serial = 1718 
type cell struct {
    x, y int
}

func main() {
    lvl, cell := solve()
    fmt.Printf("%v starting at: %v\n", lvl, cell)
}

func solve() (int, *cell) {
    
    cells := make([][]*cell, 0)
    powers := make(map[cell]int)
    for i := 0; i < 297; i++ {
        for j := 0; j < 297; j++ {
            clust := cluster(&cell{i,j})
            cells = append(cells, clust)
            for _, c := range clust {
                if powers[(*c)] == 0 {
                    powers[(*c)] = power(c)
                }
            }
        }
    }
    
    // now find the cluster with the highest value
    var max int
    var startCell *cell
    for _,clust := range cells {
        var powerlevel int
        for _, cell := range clust {
            powerlevel += powers[(*cell)]
        }
        if powerlevel > max {
            max = powerlevel
            startCell = clust[0]
        }
    }
    return max, startCell
}

// find the cluster where cell is the starting point
func cluster(c *cell) []*cell {
    cells := make([]*cell, 0)
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            nc := &cell{c.x + i, c.y + j}
            cells = append(cells, nc)
        }
    }
    return cells
}

func power(c *cell) int {
    rackID := c.x + 10
    powerlevel := rackID * c.y
    withSerial := powerlevel + serial
    multi := withSerial * rackID
    // find the 100th digit or 0
    if multi > 100 {
        //digits := int(math.Floor(math.Log10(math.Abs(float64(multi))))) + 1
        //fmt.Println(digits)
        multi %= int(math.Pow(10, float64(3)))
        //fmt.Printf("%v\n", multi)
        multi /= 100
        //fmt.Printf("%v\n", multi)
    } else {
        multi = 0
    }
    return multi - 5
}
