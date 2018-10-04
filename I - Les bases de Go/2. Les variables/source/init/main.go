package main

import (
	"fmt"
	"time"
)

var (
	timeS int64
)

func init() {
	fmt.Println(timeS)
	timeS = time.Now().Unix()
}

func main() {
	fmt.Println(timeS)
}
