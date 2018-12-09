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
    for marble := 1; marble < 8; marble++{ 
        fmt.Printf("%v\n", l.String())
        fmt.Println(marble)
        fmt.Println(currentIndex)
        fmt.Println(l.Size)
        if currentIndex == l.Size {
            l.AddAt(marble, 1)
            currentIndex = 2
        } else {
            l.AddAt(marble, currentIndex + 1)
            currentIndex += 2
        }
    }
    fmt.Printf("%v\n", l.String())
}

