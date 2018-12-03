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
}

func overlaps(cs []claim) map[point]int {
	pm := make(map[point]int, 0)
	clms := make(map[string]int, 0)
	for _,c := range cs {
		clms[c.id]++
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

