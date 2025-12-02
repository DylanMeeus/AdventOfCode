package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

	/*
		is = []instr{
			{
				dir: "L", val: 60,
			},
			{
				dir: "R", val: 50,
			},
		}
	*/
	//fmt.Println(solve1(is))
	fmt.Println(solve2(is))
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

func solve2(is []instr) int {
	pos := 50
	c := 0
	for _, inst := range is {
		var nextPos int
		if inst.dir == "R" {
			nextPos = (pos + inst.val) % DIAL_LENGTH
			rollOver := (pos + inst.val) / DIAL_LENGTH
			fmt.Printf("roll over: %v\n", rollOver)
			c += rollOver
		} else if inst.dir == "L" {
			nextPos = (pos - inst.val)
			if nextPos == 0 {
				c++
			} else if nextPos < 0 {
				if pos != 0 {
					v := pos - inst.val
					div := math.Abs(float64(v / DIAL_LENGTH))
					if int(div) > 0 {
						fmt.Printf("roll over (L): %v\n", div)
						c += int(div) + 1
					} else {
						fmt.Printf("roll over (L): %v\n", 1)
						c++
					}
				}
				rem := nextPos % DIAL_LENGTH
				nextPos = DIAL_LENGTH + rem
			}
			// figure out how much smaller than 0
		} else {
			panic("should not happen yo")
		}

		pos = nextPos
		fmt.Println(pos)
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
