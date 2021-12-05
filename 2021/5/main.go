package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	fmt.Printf("%v\n", solve())
}

type Point struct {
	x, y int
}

type Line struct {
	start Point
	end   Point
}

func (l Line) isHorizontal() bool {
	return l.start.y == l.end.y
}

func (l Line) isVertical() bool {
	return l.start.x == l.end.x
}

func (l Line) pointsOnLine() []Point {
	points := []Point{}
	if l.isHorizontal() {

		if l.start.x < l.end.x {
			for i := l.start.x; i <= l.end.x; i++ {
				points = append(points, Point{x: i, y: l.start.y})
			}
		} else {
			for i := l.end.x; i <= l.start.x; i++ {
				points = append(points, Point{x: i, y: l.start.y})
			}
		}

	} else if l.isVertical() {

		if l.start.y < l.end.y {
			for i := l.start.y; i <= l.end.y; i++ {
				points = append(points, Point{x: l.start.x, y: i})
			}
		} else {
			for i := l.end.y; i <= l.start.y; i++ {
				points = append(points, Point{x: l.start.x, y: i})
			}
		}

	} else {
		// no-op
	}
	return points
}

func getData() []Line {
	input, _ := ioutil.ReadFile("./input.txt")
	strInput := string(input)

	inputLines := strings.Split(strings.Replace(strInput, " ", "", -1), "\n")

	lines := []Line{}
	for _, line := range inputLines {
		if line == "" {
			continue
		}
		line = strings.TrimSpace(line)
		parts := strings.Split(line, "->")
		firstPoint := stringPointToPoint(parts[0])
		secondPoint := stringPointToPoint(parts[1])

		lines = append(lines, Line{start: firstPoint, end: secondPoint})
	}
	return lines
}

func stringPointToPoint(s string) Point {
	parts := strings.Split(s, ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return Point{x, y}
}

func solve() int {
	lines := getData()
	fumeMap := map[Point]int{}
	for _, line := range lines {
		points := line.pointsOnLine()
		for _, point := range points {
			fumeMap[point]++
		}
	}

	result := 0

	for _, value := range fumeMap {
		if value >= 2 {
			result++
		}
	}

	return result
}
