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
	school := simulate(data, 256)
	return len(school)
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
