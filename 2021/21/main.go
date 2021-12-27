package main

import "fmt"

func main() {
	fmt.Printf("%v\n", solve())
	//d := &Die{0}
	//fmt.Printf("%v\n", d.next3())
}

const (
	BOARD_LENGTH = 10
)

type Player struct {
	score    int
	position int
}

type Die struct {
	face int // current face value of die
}

func (d *Die) next3() int {
	steps := 0
	for i := 0; i < 3; i++ {
		d.roll()
		steps += d.face
	}
	return steps
}

func (d *Die) roll() {
	d.face++
	if d.face > 100 {
		d.face = 1
	}
}

func (p *Player) move(steps int) {
	for i := 0; i < steps; i++ {
		p.position++
		if p.position > 10 {
			p.position = 1
		}
	}
	p.score += p.position
}

func getData() (*Player, *Player) {
	return &Player{
			score:    0,
			position: 8,
		}, &Player{
			score:    0,
			position: 4,
		}
}

func solve() int {
	p1, p2 := getData()

	// 100-sided die
	die := &Die{0}
	rolls := 0
	for p1.score < 1000 && p2.score < 1000 {
		total_moves := die.next3()
		p1.move(total_moves)
		rolls++
		if p1.score >= 1000 {
			break
		}
		total_moves = die.next3()
		p2.move(total_moves)
		rolls++
	}

	if p2.score < p1.score {
		return p2.score * (rolls * 3)
	} else {
		return p1.score * (rolls * 3)
	}
}
