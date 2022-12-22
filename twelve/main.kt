import java.io.File
import java.util.Arrays

data class Coordinate(var x: Int, var y: Int, var h: Int, var d: Char)

var w = 0
var h = 0
var x: MutableList<MutableList<Coordinate>> = ArrayList()

fun pointNotTraversed(route: MutableList<Coordinate>, targetCoordinate: Coordinate): Boolean {
  return route.filter { it.x == targetCoordinate.x && it.y == targetCoordinate.y }.isEmpty()
}

fun walk(coordinate: Coordinate, destination: Coordinate, src: Array<IntArray>, r: MutableList<Coordinate>, dir: Char): MutableList<Coordinate> {
  var route = r.toMutableList()

  if (route.size > 0) {
    route[route.size - 1] = Coordinate(coordinate.x, coordinate.y, coordinate.h, dir)
  }
  route.add(coordinate)

  if (coordinate.x + 1 in 0 until w) { // Right
    val targetCoordinate = Coordinate(coordinate.x + 1, coordinate.y, src[coordinate.y][coordinate.x + 1], "#"[0])
    if (targetCoordinate.h - coordinate.h in 0..1 && pointNotTraversed(route, targetCoordinate) ) {
      walk(targetCoordinate, destination, src, route, ">"[0])
    }
  }
  if (coordinate.x - 1 in 0 until w) { // Left
    val targetCoordinate = Coordinate(coordinate.x - 1, coordinate.y, src[coordinate.y][coordinate.x - 1], "#"[0])
    if (targetCoordinate.h - coordinate.h in 0..1 && pointNotTraversed(route, targetCoordinate)) {
      walk(targetCoordinate, destination, src, route, "<"[0])
    }
  }
  if (coordinate.y + 1 in 0 until h) { // Down
    val targetCoordinate = Coordinate(coordinate.x, coordinate.y + 1, src[coordinate.y + 1][coordinate.x], "#"[0])
    if (targetCoordinate.h - coordinate.h in 0..1 && pointNotTraversed(route, targetCoordinate)) {
      walk(targetCoordinate, destination, src, route, "v"[0])
    }
  }
  if (coordinate.y - 1 in 0 until h) { // Up
    val targetCoordinate = Coordinate(coordinate.x, coordinate.y - 1, src[coordinate.y - 1][coordinate.x], "#"[0])
    if (targetCoordinate.h - coordinate.h in 0..1 && pointNotTraversed(route, targetCoordinate)) {
      walk(targetCoordinate, destination, src, route, "^"[0])
    }
  }

  if (coordinate.x == destination.x && coordinate.y == destination.y) {
    x.add(route)
  }

  return route
}

fun main() {
  val src = File("heightmap.txt").readLines()
  w = src[0].length
  h = src.size
  var s = Coordinate(0, 0, 0, "#"[0])
  var e = Coordinate(0, 0, 0, "#"[0])
  var heights = Array(h, {IntArray(w)})

  for (j in 0 until h) {
    for (i in 0 until w) {
      val height = src[j][i]
      var code = height.code - 96
      if (height == "S"[0]) {
        s = Coordinate(i, j, 1, "#"[0])
        code = 1
      } else if (height == "E"[0]) {
        e = Coordinate(i, j, 26, "#"[0])
        code = 26
      }

      heights[j][i] = code
    }
  }

  walk(s, e, heights, mutableListOf<Coordinate>(), "#"[0])

  println("Paths calculated (in descending order of efficiency): ${x.size}\n")
  val sorted = x.sortedBy { it.size }
  for (route in sorted) {
    println("\nRoute length: ${route.size - 1}")

    var plot = Array(h) { Array(w) { "." } }
    for (point in route) {
      plot[point.y][point.x] = Character.toString(point.d)
    }

    for (line in plot) {
      println(line.joinToString(""))
    }
  }
}