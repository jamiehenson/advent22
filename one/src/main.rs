use std::collections::HashMap;

fn main() {
    let elves = HashMap::from([
        ("Elf 1", vec![1000, 2000, 3000]),
        ("Elf 2", vec![4000]),
        ("Elf 3", vec![5000, 6000]),
        ("Elf 4", vec![7000, 8000, 9000]),
        ("Elf 5", vec![10000]),
    ]);

    let mut winner = ("", 0);

    for (elf, num) in &elves {
        let sum = num.iter().sum();

        if sum > winner.1 {
            winner = (elf, sum);
        }

        println!("{} - {}", elf, sum);
    }

    println!("Winning elf: {} with {}", winner.0, winner.1);
}
