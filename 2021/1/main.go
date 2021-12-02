package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%v\n", solve())
	fmt.Printf("%v\n", solve2())
}

func getData() (out []int) {
	in, _ := ioutil.ReadFile("./input.txt")

	for _, part := range strings.Split(string(in), "\n") {
		if part == "" {
			continue
		}

		intgr, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		out = append(out, intgr)
	}
	return
}

func solve() (sum int) {
	data := getData()
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {

			sum++
		}
	}
	return sum
}

func solve2() (sum int) {
	defer func() {
		recover()
	}()

	data := getData()
	for i := 0; i < len(data); i++ {
		a := data[i] + data[i+1] + data[i+2]
		b := data[i+1] + data[i+2] + data[i+3]
		if b > a {
			sum++
		}
	}
	return sum
}
