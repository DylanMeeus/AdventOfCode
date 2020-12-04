package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type (
	Passport       string
	ValidationFunc func(string) bool
)

var (
	MANDATORY_FIELDS = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	OPTIONAL_FIELDS  = []string{"cid"}

	ValidEcl = map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}

	ValidationMap = map[string]ValidationFunc{
		"byr": ValidateByr,
		"iyr": ValidateIyr,
		"eyr": ValidateEyr,
		"hgt": ValidateHgt,
		"hcl": ValidateHcl,
		"ecl": ValidateEcl,
		"pid": ValidatePid,
	}
)

func ValidateByr(s string) bool {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i >= 1920 && i <= 2002
}

func ValidateIyr(s string) bool {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i >= 2010 && i <= 2020
}

func ValidateEyr(s string) bool {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i >= 2020 && i <= 2030
}

func ValidateHgt(s string) bool {
	s = strings.TrimSpace(s)
	value, unit := s[:len(s)-2], s[len(s)-2:]
	fmt.Printf("unit: %v\n", unit)
	fmt.Printf("value: %v\n", value)
	if unit != "cm" && unit != "in" {
		return false
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("error parsing height:%v, %v", s, err))
	}
	if unit == "cm" {
		return i >= 150 && i <= 193
	}
	return i >= 59 && i <= 76
}

func ValidateHcl(s string) bool {
	ok, err := regexp.Match("#[0-9]*[a-f]*", []byte(s))
	if err != nil {
		panic(err)
	}
	return ok
}

func ValidateEcl(s string) bool {
	_, ok := ValidEcl[s]
	return ok
}

func ValidatePid(s string) bool {
	if len(s) != 9 {
		return false
	}
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}

func main() {
	//fmt.Printf("%v\n", solve1())
	fmt.Printf("%v\n", solve2())
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

func solve2() int {
	in := getInput()
	pps := getPassports(in)
	valid := 0
	for _, pass := range pps {
		// check if they contain all fields - "cid" is optional
		ok := true
		fields := Fields(pass)
		fmt.Printf("pass: %v\n", pass)
		for _, field := range MANDATORY_FIELDS {
			if field == "cid" {
				continue
			}
			if value, o := fields[field]; o {
				if !ValidationMap[field](value) {
					ok = false
				}
			} else {
				ok = false
			}
		}
		if ok {
			valid++
		}
	}
	return valid
}

// Fields returns the fields of the passport as a map
func Fields(p Passport) map[string]string {
	parts := strings.Split(string(p), " ")
	m := map[string]string{}
	for _, part := range parts {
		if part == "" {
			continue
		}
		kv := strings.Split(part, ":")
		m[kv[0]] = strings.TrimSpace(kv[1])

	}
	return m
}

func getPassports(ss []string) []Passport {
	pass := []Passport{}
	var current Passport
	for _, s := range ss {
		if s == "" {
			pass = append(pass, current)
			current = ""
		} else {
			current = Passport(strings.Join([]string{string(current), (s)}, " "))
		}
	}
	return pass
}

func getInput() []string {
	in, _ := ioutil.ReadFile("input.txt")
	return strings.Split(string(in), "\n")
}
