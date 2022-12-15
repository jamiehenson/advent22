use std::{collections::HashSet, fs};

fn main() {
    let file = fs::read_to_string("src/backpacks.txt").expect("File could not be read");

    let backpacks = file.split("\n");

    let mut common_items = HashSet::new();

    for backpack in backpacks {
        let start = &backpack[..backpack.len() / 2];
        let end = &backpack[backpack.len() / 2..];

        let set: HashSet<char> = start.chars().collect();

        let intersection: HashSet<char> = end.chars().filter(|s| set.contains(&s)).collect();

        for c in intersection {
            if !common_items.contains(&c) {
                common_items.insert(c);
            }
        }
    }

    println!("Common items: {:?}", common_items);

    let priorities: Vec<u8> = common_items
        .into_iter()
        .map(|item| {
            let a = item as u8;

            if a >= 97u8 && a <= 122u8 {
                // lowercase
                return a - 96u8;
            } else if a >= 65u8 && a <= 90u8 {
                // uppercase
                return a - 38u8;
            } else {
                return 0;
            }
        })
        .collect();

    let priority_total: u8 = priorities.iter().sum();

    println!("Priority total: {:?}", priority_total);
}
