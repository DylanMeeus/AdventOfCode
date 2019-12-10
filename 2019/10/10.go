package main

import (
	"io/ioutil"
	"fmt"
	"math"
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
		vis := countVisible(a, in)
		if vis > maxtroids {
			maxtroids = vis
			bestroid = a
			fmt.Printf("bestroid..%v\n", bestroid)
		}
	}
	fmt.Printf("bestroid %v\n", bestroid)
	return maxtroids
}

func countVisible(cur asteroid, other []asteroid) (count int) {
	for _, o := range other {
		if isVisible(cur, o, other) {
			count++
		}
	}
	return
}

// if multiple ones have the same slope, only count one?
func isVisible(cur, target asteroid, other []asteroid) bool {
	// make sure there is nothing between the other asteroid and ourselves
	x1, x2 := cur.x, target.x
	y1, y2 := cur.y, target.y
	if x2 == x1 {
		return traceY(cur, target, other)
	}
	if y2 == y1 {
		return true
	}
	distanceCurTarget := math.Sqrt(math.Pow(float64(x1)-float64(x2), 2) - math.Pow(float64(y1)-float64(y2), 2))
	s := (y1 - y2) / (x1 - x2)
	// b = mx - y
	startCur := (s * x1) - y1
	for _, o := range other {
		if o != cur && o != target {
			if (startCur * o.x) == o.y {
				// it's on the line.. but is it closer?
				pow1 := math.Pow(float64(x1)-float64(o.x), 2)
				pow2 := math.Pow(float64(y1)-float64(o.y), 2)
				distanceCurO := math.Sqrt(pow1 - pow2)
				if pow1-pow2 < 2 {
					distanceCurO = 1
				}
				//fmt.Printf("distance: %v pow1: pow2: %v\n", distanceCurO, pow1, pow2)
				if distanceCurO < distanceCurTarget {
					fmt.Println("false")
					return false
				}
			}
		}
	}
	return true
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
