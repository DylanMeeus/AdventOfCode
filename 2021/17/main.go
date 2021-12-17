package main

import "fmt"

type Target struct {
	startX, endX int
	startY, endY int
}

func (t Target) contains(p Point) bool {
	inX := p.x >= t.startX && p.x <= t.endX
	inY := p.y <= t.startY && p.y >= t.endY
	return inX && inY
}

func (t Target) overShot(p Point) bool {
	return p.x > t.endX || p.y < t.endY
}

type Point struct {
	x, y int
}

type Velocity Point

var (
	INPUT = Target{
		startX: 175,
		endX:   227,
		startY: -79,
		endY:   -134,
	}

	TEST_INPUT = Target{
		startX: 20,
		endX:   30,
		startY: -5,
		endY:   -10,
	}
)

func main() {
	fmt.Printf("%v\n", solve())
}

func getData() Target { return INPUT }

func solve() int {

	area := getData()

	maxY := 0
	for xVel := 0; xVel < 2000; xVel++ {
		for yVel := 0; yVel < 2000; yVel++ {
			res, ok := shoot(Velocity{xVel, yVel}, area)
			if ok && res > maxY {
				maxY = res
			}
		}
	}

	return maxY
}

// shoot fires the probe at initial velocity (starting at <0,0>)
func shoot(velocity Velocity, target Target) (int, bool) {
	position := Point{0, 0}
	maxY := 0
	for {
		position.x += velocity.x
		position.y += velocity.y

		if position.y > maxY {
			maxY = position.y
		}
		// x drags towards 0
		if velocity.x != 0 {
			velocity.x--
		}

		if target.contains(position) {
			return maxY, true
		}

		// check if we overshot the target
		if target.overShot(position) {
			return -1, false
		}

		velocity.y--
	}
}
