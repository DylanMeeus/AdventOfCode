package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Point struct{ row, col int }

func (p Point) Down() Point {
	return Point{
		row: p.row,
		col: p.col + 1,
	}
}

func (p Point) Right() Point {
	return Point{
		row: p.row + 1,
		col: p.col,
	}
}

type Cave struct {
	Map map[Point]int
	End Point
}

func main() {
	fmt.Printf("%v\n", solve())
}

func getData() Cave {
	in, _ := ioutil.ReadFile("./input.txt")

	cave := Cave{
		Map: map[Point]int{},
	}

	colCount := 0
	rowCount := 0
	for row, line := range strings.Split(string(in), "\n") {
		if line == "" {
			continue
		}
		strValues := strings.Split(line, "")
		for col, v := range strValues {
			i, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			cave.Map[Point{row, col}] = i
		}
		colCount = len(strValues)
		rowCount++
	}

	cave.End = Point{rowCount - 1, colCount - 1}

	return cave
}

func solve() int {
	data := getData()
	return riskMapper(data) - 1
}

func riskMapper(cave Cave) int {
	risk := riskAt(cave, Point{0, 0}, map[Point]int{})
	return risk
}

func riskAt(cave Cave, point Point, memo map[Point]int) int {

	if value, ok := memo[point]; ok {
		return value
	}

	if _, ok := cave.Map[point]; !ok {
		return int(10e9) // just a ton of risk..
	}

	if point == cave.End {
		return cave.Map[point]
	}

	// else we get to do recursive fun!

	this := cave.Map[point]
	belowRisk := riskAt(cave, point.Down(), memo)
	rightRisk := riskAt(cave, point.Right(), memo)

	totalRisk := this
	if belowRisk < rightRisk {
		totalRisk += belowRisk
	} else {
		totalRisk += rightRisk
	}

	memo[point] = totalRisk
	return memo[point]
}
