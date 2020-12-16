package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var (
	rulergx   = regexp.MustCompile("[0-9]*-[0-9]*")
	my_ticket = ticket{values: []int{79, 193, 53, 97, 137, 179, 131, 73, 191, 139, 197, 181, 67, 71, 211, 199, 167, 61, 59, 127}}
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
	first, _ := solve1()
	fmt.Printf("%v\n", first)
	fmt.Printf("%v\n", solve2())
}

func solve1() (int, []ticket) {
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
	valids := []ticket{}
	for _, ticket := range tickets {
		vt := true
		for _, val := range ticket.values {
			if _, ok := valid[val]; !ok {
				result += val
				vt = false
			}
		}
		if vt {
			valids = append(valids, ticket)
		}
	}

	return result, valids
}

func solve2() int {
	// first filter for valid tickets
	_, validTickets := solve1()
	validTickets = append(validTickets, my_ticket)
	rules := getRules()

	fmt.Printf("%v\n", recursiveMatch(rules, validTickets))
	return 0

	ruleFieldMap := map[string]int{}
	for _, rule := range rules {
		res := findMatchingField(rule, validTickets)
		ruleFieldMap[rule.name] = res
	}

	fmt.Printf("%v\n", ruleFieldMap)
	// we have to map rules to a corresponding value

	return 0
}

func recursiveMatch(rules []rule, tickets []ticket) []int {
	fieldOrder := []int{}
	N := len(tickets[0].values) - 1

	var backtrack func(fields []int) bool
	backtrack = func(fields []int) bool {
		if len(fields) == len(tickets[0].values) {
			fieldOrder = fields
			return true
		}

		for i := 0; i < N; i++ {
			if !contains(fields, i) {
				for _, r := range rules {
					if allMatch(r, tickets, i) {
						cp := make([]int, len(fields))
						copy(cp, fields)
						cp = append(cp, i)
						if backtrack(cp) {
							return true
						}
					}
				}
			}
		}
		return false

	}
	backtrack([]int{})
	return fieldOrder

}

func contains(is []int, i int) bool {
	for _, x := range is {
		if x == i {
			return true
		}
	}
	return false
}

func findMatchingField(r rule, others []ticket) int {
	N := len(others[0].values) - 1

	for i := 0; i < N; i++ {
		if allMatch(r, others, i) {
			return i
		}
	}
	return -1
}

func allMatch(rule rule, others []ticket, field int) bool {
	valid := map[int]bool{}
	for _, r := range rule.ranges {
		for s := r.start; s < r.end; s++ {
			valid[s] = true
		}
	}
	for _, ticket := range others {
		if _, ok := valid[ticket.values[field]]; !ok {
			return false
		}

	}
	return true
}

func getRules() []rule {
	in, _ := ioutil.ReadFile("rules.txt")

	rules := []rule{}
	for _, line := range strings.Split(string(in), "\n") {
		if line == "" {
			continue
		}
		name := strings.Split(line, ":")[0]
		r := rule{name: name}
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
