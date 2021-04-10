use std::fs;
use std::str::Lines;

fn main() {
    //println!("{}", solve1());
    println!("{}", solve2());
}

fn lineToValues(line: &String) -> Vec<i32> {
    let mut trimmed = line.trim_start();
    let parts: Vec<&str> = trimmed.split(" ").collect();
    let (a,b,c) = (parts[0].parse::<i32>().unwrap(), parts[1].parse::<i32>().unwrap(), parts[2].parse::<i32>().unwrap());
    return vec![a,b,c];
}

fn solve2() -> isize {
    let input = read_input();

    let mut count = 0;
    let mut i = 0;

    while i < input.len() - 1 {
        let row1 = lineToValues(&input[i]);
        let row2 = lineToValues(&input[i+1]);
        let row3 = lineToValues(&input[i+2]);
        for column in 0..3 {
            let (a,b,c) = (row1[column], row2[column], row3[column]);
            if a + b > c && a + c > b && b + c > a{
                count += 1;
            }
        }

        i += 3
    }

    count
}

fn solve1() -> isize {
    let input = read_input();

    let mut count = 0;
    for line in input {
        if line.len() == 0 {
            continue;
        }
        let mut trimmed = line.trim_start();
        let parts: Vec<&str> = trimmed.split(" ").collect();
        //println!("a: {} b: {} c: {}", parts[0], parts[1], parts[2]);
        let (a,b,c) = (parts[0].parse::<i32>().unwrap(), parts[1].parse::<i32>().unwrap(), parts[2].parse::<i32>().unwrap());
        if a + b > c && a + c > b && b + c > a{
            count += 1;
        }
    }
    count
}


fn read_input() -> Vec<String> {
    let input = fs::read_to_string("input.txt").expect("could not read input");
    let lines = input.split("\n").map(str::to_string).collect();
    lines
}
