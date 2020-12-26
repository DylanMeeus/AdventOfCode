package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", solve1())
}

func solve1() int {
	value := 1
	sn := 7
	mod := 20201227

	//cardPk := 5764801
	//doorPk := 17807724

	cardPk := 15113849
	doorPk := 4206373

	pkLoop := map[int]int{}

	// how do we determine the loop size used to figure out these values?
	// determine the loop size?
	i := 0
	for pkLoop[cardPk] == 0 || pkLoop[doorPk] == 0 {
		value *= sn
		value = (value % mod)

		if value == cardPk {
			pkLoop[cardPk] = i + 1
		}
		if value == doorPk {
			pkLoop[doorPk] = i + 1
		}
		i++
	}

	// now we modify door by card?

	fmt.Printf("%v\n", pkLoop)
	sn = doorPk
	value = 1
	for i := 0; i < pkLoop[cardPk]; i++ {
		value *= sn
		value = (value % mod)
	}

	return value
}

// you have to figure out the loop size..
