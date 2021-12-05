package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%v\n", solve())
	fmt.Printf("%v\n", solve2())
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

func (l Line) horizontalLinePoints() []Point {
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
	}
	return points
}

func (l Line) verticalLinePoints() []Point {
	points := []Point{}
	if l.isVertical() {

		if l.start.y < l.end.y {
			for i := l.start.y; i <= l.end.y; i++ {
				points = append(points, Point{x: l.start.x, y: i})
			}
		} else {
			for i := l.end.y; i <= l.start.y; i++ {
				points = append(points, Point{x: l.start.x, y: i})
			}
		}
	}
	return points
}

func (l Line) diagonaLinePoints() []Point {

	points := []Point{}
	if l.isHorizontal() || l.isVertical() {
		return points
	}

	// we need to find out the direction of the line

	if l.start.x < l.end.x {
		// this point starts first
		if l.start.y < l.end.y {

			i := l.start.x
			j := l.start.y

			for i <= l.end.x && j <= l.end.y {
				points = append(points, Point{x: i, y: j})
				i++
				j++
			}

		} else {

			i := l.start.x
			j := l.start.y

			for i <= l.end.x && j >= l.end.y {
				points = append(points, Point{x: i, y: j})
				i++
				j--
			}
		}

	} else {
		if l.start.y < l.end.y {

			i := l.start.x
			j := l.start.y

			for i >= l.end.x && j <= l.end.y {
				points = append(points, Point{x: i, y: j})
				i--
				j++
			}
		} else {

			i := l.start.x
			j := l.start.y

			for i >= l.end.x && j >= l.end.y {
				points = append(points, Point{x: i, y: j})
				i--
				j--
			}
		}
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
		points := append(line.horizontalLinePoints(), line.verticalLinePoints()...)
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

func solve2() int {
	lines := getData()
	fumeMap := map[Point]int{}
	for _, line := range lines {
		points := append(line.horizontalLinePoints(), line.verticalLinePoints()...)
		points = append(points, line.diagonaLinePoints()...)
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
