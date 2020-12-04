package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Passport string

var (
	MANDATORY_FIELDS = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	OPTIONAL_FIELDS  = []string{"cid"}
)

func main() {
	fmt.Printf("%v\n", solve1())
}

func solve1() int {
	in := getInput()
	pps := getPassports(in)

	valid := 0
	for _, pass := range pps {
		// check if they contain all fields - "cid" is optional
		ok := true
		for _, field := range MANDATORY_FIELDS {
			if ok && !strings.Contains(string(pass), field) {
				ok = false
			}
		}
		if ok {
			valid++
		}
	}

	return valid
}

func getPassports(ss []string) []Passport {
	pass := []Passport{}
	var current Passport
	for _, s := range ss {
		if s == "" {
			pass = append(pass, current)
			current = ""
		} else {
			current += Passport(s)
		}
	}
	return pass
}

func getInput() []string {
	in, _ := ioutil.ReadFile("input.txt")
	return strings.Split(string(in), "\n")
}
