package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type instruction struct {
	direction string
	value     int
}

func main() {
	fmt.Printf("%v\n", solve1())
	fmt.Printf("%v\n", solve2())
}

func solve1() float64 {
	instr := getInput()

	var x, y int

	walk := map[int]func(int){
		0: func(v int) { y -= v },
		1: func(v int) { x += v },
		2: func(v int) { y += v },
		3: func(v int) { x -= v },
	}

	orientation := 0
	for _, in := range instr {
		if in.direction == "R" {
			orientation = (orientation + 1) % 4
		} else {
			orientation--
			if orientation < 0 {
				orientation = 3
			}
		}
		// now walk this way :-)
		walk[orientation](in.value)
	}

	return math.Abs(float64(x)) + math.Abs(float64(y))
}

func solve2() float64 {
	instr := getInput()

	var x, y int
	places := map[point]bool{}

	fwalk := func(o, v int) {
		cache := func() {
			p := point{x, y}
			if v, ok := places[p]; ok && v {
				// this is the right answer lol
				panic(fmt.Sprintf("%v\n", math.Abs(float64(x))+math.Abs(float64(y))))
			}
			places[p] = true
		}

		switch o {
		case 0:
			end := y - v
			for ; y != end; y-- {
				cache()
			}
		case 1:
			end := x + v
			for ; x != end; x++ {
				cache()
			}
		case 2:
			end := y + v
			for ; y != end; y++ {
				cache()
			}
		case 3:
			end := x - v
			for ; x != end; x-- {
				cache()
			}

		}

	}

	orientation := 0
	for _, in := range instr {
		if in.direction == "R" {
			orientation = (orientation + 1) % 4
		} else {
			orientation--
			if orientation < 0 {
				orientation = 3
			}
		}
		fwalk(orientation, in.value)
	}

	return math.Abs(float64(x)) + math.Abs(float64(y))

}

func getInput() []instruction {
	in, _ := ioutil.ReadFile("input.txt")

	parts := strings.Split(string(in), ",")

	out := []instruction{}
	for _, part := range parts {
		part = strings.TrimSpace(part)
		dir := string(part[0])
		val, err := strconv.Atoi(part[1:])
		if err != nil {
			panic(err)
		}
		out = append(out, instruction{dir, val})
	}

	return out
}
