package opcodes

// addition

var Functions = []Op{Addr,Mulr,Banr,Borr,Setr,Gtrr,Eqrr,Addi,Muli,Bani,Bori,Seti,Gtri,Eqri,Gtir,Eqir}


var Operators = map[string]Op{
    "addi": Addi,
    "seti": Seti,
    "muli": Muli,
    "mulr": Mulr,
    "eqrr": Eqrr,
    "addr": Addr,
    "gtrr": Gtrr,
}


type Op func(in1, in2, out int, regs *[6]int)

func Addr(in1, in2, out int, regs *[6]int) {
    regs[out] = regs[in1] + regs[in2]
}

func Addi(in1, in2, out int, regs *[6]int) {
    regs[out] = regs[in1] + in2
}


// multiplication

func Mulr(in1, in2, out int, regs *[6]int) {
    regs[out] = regs[in1] * regs[in2]
}

func Muli(in1, in2, out int, regs *[6]int) {
    regs[out] = regs[in1] * in2
}

// bitwise &

func Banr(in1, in2, out int, regs *[6]int) {
    regs[out] = regs[in1] & regs[in2]
}

func Bani(in1, in2, out int, regs *[6]int) {
    regs[out] = regs[in1] & in2
}

// bitwise |

func Borr(in1, in2, out int, regs *[6]int) {
    regs[out] = regs[in1] | regs[in2]
}

func Bori(in1, in2, out int, regs *[6]int) {
    regs[out] = regs[in1] | in2
}

// assignment

func Setr(in1, in2, out int, regs *[6]int) {
    regs[out] = regs[in1] 
}

func Seti(in1, in2, out int, regs *[6]int) {
    regs[out] = in1
}

// > testing

func Gtir(in1, in2, out int, regs *[6]int) {
    if in1 > regs[in2] {
        regs[out] = 1
    } else {
        regs[out] = 0
    }
}

func Gtri(in1, in2, out int, regs *[6]int) {
    if regs[in1] > in2 {
        regs[out] = 1 
    } else {
        regs[out] = 0
    }
}

func Gtrr(in1, in2, out int, regs *[6]int) {
    if regs[in1] > regs[in2] {
        regs[out] = 1 
    } else {
        regs[out] = 0
    }
}

// == testing


func Eqir(in1, in2, out int, regs *[6]int) {
    if in1 == regs[in2] {
        regs[out] = 1
    } else {
        regs[out] = 0
    }
}

func Eqri(in1, in2, out int, regs *[6]int) {
    if regs[in1] == in2 {
        regs[out] = 1
    } else {
        regs[out] = 0
    }
}

func Eqrr(in1, in2, out int, regs *[6]int) {
    if regs[in1] == regs[in2] {
        regs[out] = 1
    } else {
        regs[out] = 0
    }
}


