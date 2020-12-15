package main

import scala.io.Source
import scala.collection.mutable.ListBuffer

object  Main extends App {
  def next(in: List[Int]): Int = {
    val last = in.last


    if (in.count(_ == last) == 1) {
      // first time spoken
      return 0
    }

    // else find index

    return in.reverse.drop(1).takeWhile(_ != last).length + 1
  }

  def solve1(): Int = {
    var numbers = ListBuffer[Int](2,0,6,12,1,3)
    //var numbers = ListBuffer[Int](0,3,6)

    var iter = numbers.toList.length
    while (true) {
      val n = next(numbers.toList)
      numbers += n
      iter = iter + 1
      if (iter == 2020) {
        return n
      }

    }

    return 0
  }

  println(solve1())
}
