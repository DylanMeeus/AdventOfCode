package main

import (
	"fmt"
	"math"
)

type vector struct {
	x, y, z int
}

type moon struct {
	position, velocity vector
}

func (m moon) String() string {
	return fmt.Sprintf("(%v, %v, %v) {%v, %v, %v}", m.position.x, m.position.y, m.position.z, m.velocity.x, m.velocity.y, m.velocity.z)
}

func main() {
	solve1()
}

func solve1() {
	moons := readData()
	for i := 0; i < 1000; i++ {
		moons = timeStep(moons)
	}
	fmt.Printf("%v\n", calculateEnergy(moons))
}

func timeStep(moons []moon) [] moon {
	return applyVelocity(applyGravity(moons))
}

func calculateEnergy(moons []moon) (e int) {
	for _,m := range moons {
		e += m.position.energy() * m.velocity.energy()
	}
	return
}

func (v vector) energy() int {
	return int(math.Abs(float64(v.x)) + math.Abs(float64(v.y)) + math.Abs(float64(v.z)))
}

func applyVelocity(moons []moon) []moon {
	alteredMoons := make([]moon, len(moons))
	for i, m := range moons {
		m.position.x += m.velocity.x
		m.position.y += m.velocity.y
		m.position.z += m.velocity.z
		alteredMoons[i] = m
	}
	return alteredMoons
}

func applyGravity(moons []moon) []moon {
	alteredMoons := make([]moon, len(moons))
	for i, m := range moons {
		for j, m2 := range moons {
			if i == j {
				continue
			}
			// velx,y,z
			if m.position.x < m2.position.x {
				m = velocityX(m, +1)
			} else if m.position.x > m2.position.x {
				m = velocityX(m, -1)
			}
			// velx,y,z
			if m.position.y < m2.position.y {
				m = velocityY(m, +1)
			} else if m.position.y > m2.position.y {
				m = velocityY(m, -1)
			}
			// velx,y,z
			if m.position.z < m2.position.z {
				m = velocityZ(m, +1)
			} else if m.position.z > m2.position.z {
				m = velocityZ(m, -1)
			}
			alteredMoons[i] = m
		}
	}
	return alteredMoons
}

func velocityX(m moon, alter int) (moon) {
	m.velocity.x += alter
	return m
}
func velocityY(m moon, alter int) (moon) {
	m.velocity.y += alter
	return m
}
func velocityZ(m moon, alter int) (moon) {
	m.velocity.z += alter
	return m
}

func readData() []moon {
	return []moon{
		{
			position: vector{-3, 15, -11},
		},
		{
			position: vector{3, 13, -19},
		},
		{
			position: vector{-13, 18, -2},
		},
		{
			position: vector{6, 0, -1},
		},
	}
}

func testData() []moon {
	return []moon{
		{
			position: vector{-1, 0, 2},
		},
		{
			position: vector{2, -10, -7},
		},
		{
			position: vector{4, -8, 8},
		},
		{
			position: vector{3, 5, -1},
		},
	}
}
