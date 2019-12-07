package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	test = []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
)

func main() {
	//solve1()
	solve2()
}

func readPermutations() (out [][]int) {
	data, err := ioutil.ReadFile("./permutations2.txt")
	if err != nil {
		panic(err)
	}
	ss := string(data)[1 : len(string(data))-1]
	permutations := make([][]int, 0)
	for i := 0; i < len(ss); i += 12 {
		part := ss[i : i+11]
		numbers := part[1 : len(part)-1]
		nums := strings.Split(numbers, ",")
		numss := []int{}
		for _, n := range nums {
			num, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			numss = append(numss, num)
		}
		permutations = append(permutations, numss)
	}
	return permutations
}

func readData() (out []int) {
	data, err := ioutil.ReadFile("./test.txt")
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

func solve2() {
	perm := readPermutations()
	var max int
	//amps := []string{"A", "B", "C", "D", "E"}
	// go routine per amp?
	for _, p := range perm {
		p = []int{9,8,7,6,5}
		achan := make(chan int, 2)
		bchan := make(chan int, 1)
		cchan := make(chan int, 1)
		dchan := make(chan int, 1)
		echan := make(chan int, 1)
		chancount := 4
		total := make(chan int, chancount)
		achan <- p[0]
		achan <- 0
		bchan <- p[1]
		cchan <- p[2]
		dchan <- p[3]
		echan <- p[4]
		go runAmp([]int{p[0],0}, echan, achan, total)
		go runAmp([]int{p[1]}, achan, bchan, total)
		go runAmp([]int{p[2]}, bchan, cchan, total)
		go runAmp([]int{p[3]}, cchan, dchan, total)
		go runAmp([]int{p[4]}, dchan, echan, total)
		for bfr := 0; bfr < chancount+1; bfr++ {
			t := <-total
			if t > max {
				max = t
			}
		}
		break
	}
	fmt.Println(max)
}

func runAmp(start []int, r, w, total chan int) {
	data := readData()
	// + the output of the previous one
	go calculate(data, func() int {
		t := <- r
		return t
	}, w,  total)
	/*
	for _,s := range start {
		r <-s
	}
	 */
	/*
		strOut := ""
		for _, o := range output {
			strOut += strconv.Itoa(o)
		}
		intOut, _ := strconv.Atoi(strOut)
		if intOut > max {
			max = intOut
		}
	*/
}

/*
func solve1() {
	perm := readPermutations()
	var max int
	amps := []string{"A", "B", "C", "D", "E"}
	for _, p := range perm {
		var ampOutput int
		for _, _ = range amps {
			data := readData()
			ampInput := []int{}
			ampInput = append(ampInput, p[0], ampOutput)
			p = p[1:]
			// + the output of the previous one
			output := calculate(data, func() int {
				i := ampInput[0]
				ampInput = ampInput[1:]
				return i
			})
			strOut := ""
			for _, o := range output {
				strOut += strconv.Itoa(o)
			}
			intOut, _ := strconv.Atoi(strOut)
			if intOut > max {
				max = intOut
			}
			ampOutput = intOut
		}
	}
	fmt.Println(max)
}
*/

// calculate collects println statements and returns those
func calculate(input []int, readFunc func() int, w, total chan int) []int {
	printstmt := []int{}
	for i := 0; i < len(input); {
		codeparam := strconv.Itoa(input[i])
		var opcode string
		var mode1, mode2 string
		mode1, mode2 = "0", "0"
		if len(codeparam) != 1 {
			codeparam = "00" + codeparam
			opcode = string(codeparam[len(codeparam)-2]) + string(codeparam[len(codeparam)-1])
			mode1 = string(codeparam[len(codeparam)-3])
			mode2 = string(codeparam[len(codeparam)-4])
		} else {
			opcode = "0" + codeparam
		}
		switch opcode {
		case "99":
			close(w)
			total <- printstmt[len(printstmt)-1]
			return printstmt
		case "01":
			ind1, ind2, store := input[i+1], input[i+2], input[i+3]
			a := ind1
			b := ind2
			if mode1 == "0" {
				a = input[ind1]
			}
			if mode2 == "0" {
				b = input[ind2]
			}
			input[store] = a + b
			i += 4
		case "02":
			ind1, ind2, store := input[i+1], input[i+2], input[i+3]
			a := ind1
			b := ind2
			if mode1 == "0" {
				a = input[ind1]
			}
			if mode2 == "0" {
				b = input[ind2]
			}
			input[store] = a * b
			i += 4
		case "03":
			store := input[i+1]
			input[store] = readFunc()
			i += 2
		case "04":
			store := input[i+1]
			w <- input[store]
			printstmt = append(printstmt, input[store])
			// aggregate them all
			i += 2
		case "05":
			ind1, ind2 := input[i+1], input[i+2]
			a := ind1
			b := ind2
			if mode1 == "0" {
				a = input[ind1]
			}
			if mode2 == "0" {
				b = input[ind2]
			}
			if a != 0 {
				i = b
			} else {
				i += 3
			}
		case "06":
			ind1, ind2 := input[i+1], input[i+2]
			a := ind1
			b := ind2
			if mode1 == "0" {
				a = input[ind1]
			}
			if mode2 == "0" {
				b = input[ind2]
			}
			if a == 0 {
				i = b
			} else {
				i += 3
			}
		case "07":
			// less than
			ind1, ind2, store := input[i+1], input[i+2], input[i+3]
			a := ind1
			b := ind2
			if mode1 == "0" {
				a = input[ind1]
			}
			if mode2 == "0" {
				b = input[ind2]
			}
			if a < b {
				input[store] = 1
			} else {
				input[store] = 0
			}
			i += 4
		case "08":
			// equals
			ind1, ind2, store := input[i+1], input[i+2], input[i+3]
			a := ind1
			b := ind2
			if mode1 == "0" {
				a = input[ind1]
			}
			if mode2 == "0" {
				b = input[ind2]
			}
			if a == b {
				input[store] = 1
			} else {
				input[store] = 0
			}
			i += 4
		default:
			i++
		}
	}
	strOut := ""
	for _, o := range printstmt {
		strOut += strconv.Itoa(o)
	}
	intOut, err := strconv.Atoi(strOut)
	if err != nil{
		panic(err)
	}
	fmt.Printf("writing %v\n", intOut)
	w <- intOut
	return input
}
