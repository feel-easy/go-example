package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(3))
}
