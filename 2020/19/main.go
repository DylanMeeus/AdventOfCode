package main

import (
	"fmt"
	"io/ioutil"
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

func match(line string, r rule, rules map[int]rule) bool {
	// do we have to build a string from these rules?
	// can we expand the rules until we have a list of all allowed permutations?

	// expand until it's just chars?

	return true
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
