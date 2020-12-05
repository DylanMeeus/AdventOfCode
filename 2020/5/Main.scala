package main

import scala.io.Source
import scala.collection.immutable.SortedSet

object Main extends App {


  def toNumber(input: String): Int = {
    val mapping = Map ('F' -> 0, 'B' -> 1, 'L' -> 0, 'R' -> 1)
    var idx = 9
    var out = 0
    for (c <- input) {
      val bin = mapping.getOrElse(c, 0) 
      out += (bin * scala.math.pow(2, idx).toInt)
      idx = idx - 1
    }
    return out
  }

  def solve1(): Int = {
    var maxID = 0
    // read file
    val file = "input.txt"
    for (line <- Source.fromFile(file).getLines) {
      val seatID = toNumber(line)
      if (seatID > maxID) {
        maxID = seatID
      }
    }
    return maxID

  }

  def solve2(max: Int): Int = {
    var seatSet = SortedSet(0)
    val file = "input.txt"
    for (line <- Source.fromFile(file).getLines) {
      val seatID = toNumber(line)
      seatSet = seatSet + seatID
    }

    for (i <- 1 until max) {
      if (seatSet.contains(i - 1) && seatSet.contains(i + 1) && !seatSet.contains(i)) {
        return i
      }
    }
    return 0
  }

  val max = solve1()
  val missing = solve2(max)
  println(max)
  println(missing)
}


