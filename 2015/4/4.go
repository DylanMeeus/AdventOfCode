package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	input     = "yzbqklnj"
	testInput = "abcdef"
)

func main() {
	fmt.Printf("First number to break the code: %v\n", solve("00000")-1)
	fmt.Printf("First number to break the code: %v\n", solve("000000")-1)
}

func solve(prefix string) int {
	var hashed string
	i := 0
	for !strings.HasPrefix(hashed, prefix) {
		testHash := fmt.Sprintf("%v%v", input, i)
		hashed = hash(testHash)
		i += 1
	}
	return i
}

func hash(in string) (out string) {
	hasher := md5.New()
	hasher.Write([]byte(in))
	return hex.EncodeToString(hasher.Sum(nil))
}
