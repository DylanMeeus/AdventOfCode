package main

import( 
    "math"
    "io/ioutil"
    "strings"
    "unicode"
    "strconv"
    "image"
    "image/png"
    "image/color"
    "os"
)

type point struct {
    x,y int
}

type star struct {
    position, velocity *point
}

var palette = []color.Color{color.Black, color.RGBA{0xff,0xff,0x00,0xff}}
func main() {
    draw(stars())
}

func starsAlligned(stars []*star) int {
    var distance int
    for _,s := range stars {
        for _,o := range stars {
            if s == o {
                continue
            }
            d := math.Abs(float64(s.position.x - o.position.x)) + math.Abs(float64(s.position.y - o.position.y))
            distance += int(d)
        }
    }
    return distance
}

func draw(stars []*star) {
    var i int
    prev := starsAlligned(stars)
    for true { 
        for _,s := range stars {
            s.move()
        }
        newd := starsAlligned(stars)
        if newd > prev {
            for _,s := range stars {
                s.back()
            }
            break
        }
        prev = newd
        i++
    }
    rect := image.Rect(0,0,1000,1000)
    img := image.NewPaletted(rect, palette)
    for _,s := range stars {
        img.SetColorIndex(s.position.x, s.position.y, 1)
    }
    f, err := os.Create("images/img" + strconv.Itoa(i))
    if err != nil {
        panic(err)
    }
    if err := png.Encode(f, img); err != nil{
        panic(err)
    }
}

func (s *star) move() {
    s.position.x += s.velocity.x
    s.position.y += s.velocity.y
}

func (s *star) back() {
    s.position.x -= s.velocity.x
    s.position.y -= s.velocity.y
}

func stars() []*star {
    bytes, err := ioutil.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    sts := make([]*star, 0)
    for _,parts := range strings.Split(string(bytes), "\n") {
        if parts == "" {
            continue
        }
        mapped := strings.Map(func(r rune) rune {
           if unicode.IsDigit(r) || r == rune('-') || r == rune(',') {
               return r
           } 
           if r == rune('>') {
               return rune(' ')
           }
           return -1 
        }, parts)
        posvel := strings.Split(mapped, " ")
        pos := strings.Split(posvel[0], ",")
        vel := strings.Split(posvel[1], ",")
        posx,_ := strconv.Atoi(pos[0])
        posy,_ := strconv.Atoi(pos[1])
        velx,_ := strconv.Atoi(vel[0])
        vely,_ := strconv.Atoi(vel[1])
        sts = append(sts, &star{&point{posx,posy}, &point{velx, vely}})
    }
    return sts
}
