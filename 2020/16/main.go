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
	//my_ticket = ticket{values: []int{11, 12, 13}}
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
	// divide & conquer
	// find solution for first half, second half, merge solutions
	// first filter for valid tickets
	_, validTickets := solve1()
	validTickets = append(validTickets, my_ticket)
	rules := getRules()

	return recursiveRemove(rules, validTickets)
}

func recursiveRemove(rules []rule, tickets []ticket) int {
	fieldMap := map[string][]int{}

	nameValidMap := map[string]map[int]bool{}
	for _, rule := range rules {
		valid := map[int]bool{}
		for _, r := range rule.ranges {
			for s := r.start; s <= r.end; s++ {
				valid[s] = true
			}
		}
		nameValidMap[rule.name] = valid
	}
	for _, rule := range rules {
		for i := 0; i < len(tickets[0].values); i++ {
			if allMatch(rule, tickets, i, nameValidMap) {
				fieldMap[rule.name] = append(fieldMap[rule.name], i)
			}
		}
	}

	fmt.Printf("%v\n", fieldMap)

	for {

		changed := false
		for rule, values := range fieldMap {
			if len(values) == 1 {
				// remove this value from all other rules
				for otherrule, othervalues := range fieldMap {
					if contains(othervalues, values[0]) {
						if rule != otherrule {
							fieldMap[otherrule] = filter(fieldMap[otherrule], values[0])
							changed = true
						}
					}
				}
			}
		}
		if !changed {
			break
		}
	}

	// now build the possibilities (permutations) of these values?

	// filter the ones that do not match..
	// and then find a line that is unique?

	result := 1
	for rule, value := range fieldMap {
		if strings.Contains(rule, "departure") {
			result *= my_ticket.values[value[0]]
		}
	}

	return result
}

func filter(is []int, i int) (out []int) {
	for _, x := range is {
		if i != x {
			out = append(out, x)
		}
	}
	return

}

func recursiveMatch(rules []rule, tickets []ticket) [][]int {
	fieldOrder := [][]int{}

	nameValidMap := map[string]map[int]bool{}
	for _, rule := range rules {
		valid := map[int]bool{}
		for _, r := range rule.ranges {
			for s := r.start; s <= r.end; s++ {
				valid[s] = true
			}
		}
		nameValidMap[rule.name] = valid
	}

	type memos struct {
		name  string
		field int
	}

	memo := map[memos]bool{}

	var backtrack func(fields []int, remainder []rule) bool
	backtrack = func(fields []int, remainder []rule) bool {
		if len(remainder) == 0 {
			c := make([]int, len(fields))
			copy(c, fields)
			fieldOrder = append(fieldOrder, c)
			return true
		}

		head := remainder[0]
		var tail []rule
		if len(remainder) > 1 {
			tail = remainder[1:]
		}

		for field := 0; field < len(tickets[0].values); field++ {
			if !contains(fields, field) {
				if valid, ok := memo[memos{head.name, field}]; ok && !valid {
					continue
				}
				if memo[memos{head.name, field}] || allMatch(head, tickets, field, nameValidMap) {
					memo[memos{head.name, field}] = true
					//fmt.Printf("rule: %v matches %v\n", head.name, field)
					cp := make([]int, len(fields))
					copy(cp, fields)
					cp = append(cp, field)
					if backtrack(cp, tail) {
						return true
					}
				} else {
					memo[memos{head.name, field}] = false
				}
			}

		}
		return false
	}
	backtrack([]int{}, rules)
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

func allMatch(rule rule, others []ticket, field int, valid map[string]map[int]bool) bool {
	for _, ticket := range others {
		if _, ok := valid[rule.name][ticket.values[field]]; !ok {
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
