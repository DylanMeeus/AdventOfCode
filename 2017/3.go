package main

import(
    "fmt"
    "math"
    )




func main(){
    //distance := getDistance(277678)
    solve()
}

type cors struct{
    x int
    y int
}


func solve(){
    // while we have not reached 27768, get next number
    lastnumber := 0
    treshold := 277678
    cormap := make(map[cors]int)

    start := cors{x:0,y:0}
    cormap[start] = 1

    for i := 1; lastnumber <= treshold;i++ {
        xCor := int(math.Floor(getX(i)+.5))
        yCor := int(math.Floor(getY(i)+.5))
        c := cors{x:xCor, y:yCor}
        value := neighbourValues(c,&cormap)
        cormap[c] = value
        lastnumber = value
    }
    fmt.Println(lastnumber)
}

func neighbourValues(cor cors, cormap *map[cors]int) int{
    y := cor.y
    x := cor.x
    sum := 0
    for startx := x-1; startx <= x+1; startx++{
        for starty := y-1; starty <= y+1; starty++{
            dummycor := cors{x:startx, y:starty}
            value, ok := (*cormap)[dummycor]
            if ok {
                sum += value
            }
        }
    }
    return sum
}


func getDistance(n int) float64 {
    x := getX(n)
    y := getY(n)

    // we need to find the path down to X = 0, Y = 0 (sum of coordinates)
    fmt.Print("x: ")
    fmt.Println(x)
    fmt.Print("y: ")
    fmt.Println(y)
    return x + y
}

func getX(n int) float64{
    if n == 1{
    	return 0.0
    }

    result := getX(n-1)
    return result + math.Sin(float64((int64(math.Floor(math.Sqrt(float64(4*(n-2)+1)))) % 4)) * (math.Pi/2))

}

func getY(n int) float64{
    if n == 1{
    	return 0.0
    }

    result := getY(n-1)
    return result + math.Cos(float64((int64(math.Floor(math.Sqrt(float64(4*(n-2)+1)))) % 4)) * (math.Pi/2))
}
