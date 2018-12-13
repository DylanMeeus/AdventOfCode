package main


import (
    "fmt"
    "io/ioutil"
    "strings"
    "sort"
)


var (
    right = point{1,0}
    left = point{-1, 0}
    up = point{0, -1}
    down = point{0,1}
)
type point struct {
    x,y int
}


type car struct {
    position, direction *point
    turn int // 0: left 1: straight, 2: right
}

type junction struct {
    position, direction *point
}


func main() {
    fmt.Printf("%v\n", solve(parse()))
}

func solve(atlas map[point]string, cars []*car) *point {
    for {
        carpos := make(map[point]*car)
        pts := make([]*point,0)
        for _,c := range cars {
            pts = append(pts, c.position)
            carpos[*c.position] = c
        }
        sort.Slice(pts, func(i, j int) bool {
            return pts[i].y < pts[j].y
        })
        sort.Slice(pts, func(i, j int) bool {
            return pts[i].x < pts[j].x
        })
        for _,p := range pts {
            c := carpos[*p]
            c.move(atlas)
            // check the collision
            for _,o := range cars {
                if c != o {
                    if (*c.position) == (*o.position) {
                        return c.position
                    }
                }
            }
        }
    }
    return nil
}


func (c *car) move(atlas map[point]string) {
    c.position.x += c.direction.x
    c.position.y += c.direction.y
    switch atlas[*c.position] {
    case "+":
        // turn left, straight, right
        switch c.turn {
            case 0:
                if (*c.direction) == right {
                    c.direction = &up 
                } else if (*c.direction) == left {
                    c.direction = &down
                } else if (*c.direction) == up {
                    c.direction = &left
                } else if (*c.direction) == down {
                    c.direction = &right
                }
                break
            case 1:
                // nothing changes, ezpz
                break
            case 2:
                if (*c.direction) == right {
                    c.direction = &down
                } else if (*c.direction) == left {
                    c.direction = &up
                } else if (*c.direction) == up {
                    c.direction = &right
                } else if (*c.direction) == down {
                    c.direction = &left
                }
                break
        }
        c.turn++
        if c.turn == 3 {
            c.turn = 0
        }
        break
    case "/":
        if (*c.direction) == right {
            c.direction = &up
        } else if (*c.direction) == left {
            c.direction = &down
        } else if (*c.direction) == up {
            c.direction = &right
        } else if (*c.direction) == down {
            c.direction = &left
        }
        break
    case "\\":
        if (*c.direction) == right {
            c.direction = &down
        } else if (*c.direction) == left {
            c.direction = &up
        } else if (*c.direction) == up {
            c.direction = &left
        } else if (*c.direction) == down {
            c.direction = &right
        }
        break
    }    
}

func parse() (map[point]string, []*car) {
    atlas := make(map[point]string)
    cars := make([]*car,0)
    junctions := make([]*junction,0)
    bytes,_ := ioutil.ReadFile("input.txt")
    for y,part := range strings.Split(string(bytes), "\n") {
        for x,char := range strings.Split(part,"") {
            atlas[point{x,y}] = parseRoad(char)
            if c, ok := parseCar(&point{x,y}, char); ok {
                cars = append(cars,c)
            }
            if j, ok := parseJunction(&point{x,y}, char); ok {
                junctions = append(junctions,j)
            }
        }
    }
    // surpress not used
    _ = junctions
    return atlas, cars
}

func parseJunction(pos *point, s string) (*junction, bool) {
    switch s {
        case "+":
            return &junction{pos, &point{1,0}}, true
    }
    return nil, false
}

func parseCar(pos *point, s string) (*car, bool) {
    switch s {
    case ">":
        return &car{pos, &right,0}, true
    case "<":
        return &car{pos, &left, 0}, true
    case "v":
        return &car{pos, &down, 0}, true
    case "^":
        return &car{pos, &up, 0}, true
    }
    return nil, false 
}

// this data is not used
func parseRoad(s string) string {
    switch s {
    case ">":
        return "-"
    case "<":
        return "-"
    case "v":
        return "|"
    case "^":
        return "|"
    }
    return s
}
