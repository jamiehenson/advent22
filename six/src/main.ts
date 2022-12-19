import fs from "fs";

const main = () => {
  const file = fs.readFileSync("src/signals.txt");

  if (!file) {
    return;
  }

  const signals = file.toString().split("\n");

  const markerIndexes = signals.map((signal) => {
    if (signal.length < 5) {
      return;
    }

    for (let i = 0; i < signal.length - 4; i++) {
      const occurences: { [key: string]: number } = {};
      signal
        .slice(i, i + 4)
        .split("")
        .forEach((element) => {
          occurences[element] = isNaN(occurences[element])
            ? 1
            : occurences[element] + 1;
        });

      if (Object.values(occurences).filter((value) => value > 1).length > 0) {
        continue;
      } else {
        return i + 4;
      }
    }

    return -1;
  });

  markerIndexes.map((m, i) => {
    if (m != -1) {
      console.log(`Signal ${i}: marker at ${m}`);
    } else {
      console.log(`No marker for signal ${i}`);
    }
  });
};

main();
