package main

import (
	"fmt"
	"io/ioutil"
)

type point struct {
	x, y int
}

func main() {
	fmt.Printf("%v\n", solve())
	fmt.Printf("%v\n", solve2())
}

func solve() int {
	in := getInput()

	var x, y int

	funcMap := map[string]func(){
		"^": func() { y-- },
		"v": func() { y++ },
		">": func() { x++ },
		"<": func() { x-- },
	}

	m := map[point]int{point{x, y}: 1}
	for _, c := range in {
		if f, ok := funcMap[string(c)]; ok {
			f()
			m[point{x, y}]++
		}
	}

	return len(m)
}

func solve2() int {
	in := getInput()

	santa := &point{0, 0}
	robot := &point{0, 0}

	funcMap := map[string]func(p *point){
		"^": func(p *point) { p.y-- },
		"v": func(p *point) { p.y++ },
		">": func(p *point) { p.x++ },
		"<": func(p *point) { p.x-- },
	}

	m := map[point]int{point{santa.x, santa.y}: 1}
	for i, c := range in {
		if f, ok := funcMap[string(c)]; ok {
			if i%2 == 0 {
				f(santa)
			} else {
				f(robot)
			}
			m[point{santa.x, santa.y}]++
			m[point{robot.x, robot.y}]++
		}
	}

	return len(m)
}

func getInput() string {
	b, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	return string(b)
}
