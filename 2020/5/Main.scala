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


  var maxID = 0
  var seatSet = SortedSet(0)
  // read file
  val file = "input.txt"
  for (line <- Source.fromFile(file).getLines) {
    val seatID = toNumber(line)
    seatSet = seatSet + seatID
    if (seatID > maxID) {
      maxID = seatID
    }
  }
  println(maxID)
}


