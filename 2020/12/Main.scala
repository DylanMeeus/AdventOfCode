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


class ShipWaypoint() {
  var waypoint_x = 10
  var waypoint_y = 1
  var ship_x: Int = 0
  var ship_y: Int = 0
  var waypoint_direction: Int = 1
  val directions = Map(0 -> "N",1 -> "E",2 -> "S",3 -> "W")
  
  def process(input: Instruction) {
    if (input.direction == "R" || input.direction == "L") {
      if (input.direction == "L") {
        // transform to a right-turn
        input.value = 360 - input.value
      }
      turn(input.value)
    } else if (input.direction == "F") {
      moveShip(input.value)
    } else {
      moveWaypoint(input)
    }
  }

  def moveShip(times: Int) {
    ship_x = ship_x + (waypoint_x * times)
    ship_y = ship_y + (waypoint_y * times)
    printf("ship: %d %d\n", ship_x, ship_y)
  }

  def turn(degrees: Int) {
    // this should now manipulate the waypoint by turning X times right around the ship
    val amount = degrees / 90
    waypoint_direction = (waypoint_direction + amount) % 4
    amount match {
      case 1 => {
        val tmp: Int = waypoint_x
        waypoint_x = waypoint_y
        waypoint_y = -tmp
      }
      case 2 => {
        waypoint_x = -waypoint_x
        waypoint_y = -waypoint_y
      }
      case 3 => {
        val tmp: Int = waypoint_x
        waypoint_x = -waypoint_y
        waypoint_y = tmp
      }
      case default => println("nothing matches")
    }
  }


  def moveWaypoint(in: Instruction) {
    in.direction match {
      case "N" => waypoint_y = waypoint_y + in.value
      case "E" => waypoint_x = waypoint_x + in.value
      case "S" => waypoint_y = waypoint_y - in.value 
      case "W" => waypoint_x = waypoint_x - in.value
    }
    printf("%d %d\n", waypoint_x, waypoint_y)
  }

  def manhattan(): Int = {
    return abs(ship_x) + abs(ship_y)
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

  def solve2(): Int = {
    var ship = new ShipWaypoint();
    val instructions = getInput()

    for (in <- instructions) {
      ship.process(in)
    }

    // calculate manhattan distance
    return ship.manhattan()
  }


  println(solve1())
  println(solve2())
  

}
