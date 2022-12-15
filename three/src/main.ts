import { readFileSync } from "fs";

const main = () => {
  const file = readFileSync("src/backpacks.txt");

  if (!file) {
    return "all is lost";
  }

  const backpacks: String[] = file.toString().split("\n");

  const commonItems: String[] = backpacks.flatMap((backpack) => {
    const start = backpack.slice(0, backpack.length / 2);
    const end = backpack.slice(backpack.length / 2, backpack.length);

    const match = start.match(new RegExp(`[${end}]`, "g"));
    return Array.from(new Set(match));
  });

  const priorities: number[] = commonItems.map((item) => {
    let priority = 0;
    const asciiCode = item.charCodeAt(0);

    if (asciiCode >= 97 && asciiCode <= 122) {
      // lowercase
      priority = asciiCode - 96;
    } else if (asciiCode >= 65 && asciiCode <= 90) {
      // uppercase
      priority = asciiCode - 38;
    }

    return priority;
  });

  const totalPriority: number = priorities.reduce((a, v) => a + v, 0);

  console.log(`Common items: ${commonItems}`);
  console.log(`Total priority: ${totalPriority}`);
};

main();
