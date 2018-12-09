package main

import(
    "fmt"
    "./tools"
)

func main(){
    solve()
}

func solve(){
    // build the list
    l := tools.List{}
    currentIndex := 1
    // always place a marble between current / current + 1

    l.Add(0)
    players := 470 
    pScores := make(map[int]int, 0)
    currentPlayer := 0
    for marble := 1; marble <= 72170; marble++{ 
        if marble % 23 == 0 {
            pScores[currentPlayer] += marble 
            previous := currentIndex - 7 - 1
            if previous < 0 {
                diff := 0 - previous
                previous = l.Size - diff 
            }
            n := l.Get(previous)
            pScores[currentPlayer] += n.Val
            l.RemoveAt(previous)
            currentIndex = (previous + 1) % l.Size
            //fmt.Printf("%v\n", currentIndex)
            //fmt.Printf("%v played %v\n", currentPlayer, l.String())
            currentPlayer = (currentPlayer+1) % players
            continue
        }
        if currentIndex == l.Size {
            l.AddAt(marble, 1)
            currentIndex = 2
        } else {
            l.AddAt(marble, currentIndex + 1)
            currentIndex += 2
        }
       // fmt.Printf("%v played %v\n", currentPlayer, l.String())
        currentPlayer = (currentPlayer+1)%players
    }

    fmt.Printf("%v\n", l.String())
    var max int
    for _, v := range pScores {
        if v > max {
            max = v
        }
    }
    fmt.Printf("%v\n", max)
}

