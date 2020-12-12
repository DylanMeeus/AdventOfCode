package main

import scala.io.Source
import scala.collection.mutable.ListBuffer
import scala.math.abs

class Instruction(dir: String, amount: Int) {
  var direction: String = dir
  var value: Int = amount 
}

class Ship() {
  var x: Int = 0
  var y: Int = 0
  var direction: Int = 1

  val directions = Map(0 -> "N",1 -> "E",2 -> "S",3 -> "W")
  
  def process(input: Instruction) {
    if (input.direction == "R" || input.direction == "L") {
      if (input.direction == "L") {
        // transform to a right-turn
        input.value = 360 - input.value
      }
      turn(input.value)
    } else {
      move(input)
    }

  }

  def turn(degrees: Int) {
    val amount = degrees / 90
    direction = (direction + amount) % 4
  }


  def move(in: Instruction) {
    if (in.direction == "F") {
      in.direction = directions(direction)
    }

    in.direction match {
      case "N" => y = y - in.value
      case "E" => x = x + in.value
      case "S" => y = y + in.value 
      case "W" => x = x - in.value
    }
  }

  def manhattan(): Int = {
    return abs(x) + abs(y)
  }

}

object Main extends App {

  def getInput(): List[Instruction] = { 

    var instructions = new ListBuffer[Instruction]()
    val file = "input.txt"
    for (line <- Source.fromFile(file).getLines) {
      if (line != "") {
        var dir = line(0).toString
        var value = line.drop(1).toInt
        instructions += new Instruction(dir, value)
      }
    }
    return instructions.toList
  }


  def solve1(): Int = {
    var ship = new Ship();
    val instructions = getInput()

    for (in <- instructions) {
      ship.process(in)
    }

    // calculate manhattan distance
    return ship.manhattan()
  }

  println(solve1())
  

}
