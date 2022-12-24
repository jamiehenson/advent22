import "dart:io";

List<List<int>> parseLine(String src) {
  return src
      .replaceAll("],", "] ")
      .split(" ")
      .map((s) => s
          .replaceAll(RegExp(r'(\[|\]*)'), "")
          .split(",")
          .map((e) => e.isNotEmpty ? int.parse(e) : 0)
          .toList())
      .toList();
}

bool compareParts(dynamic leftParts, dynamic rightParts) {
  bool sorted = false;

  if (leftParts is List<int>) {
    for (int j = 0; j < leftParts.length; j++) {
      // Must come first, array access error will result in subsequent check otherwise
      if (j >= rightParts.length) {
        return false;
      }

      if (leftParts[j] < rightParts[j]) {
        return true;
      }

      if (leftParts[j] > rightParts[j]) {
        return false;
      }

      if (j == leftParts.length - 1 && rightParts.length - 1 > j) {
        return true;
      }
    }

    return false;
  } else if (leftParts is List<List<int>>) {
    for (int i = 0; i < leftParts.length; i++) {
      sorted = compareParts(leftParts[i], rightParts[i]);
    }
  }

  return sorted;
}

bool determineSortedPair(String pair) {
  List<String> pairLines = pair.split("\n");
  List<List<int>> leftParts = parseLine(pairLines[0]);
  List<List<int>> rightParts = parseLine(pairLines[1]);

  return compareParts(leftParts, rightParts);
}

void main() {
  String src = File("packets.txt").readAsStringSync();

  List<String> pairs = src.split("\n\n");
  List<int> sortedPairIndices = [];

  pairs.asMap().forEach((index, pair) {
    if (determineSortedPair(pair)) {
      sortedPairIndices.add(index + 1);
    }
  });

  print("Sorted pair indices: $sortedPairIndices");
  print(
      "Sum: ${sortedPairIndices.length > 0 ? sortedPairIndices.reduce((value, element) => value + element) : 0}");
}
