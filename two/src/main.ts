interface Round {
  myMove: number;
  theirMove: number;
}

(() => {
  const a = ["A Y", "B X", "C"];

  const encodeMove = (m: String): number => {
    switch (m) {
      case "A":
      case "X":
        return 0;
      case "B":
      case "Y":
        return 1;
      case "C":
      case "Z":
        return 2;
      default:
        throw new Error();
    }
  };

  const calcResult = (r: Round): number => {
    const { myMove: m, theirMove: t } = r;

    if ((m - 1) % 3 == t) {
      return 6;
    } else if ((m + 1) % 3 == t) {
      return 0;
    } else if (m == t) {
      return 3;
    } else {
      throw new Error();
    }
  };

  const b = a.map((b) => {
    const me = b[2];
    const them = b[0];

    const round: Round = {
      myMove: encodeMove(me),
      theirMove: encodeMove(them),
    };

    return round.myMove + 1 + calcResult(round);
  });

  console.log(
    b,
    b.reduce((a, b) => a + b, 0)
  );
})();
