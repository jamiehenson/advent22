use std::fs;

struct Round {
    my_move: i32,
    their_move: i32,
}

fn calculate_result(r: Round) -> i32 {
    let Round {
        my_move: x,
        their_move: y,
    } = r;

    if (x - 1).rem_euclid(3) == y {
        return 6;
    } else if (x + 1).rem_euclid(3) == y {
        return 0;
    } else if x == y {
        return 3;
    } else {
        return -1;
    }
}

fn encode_move_char(m: char) -> i32 {
    let x = "X".parse::<char>().unwrap();
    let y = "Y".parse::<char>().unwrap();
    let z = "Z".parse::<char>().unwrap();

    let a = "A".parse::<char>().unwrap();
    let b = "B".parse::<char>().unwrap();
    let c = "C".parse::<char>().unwrap();

    if m == x || m == a {
        // Rock
        return 1;
    } else if m == y || m == b {
        // Paper
        return 2;
    } else if m == z || m == c {
        // Scissors
        return 3;
    } else {
        return -1;
    }
}

fn main() {
    let scores = fs::read_to_string("src/scores.txt").expect("Can't read the file");
    let split_scores = scores.split("\n");

    let round_totals: Vec<i32> = split_scores
        .into_iter()
        .map(|x| {
            let round = Round {
                my_move: encode_move_char(x.chars().nth(2).unwrap()),
                their_move: encode_move_char(x.chars().nth(0).unwrap()),
            };

            let a = round.my_move;
            let b = calculate_result(round);

            return a + b;
        })
        .collect();

    let round_sum: i32 = round_totals.iter().sum();

    println!("Rounds: {:?}. Round total: {}", round_totals, round_sum);
}
