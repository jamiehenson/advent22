import fs from "fs";
const main = () => {
  const list = fs.readFileSync("src/cleaning.txt");

  if (!list) {
    return "file not parsed correctly";
  }

  const pairs = list.toString().split("\n");

  const pairEncapsulation = pairs.map((pair) => {
    const elves = pair.split(",");

    const cleaningIds = elves
      .map((elf) => {
        const bounds = elf.split("-");
        const lowerBound = Number(bounds[0]);
        const upperBound = Number(bounds[1]);

        const length = upperBound - lowerBound + 1;

        return Array.from({ length }, (_, i) => i + lowerBound);
      })
      .sort((a, b) => a.length - b.length);

    console.log(cleaningIds);

    return cleaningIds[0].every((id) => cleaningIds[1].includes(id));
  });

  console.log(pairEncapsulation, pairEncapsulation.filter((a) => a).length);
};

main();
