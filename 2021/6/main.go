package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Fish struct {
	Timer int8
}

func (f *Fish) Tick() bool {

	if f.Timer == 0 {
		f.Timer = 6
		return true
	}

	f.Timer--
	return false
}

func spawn() *Fish {
	return &Fish{Timer: 8}
}

func main() {
	fmt.Printf("%v\n", solve())
	fmt.Printf("%v\n", solve2())
}

func getData() []Fish {

	input, _ := ioutil.ReadFile("./test_input.txt")

	in := strings.Replace(string(input), "\n", "", -1)
	stringFishParts := strings.Split(in, ",")

	school := make([]Fish, len(stringFishParts))
	for i, timer := range stringFishParts {
		intTimer, err := strconv.Atoi(timer)
		if err != nil {
			panic(err)
		}
		school[i] = Fish{Timer: int8(intTimer)}
	}

	return school
}

func solve() int {
	data := getData()
	school := simulate(data, 80)
	return len(school)
}

func solve2() int {
	data := getData()
	lendata := len(data)

	sum := 0
	for i, fish := range data {
		sum += 1 + recursiveSpawn(256, int(fish.Timer))
		fmt.Printf("processed fish %v of %v\n", i, lendata)
	}

	return sum
}

func recursiveSpawn(days_left, start int) int {
	if days_left < start {
		return 0
	}

	total := 0
	for i := start; i < days_left; i += 7 {
		total += 1 + recursiveSpawn(days_left-i, 9)
	}
	return total
}

func copyMutable(school []Fish) []*Fish {
	mutableSchool := make([]*Fish, len(school))
	for i := 0; i < len(school); i++ {
		mutableSchool[i] = &school[i]
	}
	return mutableSchool
}

func simulate(initialSchool []Fish, days int) []*Fish {
	school := copyMutable(initialSchool)
	for day := 0; day < days; day++ {
		for _, fish := range school {
			if givesBirth := fish.Tick(); givesBirth {
				school = append(school, spawn())
			}
		}
	}

	return school
}
