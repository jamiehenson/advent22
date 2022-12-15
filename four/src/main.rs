use std::fs;

fn main() {
    let file = fs::read_to_string("src/cleaning.txt").expect("File to be present");

    let pairs = file.split("\n");

    let pair_encapsulation: Vec<bool> = pairs
        .map(|pair| {
            let elves = pair.split(",");

            let mut ranges: Vec<Vec<i32>> = elves
                .map(|elf| {
                    let bounds: Vec<&str> = elf.split("-").collect();
                    let lower_bound = bounds.get(0).expect("").parse::<i32>().expect("msg");
                    let upper_bound = bounds.get(1).expect("").parse::<i32>().expect("msg");

                    return (lower_bound..upper_bound).collect::<Vec<_>>();
                })
                .collect();

            ranges.sort_by(|a, b| a.len().cmp(&b.len()));

            return ranges
                .get(0)
                .expect("")
                .iter()
                .all(|a| ranges.iter().nth(1).expect("").contains(a));
        })
        .collect();

    println!("{:?}", pair_encapsulation);

    let count = &pair_encapsulation
        .into_iter()
        .filter(|&a| a == true)
        .count();

    println!("{:?}", count);
}
