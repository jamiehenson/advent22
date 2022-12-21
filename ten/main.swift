
import Foundation

func computeCycles(result: String) -> [Int : Int] {
  var instructions = result.components(separatedBy: "\n")
  var clock: Int = 1
  var register: Int = 1
  var queue = [Int : Int]()
  var cycles: [Int: Int] = [Int: Int]()

  while instructions.count > 0 {
    if instructions.count > 0 && queue.count == 0 {
      let parts: [String] = instructions[0].components(separatedBy: " ")
      let operation: String = parts[0]
      var value: Int = 0
      var clockOffset: Int = clock

      switch operation {
        case "addx":
          value = Int(parts[1]) ?? 0
          clockOffset += 1
        case "noop":
          break
        default:
          break
      }

      queue[clockOffset] = value      

      instructions.removeFirst()
    }

    if queue[clock] != nil {
      register += queue[clock] ?? 0
      queue.removeValue(forKey: clock)
    }

    clock += 1
    cycles[clock] = register
  }

  return cycles
}

func simulateCPU(path: String, cycleIndices: [Int]) -> Int {
  var result = ""

  do {
    result = try String(contentsOfFile: path, encoding: .utf8)
  } catch {
    print("File not read.")
  }

  let cycles = computeCycles(result: result)
  let targetCycleValues = cycleIndices.map{ (cycles[$0] ?? 0) * $0 }

  return targetCycleValues.reduce(0, +)
}

print(simulateCPU(path: "instructions.txt", cycleIndices: [1,2,3,4,5]))
print(simulateCPU(path: "instructions_large.txt", cycleIndices: [20, 60, 100, 140, 180, 220]))