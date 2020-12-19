package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	char     string
	subRules [][]int
}

type resolvedRule struct {
	resolved   []string
	unresolved [][]int
}

func (r rule) isCharRule() bool {
	return r.char != ""
}

type TreeNode struct {
	Char     string
	Children [][]TreeNode
}

func main() {
	fmt.Printf("%v\n", solve1())
	fmt.Printf("%v\n", solve2())
}

func solve1() int {
	rules := getRules()
	lines := getInput()

	// turn the rules into a tree?

	adj := map[int]*resolvedRule{}

	for k, v := range rules {
		if v.isCharRule() {
			adj[k] = &resolvedRule{resolved: []string{v.char}}
		} else {
			adj[k] = &resolvedRule{unresolved: v.subRules}
		}
	}

	resolve(adj, 0)

	out := 0
	for _, line := range lines {
		for _, opt := range adj[0].resolved {
			if line == opt {
				out++
			}
		}
	}

	return out
}

func solve2() int {
	rules := getRules()
	lines := getInput()

	// turn the rules into a tree?

	adj := map[int]*resolvedRule{}

	rules[8] = rule{subRules: [][]int{
		{42},
		{42, 8},
	}}

	rules[11] = rule{subRules: [][]int{
		{42, 31},
		{42, 11, 31},
	}}

	for k, v := range rules {
		if v.isCharRule() {
			adj[k] = &resolvedRule{resolved: []string{v.char}}
		} else {
			adj[k] = &resolvedRule{unresolved: v.subRules}
		}
	}

	// resolve these rules(?)
	resolve(adj, 42)
	resolve(adj, 31)

	r42 := fmt.Sprintf("(%s)", strings.Join(adj[42].resolved, "|"))
	r31 := fmt.Sprintf("(%s)", strings.Join(adj[31].resolved, "|"))

	// rule 8 is essentially 1 or more instances of rule 42
	rule8String := fmt.Sprintf("(%s)+", r42)

	makeRegexp := func(num int) *regexp.Regexp {
		// rule 11 is an equal number of 42 and 31 rules
		return regexp.MustCompile(fmt.Sprintf("^%s%s{%d}%s{%d}$", rule8String, r42, num, r31, num))
	}

	out := 0
	for _, line := range lines {
		for i := 1; i < 5; i++ { // magic numbers woo, try it X time sfor the recursive regex
			pattern := makeRegexp(i)
			if pattern.MatchString(line) {
				out++
				break
			}
		}
	}

	return out

}
func resolve(adj map[int]*resolvedRule, current int) []string {
	if len(adj[current].resolved) != 0 {
		// return a copy of resolved otherwise there's all kinds of side effect errors
		return append([]string{}, adj[current].resolved...)
	}

	// iterate through options, resolve children and append resolved paths
	// for the current entry point
	for _, option := range adj[current].unresolved {
		resolved := []string{""}
		for _, entryPoint := range option {
			nestedResolveVals := resolve(adj, entryPoint)
			var newResolved []string
			for _, nextPiece := range nestedResolveVals {
				for _, prev := range resolved {
					newResolved = append(newResolved, prev+nextPiece)
				}
			}
			resolved = newResolved
		}
		adj[current].resolved = append(adj[current].resolved, resolved...)
	}

	return adj[current].resolved
}

func getInput() []string {
	in, _ := ioutil.ReadFile("input.txt")
	return strings.Split(string(in), "\n")
}

// oh boy, this is a mess :-)
func getRules() map[int]rule {
	out := map[int]rule{}
	in, _ := ioutil.ReadFile("rules.txt")
	for _, line := range strings.Split(string(in), "\n") {
		if line == "" {
			continue
		}

		ruleParts := strings.Split(line, ":")
		ruleStr, spec := ruleParts[0], ruleParts[1]
		ruleIdx, _ := strconv.Atoi(ruleStr)
		if strings.Contains(spec, "\"") {
			spec = strings.TrimSpace(spec)
			spec = string(spec[1])
			out[ruleIdx] = rule{char: spec}
		} else {
			// parse other rules
			r := rule{}
			if strings.Contains(spec, "|") {
				// split
				for _, part := range strings.Split(spec, "|") {
					nums := []int{}
					numbers := strings.Split(part, " ")
					for _, num := range numbers {
						if num == "" {
							continue
						}
						i, _ := strconv.Atoi(num)
						nums = append(nums, i)
					}
					r.subRules = append(r.subRules, nums)
				}
			} else {
				nums := []int{}
				for _, num := range strings.Split(spec, " ") {
					if num == "" {
						continue
					}
					i, _ := strconv.Atoi(num)
					nums = append(nums, i)
				}
				r.subRules = append(r.subRules, nums)

			}
			out[ruleIdx] = r
		}

	}
	return out

}
