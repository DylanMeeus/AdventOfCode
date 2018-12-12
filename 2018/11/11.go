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
    //lvl, cell := solve()
    //fmt.Printf("%v starting at: %v\n", lvl, cell)
    solve2()
}


func solve2() {
    // make a map of point -> value?
    cloud := pointCloud()
    type result struct {
        x,y,dial,lvl int
    }

    results := make(chan result)

    maxdial := 300 
    for dial := 2; dial < maxdial; dial++ {
        // make clusters of different sizes
        go func(dial int) {
            fmt.Println(dial)
            var x,y,max int
            for i := 0; i < 300-dial; i++ {
                for j := 0; j< 300-dial; j++ {
                    lvl := clusterPower(&cell{i,j}, cloud, dial)
                    if lvl > max {
                        max = lvl
                        x = i
                        y = j
                    }
                }
            }
            results <- result{x,y,dial,max}
        }(dial)
    }

    ress := make([]result, 298)
    for x := 2; x < maxdial; x++ {
        res := <-results
        fmt.Printf("dial %v max %v, result: %v\n", res.dial, res.lvl, res)
        ress = append(ress, res)
    }


    var bestresult result
    for _,r := range ress {
        if r.lvl > bestresult.lvl {
            bestresult = r
        }
    }
    fmt.Printf("dial %v max %v, result: %v\n", bestresult.dial, bestresult.lvl, bestresult)
}

func clusterPower(start *cell, cells map[cell]int, dial int) int {
    var pow int
    for i := 0; i < dial; i++ {
        for j := 0; j < dial; j++ {
            pow += cells[cell{start.x + i, start.y + j}]
        }
    }
    return pow
}



func pointCloud() (map[cell]int) {
    cloud := make(map[cell]int,0)
    for i := 0; i < 300; i++ {
        for j := 0; j < 300; j++ {
            c := cell{i,j}
            powlvl := power(&c)
            cloud[c] = powlvl
        }
    }
    return cloud 
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
