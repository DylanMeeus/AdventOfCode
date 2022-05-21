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

func getInput() string {
	b, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", string(b))
	return string(b)
}
