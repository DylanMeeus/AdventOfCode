package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type board [][]int

type mask [5][5]bool

type markedBoard struct {
	board     board
	boardMask *mask
}

func (m *markedBoard) mark(num int) {
	for row := 0; row < len(m.board); row++ {
		for col := 0; col < len(m.board[row]); col++ {
			if m.board[row][col] == num {
				m.boardMask[row][col] = true
			}
		}
	}
}

func (m markedBoard) checkWinCondition() bool {
	// check all rows and and columns..
	// TODO: this would again be easier with transpose functions (like a few  days ago.. make a
	// transpose function for future AoC challenges..)
	for i := 0; i < len(m.board); i++ {
		if allTrue(m.boardMask.getRow(i)) || allTrue(m.boardMask.getColumn(i)) {
			return true
		}
	}
	return false
}

// sumUnmarked gets the sum of all unmarked numbers
func (m markedBoard) sumUnmarked() int {
	sum := 0

	for row := 0; row < len(m.board); row++ {
		for col := 0; col < len(m.board[row]); col++ {
			if !m.boardMask[row][col] {
				sum += m.board[row][col]
			}
		}
	}
	return sum
}

func allTrue(in [5]bool) bool {
	for _, b := range in {
		if !b {
			return false
		}
	}
	return true
}

func (m mask) getRow(row int) [5]bool {
	out := [5]bool{}
	for col := 0; col < len(m[row]); col++ {
		out[col] = m[row][col]
	}
	fmt.Printf("%v\n", m)
	fmt.Printf("%v\n", out)
	return out
}

func (m mask) getColumn(col int) [5]bool {
	out := [5]bool{}
	for row := 0; row < 5; row++ {
		out[row] = m[row][col]
	}
	return out
}

func newBoard() board {
	b := make(board, 5)
	for i := 0; i < 4; i++ {
		b[i] = make([]int, 5)
	}
	return b
}

func main() {
	fmt.Printf("%v\n", solve())
}

func getData() ([]int, []markedBoard) {
	in, _ := ioutil.ReadFile("./input.txt")

	lines := strings.Split(string(in), "\n")
	intStream := getIntStream(lines[0], ",")
	fmt.Printf("%v\n", intStream)

	boards := []markedBoard{}
	b := newBoard()
	idx := 0
	for _, line := range lines[1:] {
		if line == "" {
			b = newBoard()
			idx = 0
			continue
		}
		line = strings.Replace(line, "  ", " ", -1)
		line = strings.TrimSpace(line)
		b[idx] = getIntStream(line, " ")
		if idx == 4 {
			boards = append(boards, markedBoard{
				board:     b,
				boardMask: &mask{},
			})
		}
		idx++
	}

	return intStream, boards
}

func solve() int {
	numbers, boards := getData()

	for _, num := range numbers {
		for _, b := range boards {
			b.mark(num)
			if b.checkWinCondition() {
				return b.sumUnmarked() * num
			}
		}
	}

	return 0
}

func getIntStream(line string, sep string) []int {
	strNums := strings.Split(line, sep)
	ints := make([]int, len(strNums))

	var err error
	for i, s := range strNums {
		ints[i], err = strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
	}
	return ints
}
