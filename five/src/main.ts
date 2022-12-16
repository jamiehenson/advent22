import fs from "fs";

const main = () => {
  const file = fs.readFileSync("src/stacks.txt");

  if (!file) {
    return "source material not available, file cannot be read";
  }

  const parts = file.toString().split(" 1   2   3 \n");

  const stacks: string[][] = [[], [], []];
  const stackLines = parts[0].split("\n");

  stackLines
    .filter((line) => line.length > 0)
    .forEach((line) => {
      line[1].match(/[A-Z]/g) && stacks[0].unshift(line[1]);
      line[5].match(/[A-Z]/g) && stacks[1].unshift(line[5]);
      line[9].match(/[A-Z]/g) && stacks[2].unshift(line[9]);
    });

  const instructions = parts[1]
    .split("\n")
    .slice(1)
    .map((line) =>
      line
        .split(" ")
        .filter((_, index) => index % 2 == 1)
        .map((value) => Number(value))
    );

  instructions.forEach((line) => {
    for (let i = 0; i < line[0]; i++) {
      const transitValue = stacks[line[1] - 1].pop() || "";
      stacks[line[2] - 1].push(transitValue);
    }
  });

  console.log(stacks.map((stack) => stack[stack.length - 1]));
};

main();
