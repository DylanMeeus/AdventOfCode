package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	numberRegex = regexp.MustCompile(`[0-9]+`)
)

type Monkey struct {
	ID        int
	Items     *Queue
	Operation func(i int) int
	TestDiv   int
	Ytarget   int
	Ntarget   int
}

func main() {
	fmt.Println(solve1())
}

func solve1() int {
	monkeys := getData()
	monkeyID := 0
	next := func() {
		monkeyID++
		if monkeyID > len(monkeys)-1 {
			monkeyID = 0
		}
	}

	inspections := map[int]int{}

	for round := 0; round < 20; round++ {
		for turn := 0; turn < len(monkeys); turn++ {
			currentMonkey := monkeys[monkeyID]
			for !currentMonkey.Items.Empty() {
				item, ok := currentMonkey.Items.Pop()
				if ok {
					inspections[currentMonkey.ID]++
					item = currentMonkey.Operation(item)
					item /= 3
					if item%currentMonkey.TestDiv == 0 {
						monkeys[currentMonkey.Ytarget].Items.Push(item)
					} else {
						monkeys[currentMonkey.Ntarget].Items.Push(item)
					}
				}
			}
			next()
		}
	}

	out := []int{}
	for _, value := range inspections {
		out = append(out, value)
	}

	sort.Ints(out)

	return out[len(out)-1] * out[len(out)-2]
}

func getData() []Monkey {
	f, err := ioutil.ReadFile("./input.txt")
	handleError(err)
	lines := strings.Split(string(f), "\n")

	monkeys := []Monkey{}

	for i := 0; i < len(lines)-5; i += 6 {
		monkeys = append(monkeys, parseMonkey(lines[i:i+6]))
	}

	fmt.Println(monkeys)
	return monkeys
}

func removeSpace(s string) string {
	return strings.Replace(s, " ", "", -1)
}

func parseItems(itemline string) *Queue {
	values := removeSpace(itemline[len("starting items: "):])

	q := &Queue{}
	for _, item := range strings.Split(values, ",") {
		i, err := strconv.Atoi(item)
		handleError(err)
		q.Push(i)
	}

	return q
}

func parseOperation(operationline string) func(int) int {
	if strings.Contains(operationline, "old + ") {
		output := numberRegex.Find([]byte(operationline))
		num, err := strconv.Atoi(string(output))
		handleError(err)
		return func(i int) int {
			return i + num
		}
	}

	if strings.Contains(operationline, "old * old") {
		return func(i int) int { return i * i }
	}

	if strings.Contains(operationline, "old * ") {
		output := numberRegex.Find([]byte(operationline))
		num, err := strconv.Atoi(string(output))
		handleError(err)
		return func(i int) int {
			return i * num
		}
	}
	panic("should not be here")
}

func parseSingleNum(line string) int {
	output := numberRegex.Find([]byte(line))
	num, err := strconv.Atoi(string(output))
	handleError(err)
	return num
}

func parseMonkey(lines []string) Monkey {
	idline := lines[0]
	itemline := strings.TrimSpace(lines[1])
	operationline := strings.TrimSpace(lines[2])
	testline := strings.TrimSpace(lines[3])
	trueline := strings.TrimSpace(lines[4])
	falseline := strings.TrimSpace(lines[5])

	ids := idline[len(idline)-2]
	id, err := strconv.Atoi(string(ids))
	handleError(err)

	return Monkey{
		ID:        id,
		Items:     parseItems(itemline),
		Operation: parseOperation(operationline),
		TestDiv:   parseSingleNum(testline),
		Ytarget:   parseSingleNum(trueline),
		Ntarget:   parseSingleNum(falseline),
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

type Queue struct {
	inner []int
}

func (q *Queue) Empty() bool {
	return len(q.inner) == 0
}

func (q *Queue) Push(i int) {
	q.inner = append(q.inner, i)
}

func (q *Queue) Pop() (int, bool) {
	if q.Empty() {
		return 0, false
	}
	first := q.inner[0]
	q.inner = q.inner[1:]
	return first, true
}

func (q *Queue) Print() {
	for _, i := range q.inner {
		fmt.Println(i)
	}
}
