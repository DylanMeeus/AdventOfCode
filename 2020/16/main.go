package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var (
	rulergx = regexp.MustCompile("[0-9]*-[0-9]*")
)

type ticket struct {
	values []int
}

type vrange struct {
	start, end int
}

type rule struct {
	name   string
	ranges []vrange
}

func main() {
	fmt.Printf("%v\n", solve1())
}

func solve1() int {
	rules := getRules()
	tickets := getOthers()
	// if a field does not match _anything_ it is invalid
	// let's collapse all the rules so we don't have to keep iterating over all of them

	valid := map[int]bool{}
	for _, rule := range rules {
		for _, r := range rule.ranges {
			for s := r.start; s < r.end; s++ {
				valid[s] = true
			}
		}
	}

	result := 0
	for _, ticket := range tickets {
		for _, val := range ticket.values {
			if _, ok := valid[val]; !ok {
				result += val
			}
		}
	}

	return result
}

func getRules() []rule {
	in, _ := ioutil.ReadFile("rules.txt")

	rules := []rule{}
	for _, line := range strings.Split(string(in), "\n") {
		r := rule{}
		matches := rulergx.FindAllString(line, -1)
		for _, match := range matches {
			parts := strings.Split(match, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			r.ranges = append(r.ranges, vrange{start, end})
		}
		rules = append(rules, r)
	}

	return rules
}

func getOthers() []ticket {
	in, _ := ioutil.ReadFile("other_tickets.txt")

	tickets := []ticket{}
	for _, line := range strings.Split(string(in), "\n") {
		if line == "" {
			continue
		}
		t := ticket{}
		numbers := strings.Split(line, ",")

		for _, num := range numbers {
			i, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			t.values = append(t.values, i)
		}
		tickets = append(tickets, t)
	}

	return tickets
}
