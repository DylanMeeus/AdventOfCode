package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	w = 25
	h = 6
	//w = 3
	//h = 2
)

func main() {
	rd := readData()
	fmt.Printf("%v \n", solve1(rd))
	solve2(rd)
}

func solve2(input []int) []int {
	// divide it in layers
	cs := w * h
	layers := make([][]int, 0)
	for i := 0; i < len(input); i += cs {
		layers = append(layers, input[i:i+cs])
	}
	// collapse the layers..
	current := layers[0]
	for i := 1; i < len(layers); i++ {
		next := layers[i]
		current = merge(current, next)
	}
	printImage(current)
	return []int{}
}

func printImage(img []int) {
	fmt.Println(img)
	for i, n := range img {
		if i%w == 0 {
			fmt.Println()
		}
		if n == 1 {
			fmt.Print(n)
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func merge(l1, l2 []int) []int {
	out := make([]int, len(l1))
	for i, n := range l1 {
		if n == 2 {
			out[i] = l2[i]
		} else {
			out[i] = n
		}
	}
	return out
}

func solve1(input []int) int {
	// divide it in layers
	cs := w * h
	layers := make([][]int, 0)
	for i := 0; i < len(input); i += cs {
		layers = append(layers, input[i:i+cs])
	}
	var minzero *int
	var minlayer []int
	for _, l := range layers {
		zeros := count(l, 0)
		if minzero == nil || *minzero > zeros {
			minzero = &zeros
			minlayer = l
		}
	}
	return count(minlayer, 1) * count(minlayer, 2)
}

func count(hay []int, needle int) (out int) {
	for _, h := range hay {
		if h == needle {
			out++
		}
	}
	return
}

func readData() (out []int) {
	d, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	ss := strings.Replace(string(d), "\n", "", -1)
	for _, s := range strings.Split(ss, "") {
		i, e := strconv.Atoi(s)
		if e != nil {
			panic(e)
		}
		out = append(out, i)
	}
	return
}
