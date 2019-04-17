package main


import (
    "fmt"
    "strings"
    "crypto/md5"
    "encoding/hex"
)

const (
    input = "yzbqklnj"
    testInput = "abcdef"
)

func main(){
    var hashed string
    i := 0
    for !strings.HasPrefix(hashed,"00000") {
        testHash := fmt.Sprintf("%v%v",input,i)
        hashed = hash(testHash)
        i += 1
    }
    fmt.Printf("First number to break the code: %v\n", i - 1)
}

func hash(in string) (out string) {
    hasher := md5.New()
    hasher.Write([]byte(in))
    return hex.EncodeToString(hasher.Sum(nil))
}
