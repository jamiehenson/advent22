use std::{collections::HashMap, fs};

fn main() {
    let file = fs::read_to_string("src/stacks.txt").expect("File to be present");

    let parts: Vec<&str> = file.split("1   2   3").collect();

    let stack_text = parts.get(0);
    let instruction_text = parts.get(1);
    let mut stacks: HashMap<i32, Vec<char>> =
        HashMap::from([(1, Vec::new()), (2, Vec::new()), (3, Vec::new())]);

    let mut lines: Vec<&str> = stack_text
        .expect("First part should be available")
        .split("\n")
        .collect();

    lines.pop();

    for line in &lines {
        let chars: Vec<char> = line.chars().collect();
        let first_char = chars.get(1).expect("Should have char at 1st index");
        let second_char = chars.get(5).expect("Should have char at 4th index");
        let third_char = chars.get(9).expect("Should have char at 7th index");

        if first_char.is_alphabetic() {
            stacks.get_mut(&1).unwrap().insert(0, *first_char);
        }

        if second_char.is_alphabetic() {
            stacks.get_mut(&2).unwrap().insert(0, *second_char);
        }

        if third_char.is_alphabetic() {
            stacks.get_mut(&3).unwrap().insert(0, *third_char);
        }
    }

    let instructions: Vec<Vec<i32>> = instruction_text
        .expect("Should have instruction text")
        .split("\n")
        .map(|line| {
            return line
                .split(" ")
                .enumerate()
                .filter(|(i, _)| i % 2 == 1)
                .map(|(_, e)| {
                    if !e.is_empty() {
                        e.parse::<i32>().unwrap()
                    } else {
                        0
                    }
                })
                .collect();
        })
        .collect();

    for inst in &instructions[2..] {
        let first = inst.get(0).unwrap();
        let second = inst.get(1).unwrap();
        let third = inst.get(2).unwrap();

        for _ in 0..*first {
            let transit_value = stacks.get_mut(&second).unwrap().pop().unwrap();
            stacks.get_mut(&third).unwrap().push(transit_value);
        }
    }

    for i in 1..=3 {
        println!(
            "Top of stack {}: {:?}",
            i,
            stacks.get(&i).unwrap().last().unwrap()
        );
    }
}
