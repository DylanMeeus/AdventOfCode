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
