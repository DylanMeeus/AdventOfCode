package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

type point struct {
	x,y int
}

type claim struct {
	id string
	x, y, w, h int
}

func main() {
	cs := claims()
	var sum int
	for _,v := range overlaps(cs){
		if v >= 2 {
			sum++
		}
	}
	fmt.Printf("%v\n", sum)
	fmt.Printf("%v\n", notOverlapping(cs))
}

func hasOverlappingPoints(c1, c2 *claim, points map[claim][]point) bool {
	for _,p1 := range points[*c1] {
		for _,p2 := range points[*c2] {
			if p1 == p2 {
				return true
			}
		}
	}
	return false
}

func notOverlapping(cs []claim) string {
	claimPoints := make(map[claim][]point)
	for _,c := range cs {
		claimPoints[c] = pointsInClaim(c)
	}
	for _,c := range cs {
		var overlaps bool = false
		for _,o := range cs {
			if c.id != o.id {
				if hasOverlappingPoints(&c, &o, claimPoints) {
					overlaps = true
					continue
				}
			}
		}
		if !overlaps {
			fmt.Println(c.id)
		}
	}
	return ""
}

func pointsInClaim(c claim) []point {
	points := []point{}
	for column := c.x; column < c.x + c.w; column++ {
		for row := c.y; row < c.y + c.h; row++ {
			points = append(points, point{column, row})
		}
	}
	return points
}

func overlaps(cs []claim) map[point]int {
	pm := make(map[point]int, 0)
	for _,c := range cs {
		for column := c.x; column < c.x + c.w; column++ {
			for row := c.y; row < c.y + c.h; row++ {
				pm[point{column, row}]++
			}
		}
	}
	return pm
}

func claims() []claim {
	b, err := ioutil.ReadFile(`C:\Development\Go\src\AoC\2018\3\input.txt`)
	if err != nil {
		panic(err)
	}
	cs := make([]claim, 0)
	for _,line := range strings.Split(string(b), "\r\n") {
		var err error
		parts := strings.Split(line, " ")
		c := claim{}
		c.id = parts[0]
		xy := strings.Split(parts[2], ",")
		c.x, err = strconv.Atoi(xy[0])
		c.y, err = strconv.Atoi(xy[1][:len(xy[1])-1])
		wh := strings.Split(parts[3], "x")
		c.w, err = strconv.Atoi(wh[0])
		c.h, err = strconv.Atoi(wh[1])
		if err != nil {
			panic(err)
		}
		cs = append(cs, c)
	}
	fmt.Printf("%v\n", cs)
	return cs
}

