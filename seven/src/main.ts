import fs from "fs";

type Entry = FileEntry | DirectoryEntry;

interface FileEntry {
  name: string;
  fileSize: number;
}

interface DirectoryEntry {
  name: string;
  children: Entry[];
}

interface DirectorySize {
  name: string;
  recursiveSize: number;
}

const isDirectoryEntry = (entry: Entry): entry is DirectoryEntry =>
  (entry as DirectoryEntry).children !== undefined;

const isFileEntry = (entry: Entry): entry is FileEntry =>
  (entry as FileEntry).fileSize !== undefined;

const buildFileSystem = (instructions: string[]): DirectoryEntry => {
  const fileSystem: DirectoryEntry = {
    name: "/",
    children: [],
  };

  let currentPath: string[] = [];
  let currentDirectory: DirectoryEntry = fileSystem;

  instructions.forEach((instruction) => {
    const splitInstruction = instruction.split(" ");

    if (splitInstruction.includes("cd")) {
      // change directory command
      if (splitInstruction[2] == "/") {
        currentPath = [];
        currentDirectory = fileSystem;
      } else {
        if (splitInstruction[2] == "..") {
          currentPath.pop();
        } else {
          currentPath.push(splitInstruction[2]);
        }

        let temporaryEntry = fileSystem;
        currentPath.forEach((level) => {
          const targetDir = temporaryEntry.children
            .filter(isDirectoryEntry)
            .find(({ name }) => name === level);

          if (targetDir) {
            temporaryEntry = targetDir;
          }
        });

        currentDirectory = temporaryEntry;
      }
    } else if (splitInstruction.includes("ls")) {
      // list command - noop
      return;
    } else if (splitInstruction.includes("dir")) {
      // directory
      currentDirectory.children.push({
        name: splitInstruction[1],
        children: [],
      });
    } else {
      // file
      currentDirectory.children.push({
        name: splitInstruction[1],
        fileSize: Number(splitInstruction[0]),
      });
    }
  });

  return fileSystem;

  // console.log(JSON.stringify(fileSystem, null, 2)); // Filesystem log
};

const getDirectories = (
  directory: DirectoryEntry,
  parent?: DirectoryEntry
): DirectoryEntry[] => {
  const localDirs = directory.children.filter(isDirectoryEntry);

  if (localDirs.length > 0) {
    return localDirs.flatMap((dir) => getDirectories(dir, directory));
  } else {
    return [directory, parent || directory];
  }
};

const calculateDirectorySize = (directory: DirectoryEntry): number => {
  const sizes = directory.children.flatMap((child) => {
    if (isFileEntry(child)) {
      return child.fileSize;
    } else if (isDirectoryEntry(child)) {
      return calculateDirectorySize(child);
    }
  });

  return sizes.reduce((acc, cur) => (acc || 0) + (cur || 0), 0) || 0;
};

const main = () => {
  const file = fs.readFileSync("src/instructions.txt");

  if (!file) {
    return;
  }

  const instructions = file.toString().split("\n");

  const fileSystem = buildFileSystem(instructions);
  const directories = getDirectories(fileSystem);
  const directorySizes = directories.map(
    (dir): DirectorySize => ({
      name: dir.name,
      recursiveSize: calculateDirectorySize(dir),
    })
  );

  console.log("Directories:");
  directorySizes
    .sort((a, b) => b.recursiveSize - a.recursiveSize)
    .forEach((dir) => console.log(`${dir.name}: ${dir.recursiveSize}`));

  const eligibleDirectories = directorySizes
    .filter((dir) => dir.recursiveSize <= 100000)
    .sort((a, b) => b.recursiveSize - a.recursiveSize);

  console.log("\nEligible directories (< 100000):");
  eligibleDirectories.map((dir) => {
    console.log(`${dir.name}: ${dir.recursiveSize}`);
  });
  console.log(
    `\nSum: ${eligibleDirectories.reduce(
      (acc, curr) => acc + curr.recursiveSize,
      0
    )}`
  );
};

main();
