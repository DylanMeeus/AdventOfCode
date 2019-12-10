package main

import (
	"io/ioutil"
	"fmt"
	"strings"
)

type space [][]int

type asteroid struct {
	x, y int
}

func (a asteroid) String() string {
	return fmt.Sprintf("{x:%v, y:%v}\n", a.x, a.y)
}

func main() {
	fmt.Printf("maxtroids: %v\n", solve1(readData()))
}

func solve1(in []asteroid) int {
	// find the best asteroid in the world.
	var maxtroids int
	var bestroid asteroid
	for _, a := range in {
		vis := uniqSlopes(a, in)
		if vis > maxtroids {
			maxtroids = vis
			bestroid = a
			fmt.Printf("bestroid..%v\n", bestroid)
		}
	}
	fmt.Printf("bestroid %v\n", bestroid)
	return maxtroids
}

// if multiple ones have the same slope, only count one?
func uniqSlopes(cur asteroid, others []asteroid) int {
	// make sure there is nothing between the other asteroid and ourselves
	// float64?
	slopes := map[string]struct{}{}
	var up bool
	var down bool
	for _, target := range others {
		if cur == target {
			continue
		}
		x1, x2 := cur.x, target.x
		y1, y2 := cur.y, target.y
		if x1 == x2 {
			if y1 > y2 {
				up = true
			}
			if y1 < y2 {
				down = true
			}
			continue
		}
		s := (float64(y1) - float64(y2)) / (float64(x1) - float64(x2))
		key := "R"
		if x1 < x2 {
			key="L"
		}
		slopes[key+fmt.Sprintf("%.3f",s)] = struct{}{}
	}
	fmt.Println(slopes)
	count := len(slopes)
	if up {
		count++
	}
	if down {
		count++
	}
	return count
}

func traceY(cur, target asteroid, other []asteroid) bool {
	// trace path to target
	if cur.y < target.y {
		// down
		for _, o := range other {
			if o.x == cur.x && o.y > cur.y && o.y < target.y {
				return false
			}
		}
	}
	if cur.y > target.y {
		// up
		for _, o := range other {
			if o.x == cur.x && o.y > target.y && o.y < cur.y {
				return false
			}
		}
	}
	return true
}

func readData() (out []asteroid) {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	rows := strings.Split(string(bs), "\n")
	for i, r := range rows {
		cols := strings.Split(r, "")
		for k, c := range cols {
			if c == "#" {
				fmt.Printf("#")
				out = append(out, asteroid{x: k, y: i})
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
	return
}
