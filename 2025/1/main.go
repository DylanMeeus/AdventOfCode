package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const DIAL_LENGTH = 100

type instr struct {
	dir string
	val int
}

func main() {
	lines := readLines()
	fmt.Println(len(lines))
	is := mapInstr(lines)
	fmt.Println(len(is))
	fmt.Println("-----")

	//fmt.Println(countRollOvers(99, -99))
	/*
		is = []instr{
			{
				dir: "R", val: 5,
			},
			{
				dir: "L", val: 5,
			},
		}
	*/

	fmt.Println(getPos(99, 1))
	fmt.Println(getPos(1, -1))

	fmt.Println(countRollOvers(0, -200))

	//fmt.Println(solve1(is))
	//fmt.Println(solve2(is))

	moves := MapToMoves(is)
	fmt.Println(moves)
	fmt.Println(solveAlt(moves))
}

func solveAlt(moves []int) int {
	pos := 50

	c := 0
	for _, move := range moves {
		pos += move
		if pos == -1 {
			pos = 99
		}
		if pos == 100 {
			pos = 0
		}
		if pos == 0 {
			c++
		}
	}

	return c
}

func solve1(is []instr) int {
	pos := 50
	m := map[int]int{}
	for _, inst := range is {
		if inst.dir == "R" {
			pos = (pos + inst.val) % DIAL_LENGTH
		} else if inst.dir == "L" {
			pos = (pos - inst.val)
			if pos < 0 {
				rem := pos % DIAL_LENGTH
				pos = DIAL_LENGTH + rem
			}
		} else {
			panic("should not happen yo")
		}
		fmt.Println(pos)
		m[pos]++
	}
	return m[0]
}

func getPos(pos, moves int) int {

	if moves > 0 {
		return (pos + moves) % DIAL_LENGTH
	}

	moves *= -1
	nextPos := pos - moves
	fmt.Println(nextPos)
	if nextPos < 0 {
		// how many times can we loop around back to 0?
		remainingMoves := moves - pos
		fmt.Printf("remaining moves: %v\n", remainingMoves)
		loops := remainingMoves / DIAL_LENGTH
		remainingMoves = remainingMoves - (loops * DIAL_LENGTH)
		return DIAL_LENGTH - remainingMoves
	} else {
		return nextPos
	}

	return -1337
}

func countRollOvers(pos int, moves int) int {
	if moves > 0 {
		// remMoves is how to calculate it yo

		if pos == 0 && moves < DIAL_LENGTH {
			return 0
		}

		distToL := (DIAL_LENGTH) - pos
		remMoves := moves - distToL
		if remMoves < 0 {
			// can't reach the end of the dial
			return 0
		}

		if remMoves == 0 {
			return 1
		}

		return 1 + (remMoves / DIAL_LENGTH)
	} else {

		// how often can it go negative..
		// how many positional moves left after we go to 0
		moves *= -1
		if pos == 0 && moves < DIAL_LENGTH {
			return 0
		}
		if moves < pos {
			return 0
		}
		remMoves := moves - pos
		if remMoves == 0 {
			return 1
		}
		return 1 + (remMoves / DIAL_LENGTH)
	}
	return 0
}

func MapToMoves(is []instr) []int {
	out := []int{}
	for _, i := range is {
		input := 1
		if i.dir == "L" {
			input = -1
		}
		for j := 0; j < i.val; j++ {
			out = append(out, input)
		}
	}
	return out
}

func solve2(is []instr) int {
	pos := 50
	c := 0
	for _, inst := range is {
		var nextPos int
		if inst.dir == "R" {
			nextPos = getPos(pos, inst.val)
			count := countRollOvers(pos, inst.val)
			c += count
			fmt.Printf("passed 0 %v times going from %v to %v in %v moves\n", count, pos, nextPos, inst.val)
		} else if inst.dir == "L" {
			count := countRollOvers(pos, -inst.val)
			c += count
			nextPos = getPos(pos, -inst.val)

			fmt.Printf("--passed 0 %v times going from %v to %v in %v moves\n", count, pos, nextPos, inst.val)
		} else {
			panic("should not happen yo")
		}

		pos = nextPos
	}
	return c
}

func mapInstr(ss []string) []instr {
	is := make([]instr, len(ss))
	for i, s := range ss {
		d := string(s[0])
		val, err := strconv.Atoi(s[1:])
		if err != nil {
			panic(err)
		}
		is[i] = instr{d, val}
	}
	return is
}

func filter(ss []string) []string {
	out := []string{}
	for _, s := range ss {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func readLines() []string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")
	return filter(lines)
}
