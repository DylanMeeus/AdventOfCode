package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%v\n", solve())
}

type Cmd struct {
	direction string
	unit      int
}

func getData() []Cmd {
	in, _ := ioutil.ReadFile("./input.txt")

	lines := strings.Split(string(in), "\n")

	cmds := make([]Cmd, len(lines)-1)
	for i, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		direction, unit := parts[0], parts[1]
		unitIntgr, err := strconv.Atoi(unit)
		if err != nil {
			panic(err)
		}
		cmds[i] = Cmd{direction: direction, unit: unitIntgr}
	}
	return cmds
}

func solve() int {
	data := getData()

	x, y := 0, 0
	actions := map[string]func(i int){
		"forward": func(i int) { x += i },
		"down":    func(i int) { y += i },
		"up":      func(i int) { y -= i },
	}

	for _, cmd := range data {
		actions[cmd.direction](cmd.unit)

	}
	return x * y
}
