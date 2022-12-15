a = ["A Y", "B X", "C Z"]

b = a.map(b => {
    me = b[2];
    them = b[0];

    meCode = 0;
    themCode = 0;
    result = 0;

    if (me == "B" || me == "Y") {
        meCode = 1;
    } else if (me == "C" || me == "Z") {
        meCode = 2;
    }

    if (them == "B" || them == "Y") {
        themCode = 1;
    } else if (them == "C" || them == "Z") {
        themCode = 2;
    }

    if (((meCode - 1) % 3) == themCode) {
        result = 6;
    } else if (meCode == themCode) {
        result = 3;
    }

    c = meCode + 1;
    
    return c + result;
})

console.log(b, b.reduce((a, b) => a + b, 0))