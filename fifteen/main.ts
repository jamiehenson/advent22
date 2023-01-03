import fs from "fs"

interface Coordinate {
  x: number
  y: number
}

interface Bounds {
  minX: number
  maxX: number
}

const calculateManhattanDistance = (sensor: Coordinate, beacon: Coordinate): number =>
  Math.abs(sensor.x - beacon.x) + Math.abs(sensor.y - beacon.y)

const main = (y: number) => {
  const file = fs.readFileSync("beacons.txt")
  const splitFile = file.toString().split("\n")

  const sensors: Coordinate[] = []
  const beacons: Coordinate[] = []
  const bounds: Bounds = {minX: 0, maxX: 0}
  const noBeacons: Set<number> = new Set()

  splitFile.forEach((line) => {
    const values = line.match(/[\d-]+/g)?.map((value) => Number(value))

    if (values && (values?.filter(value => value) || []).length > 0) {
      sensors.push({x: values[0], y: values[1]})
      beacons.push({x: values[2], y: values[3]})

      bounds.minX = Math.min(bounds.minX, values[0], values[2])
      bounds.maxX = Math.max(bounds.maxX, values[0], values[2])
    }
  })

  for (let s = 0; s < sensors.length; s++) {
    const closestBeaconDistance = calculateManhattanDistance(sensors[s], beacons[s])

    for (let x = bounds.minX; x < bounds.maxX; x++) {
      const manhattanDistance = calculateManhattanDistance(sensors[s], {x, y})
      if (manhattanDistance > 1 && manhattanDistance <= closestBeaconDistance) {
        noBeacons.add(x)
      }
    }
  }

  return noBeacons.size
}

console.log("Possible beacons at y=10: " + main(10))
console.log("Possible beacons at y=2000000: " + main(2000000))