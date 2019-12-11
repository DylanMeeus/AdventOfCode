package main

import (
	"io/ioutil"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	solve1()
	solve2()
}

func readData() (out []int) {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	ss := string(data)
	for _, s := range strings.Split(strings.ReplaceAll(ss, "\n", ""), ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		out = append(out, i)
	}
	return
}

func solve1() {
	data := readData()
	calculate(data)
}

func solve2() {
	data := readData()
	tiles := calculate(data)
	// plot the tiles?
	// we need to know the min(x), max(x) and min(y), max(y) to paint
	var minx, miny, maxx, maxy int
	miny, minx = 10000, 10000
	for k, _ := range tiles {
		if k.x > maxx {
			maxx = k.x
		}
		if k.x < minx {
			minx = k.x
		}
		if k.y > maxy {
			maxy = k.y
		}
		if k.y < miny {
			miny = k.y
		}
	}
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx+20; x++ {
			if p,ok := tiles[point{x,y}]; ok {
				if p == 1 {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Printf("%v %v %v %v\n", minx, maxx, miny, maxy)
}

type robot struct {
	location point
	rotation int // (0 up, 1 right, 2 down, 3 left)
}

type point struct {
	x, y int
}

func (r *robot) step() {
	// based on rotation, move robot marvin
	switch r.rotation {
	case 0:
		r.location.y--
	case 1:
		r.location.x++
	case 2:
		r.location.y++
	case 3:
		r.location.x--
	default:
		panic(r.rotation)
	}
}

func calculate(input []int) map[point]int {
	marvin := robot{}
	input = append(input, make([]int, 3000)...)
	tiles := map[point]int{}
	tiles[point{0, 0}] = 1
	readFunc := func() int {
		tile := tiles[marvin.location]
		return tile
	}
	procColor := true
	processOutput := func(num int) {
		if procColor {
			tiles[marvin.location] = num
			procColor = false
			return
		}
		// rotate marvin
		if num == 1 {
			marvin.rotation = (marvin.rotation + 1) % 4
		} else {
			marvin.rotation = (marvin.rotation - 1)
			if marvin.rotation == -1 {
				marvin.rotation = 3
			}
		}
		marvin.step()
		procColor = true
	}
	var relativeBase int
	relativeBase = 0
	for i := 0; i < len(input); {
		codeparam := strconv.Itoa(input[i])
		var opcode string
		var mode1, mode2, mode3 string
		mode1, mode2, mode3 = "0", "0", "0"
		_ = mode3
		if len(codeparam) != 1 {
			codeparam = "000" + codeparam
			opcode = string(codeparam[len(codeparam)-2]) + string(codeparam[len(codeparam)-1])
			mode1 = string(codeparam[len(codeparam)-3])
			mode2 = string(codeparam[len(codeparam)-4])
			mode3 = string(codeparam[len(codeparam)-5])
		} else {
			opcode = "0" + codeparam
		}
		switch opcode {
		case "99":
			fmt.Printf("painted %v tiles\n", len(tiles))
			return tiles
		case "01":
			ind1, ind2, store := input[i+1], input[i+2], input[i+3]
			a := parseMode(mode1, relativeBase, ind1, input)
			b := parseMode(mode2, relativeBase, ind2, input)
			input = storeMode(mode3, a+b, relativeBase, store, input)
			i += 4
		case "02":
			ind1, ind2, store := input[i+1], input[i+2], input[i+3]
			a := parseMode(mode1, relativeBase, ind1, input)
			b := parseMode(mode2, relativeBase, ind2, input)
			input = storeMode(mode3, a*b, relativeBase, store, input)
			i += 4
		case "03":
			ind := input[i+1]
			if mode1 == "2" {
				input[relativeBase+ind] = readFunc()
			} else if mode1 == "1" {
				// no-op
			} else {
				input[ind] = readFunc()
			}
			i += 2
		case "04":
			store := input[i+1]
			a := store
			var instruction int
			if mode1 == "2" {
				//fmt.Printf("mode 2 out: %v\n", input[relativeBase+a])
				instruction = input[relativeBase+a]
			} else if mode1 == "1" {
				//fmt.Printf("mode 1 out: %v\n", store)
				instruction = store
			} else {
				// mode 0
				//fmt.Printf("mode 0 %v\n", input[a])
				instruction = input[a]
			}
			processOutput(instruction)
			i += 2
		case "05":
			ind1, ind2 := input[i+1], input[i+2]
			a := parseMode(mode1, relativeBase, ind1, input)
			b := parseMode(mode2, relativeBase, ind2, input)
			if a != 0 {
				i = b
			} else {
				i += 3
			}
		case "06":
			ind1, ind2 := input[i+1], input[i+2]
			a := parseMode(mode1, relativeBase, ind1, input)
			b := parseMode(mode2, relativeBase, ind2, input)
			if a == 0 {
				i = b
			} else {
				i += 3
			}
		case "07":
			// less than
			ind1, ind2, store := input[i+1], input[i+2], input[i+3]
			a := parseMode(mode1, relativeBase, ind1, input)
			b := parseMode(mode2, relativeBase, ind2, input)
			if a < b {
				input = storeMode(mode3, 1, relativeBase, store, input)
			} else {
				input = storeMode(mode3, 0, relativeBase, store, input)
			}
			i += 4
		case "08":
			// equals
			ind1, ind2, store := input[i+1], input[i+2], input[i+3]
			a := parseMode(mode1, relativeBase, ind1, input)
			b := parseMode(mode2, relativeBase, ind2, input)
			//store = parseMode(mode3, relativeBase, store, input)
			if a == b {
				input = storeMode(mode3, 1, relativeBase, store, input)
			} else {
				input = storeMode(mode3, 0, relativeBase, store, input)
			}
			i += 4
		case "09":
			alt := parseMode(mode1, relativeBase, input[i+1], input)
			relativeBase += alt
			i += 2
		default:
			i++
		}
	}
	return tiles
}

// return the location of the blahblabhlah?
func parseMode(mode string, relbase, value int, input []int) int {
	switch mode {
	case "0":
		return input[value]
	case "1":
		return value
	case "2":
		return input[relbase+value]
	default:
		panic("fubar")
	}
}

func storeMode(mode string, value, relbase, storeLocation int, input []int) []int {
	switch mode {
	case "0":
		input[storeLocation] = value
		return input
	case "1":
		input[input[storeLocation]] = value
		return input
	case "2":
		input[relbase+storeLocation] = value
		return input
	default:
		panic("fubar")
	}
}
