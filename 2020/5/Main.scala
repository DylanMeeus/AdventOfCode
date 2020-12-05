package main

import scala.io.Source

object Main extends App {
  def findLocation(input: String): Int  = {
    // first four characters to take location
    // last characters to take seat

    val first: String = input.slice(0,7)
    val last: String = input.slice(7, 10)


    var lo = 0
    var hi = 127
    var mid = 0

    for (c <- first) {
      mid = lo + ((hi - lo) / 2)
      c match {
        case 'F' => hi = mid
        case 'B' => lo = mid
      }
    }

    val row = mid
    println(row)
    lo = 0
    hi = 8
    mid = 0


    for (c <- last) {
      mid = lo + ((hi - lo) / 2)
      c match {
        case 'L' => hi = mid
        case 'R' => lo = mid
      }
    }

    val col = mid

    

    return (row * 8) + col
  }


  //println(findLocation("FBFBBFFRLR"))

  var maxID = 0
  // read file
  val file = "input.txt"
  for (line <- Source.fromFile(file).getLines) {
    val seatID = findLocation(line)
    if (seatID > maxID) {
      maxID = seatID
    }
  }
  println(maxID)



}


