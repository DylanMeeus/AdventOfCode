use std::fs;
use std::collections::HashMap;

fn main() {
    println!("{}", solve2());
}

fn solve2() -> String {
    let instr = read_file();

    let mut m: HashMap<(isize, isize), &str> = HashMap::new();
    
    // 1
    m.insert((-2,0), "1");

    // 2 3 4 
    m.insert((-1,-1), "2");
    m.insert((-1,0), "3");
    m.insert((-1,1), "4");

    // 5 6 7 8 9 
    m.insert((0,-2), "5");
    m.insert((0,-1), "6");
    m.insert((0,0), "7");
    m.insert((0,1), "8");
    m.insert((0,2), "9");

    // A B C 
    m.insert((1,0), "B");
    m.insert((1,1), "C");
    m.insert((1,-1), "A");

    // D 
    m.insert((2,0), "D");

    let mut position = (0,-2);

    let mut result = String::new();
    for line in instr.lines() {
        for instruction in line.split("") {
            if instruction == "U" {
                let mut cp = position.clone();
                cp.0 -= 1;
                if m.contains_key(&(cp)) {
                    position.0 -= 1;
                }
            }

            if instruction == "D" {
                let mut cp = position.clone();
                cp.0 += 1;
                if m.contains_key(&(cp)) {
                    position.0 += 1;
                }
            }

            if instruction == "L" {
                let mut cp = position.clone();
                cp.1 -= 1;
                if m.contains_key(&(cp)) {
                    position.1 -= 1;
                }
            }

            if instruction == "R" {
                let mut cp = position.clone();
                cp.1 += 1;
                if m.contains_key(&(cp)) {
                    position.1 += 1;
                }
            }
        }
        // map the position to a key
        match m.get(&position) {
            None => println!("could not find {}", position.0),
            Some(r) => result.push_str(r)
        }
    }

    result
}

fn solve1() -> String {
    let instr = read_file();

    let mut m: HashMap<(isize, isize), &str> = HashMap::new();
    m.insert((0,0), "1");
    m.insert((0,1), "2");
    m.insert((0,2), "3");
    m.insert((1,0), "4");
    m.insert((1,1), "5");
    m.insert((1,2), "6");
    m.insert((2,0), "7");
    m.insert((2,1), "8");
    m.insert((2,2), "9");
    /*
     1 2 3  -> (0,0) (0,1) (0,2)
     4 5 6  -> (1,0) (1,1) (1,2)
     7 8 9  -> (2,0) (2,1) (2,2)
    */
    let mut position = (1,1);

    let mut result = String::new();
    for line in instr.lines() {
        for instruction in line.split("") {
            if instruction == "L" && position.1 > 0 {
                position.1 -= 1;
            }
            if instruction == "U" && position.0 > 0 {
                position.0 -= 1;
            }
            if instruction == "R" && position.1 < 2 {
                position.1 += 1;
            }
            if instruction == "D" && position.0 < 2 {
                position.0 += 1;
            }
        }
        // map the position to a key
        match m.get(&position) {
            None => println!("could not find {}", position.0),
            Some(r) => result.push_str(r)
        }
    }

    result
}


fn read_file() -> String {
    let input = fs::read_to_string("input.txt").expect("could not read input");
    input
}
