package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	in, _ := ioutil.ReadFile("input.txt")

	nums := map[int]bool{}
	parts := strings.Split(string(in), "\n")
	for _, part := range parts {
		i, err := strconv.Atoi(part)
		if err != nil {
			continue
		}
		delta := 2020 - i
		if _, ok := nums[delta]; ok {
			result := delta * i
			fmt.Printf("%v\n", result)
		}
		nums[i] = true
	}

	fmt.Println("done")
}
