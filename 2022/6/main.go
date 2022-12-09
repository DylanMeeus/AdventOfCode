package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println(solve1())
	fmt.Println(solve2())
}

func solve2() int {
	data := getData()
	for i := 13; i < len(data); i++ {
		m := map[byte]bool{}
		for j := i - 13; j <= i; j++ {
			m[data[j]] = true
		}
		if len(m) == 14 {
			fmt.Println(m)
			return i + 1
		}
	}
	return -1
}

func solve1() int {
	data := getData()
	for i := 3; i < len(data); i++ {
		m := map[byte]bool{}
		for j := i - 3; j <= i; j++ {
			m[data[j]] = true
		}
		if len(m) == 4 {
			fmt.Println(m)
			return i + 1
		}
	}
	return -1
}

func getData() string {
	f, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	return string(f)
}
