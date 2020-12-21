package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type food struct {
	ingredients []string
	allergens   []string
}

func main() {
	fmt.Printf("%v\n", solve1())
	fmt.Printf("%v\n", solve2())
}

func solve2() string {
	foods := getInput()

	ingredientAllergen := map[string][]string{}
	allIngs := map[string]bool{}
	for _, food := range foods {
		for _, ing := range food.ingredients {
			allIngs[ing] = true
			for _, alg := range food.allergens {
				ingredientAllergen[ing] = append(ingredientAllergen[ing], alg)
			}
		}
	}

	m := map[string]map[string]bool{}
	for ingredient, possibleAllergens := range ingredientAllergen {
		for _, allergen := range possibleAllergens {
			isMatch := true
			for _, food := range foods {
				// this allergen can not appear in a line where the food does not appear
				if contains(allergen, food.allergens) && !contains(ingredient, food.ingredients) {
					isMatch = false
				}
			}
			if isMatch {
				if m[ingredient] == nil {
					m[ingredient] = map[string]bool{}
				}
				m[ingredient][allergen] = true
			}
		}
	}

	// we have to make sure they are unique..
	for k, v := range m {
		fmt.Printf("%v contains:\n", k)
		for x, _ := range v {
			fmt.Printf("\t%v\n", x)
		}
	}

	fmt.Printf("%v\n", m)

	keys := []string{}

	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return strings.Join(keys, ",")
}

func solve1() int {
	foods := getInput()

	// bruteforce: for each allergen, check if there is an ingredient that appears in every line
	// with this allergen?

	// for each food, find the list of possible allergens?
	// then for each found, count the times this allergens appears?

	/*
		mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
		trh fvjkl sbzzf mxmxvkd (contains dairy)
		sqjhc fvjkl (contains soy)
		sqjhc mxmxvkd sbzzf (contains fish)

		-> kfcds can not be fish because fish needs to be in [sgjhc, mxmvkd, sbzzf]
	*/

	// if an ingredient is missing from a line that contains the allergen, it can not be that
	// allergen?

	// for an ingredient to match an allergen, the allergen can not appear in a line where the
	// ingredient does not appear

	ingredientAllergen := map[string][]string{}

	allIngs := map[string]bool{}
	for _, food := range foods {
		for _, ing := range food.ingredients {
			allIngs[ing] = true
			for _, alg := range food.allergens {
				ingredientAllergen[ing] = append(ingredientAllergen[ing], alg)
			}
		}
	}

	out := 0
	m := map[string]string{}
	for ingredient, possibleAllergens := range ingredientAllergen {
		for _, allergen := range possibleAllergens {
			isMatch := true
			for _, food := range foods {
				// this allergen can not appear in a line where the food does not appear
				if contains(allergen, food.allergens) && !contains(ingredient, food.ingredients) {
					isMatch = false
				}
			}
			if isMatch {
				m[ingredient] = allergen
			}
		}
	}

	withoutAllergens := []string{}

	for ing, _ := range allIngs {
		if _, ok := m[ing]; !ok {
			withoutAllergens = append(withoutAllergens, ing)
		}
	}

	// now we count how many times they appear..

	for _, food := range foods {
		for _, ing := range food.ingredients {
			if contains(ing, withoutAllergens) {
				out++
			}
		}
	}

	return out
}

func contains(s string, ss []string) bool {
	for _, x := range ss {
		if s == x {
			return true
		}
	}
	return false

}

func getInput() []food {
	in, _ := ioutil.ReadFile("input.txt")

	foods := []food{}
	for _, line := range strings.Split(string(in), "\n") {
		if line == "" {
			continue
		}
		f := food{}
		ingreds := takeWhile(line, func(s string) bool { return s != "(" })
		alergs := dropWhile(line, func(s string) bool { return s != "(" })
		f.ingredients = nonEmpty(strings.Split(ingreds, " "))
		alergs = alergs[len("(contains"):]
		alergs = alergs[0 : len(alergs)-1]

		f.allergens = nonEmpty(strings.Split(alergs, ","))

		foods = append(foods, f)
	}

	return foods

}

// format the input correctly
func nonEmpty(s []string) (out []string) {
	for _, x := range s {
		if len(x) != 0 {
			out = append(out, strings.TrimSpace(x))
		}
	}
	return
}

func takeWhile(s string, pred func(s string) bool) (out string) {
	for _, c := range s {
		if pred(string(c)) {
			out += string(c)
		} else {
			return out
		}
	}
	return out
}

func dropWhile(s string, pred func(s string) bool) (out string) {
	prefix := takeWhile(s, pred)
	return s[len(prefix):]
}
