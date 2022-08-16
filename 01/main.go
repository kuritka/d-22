package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Item{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 10; i++ {
		id := rnd.Intn(7) + 1
		if i, ok := queryCacheFast(id); ok {
			fmt.Println("From cache ", i)
			continue
		}
		if i, ok := queryDatabaseSlow(id); ok {
			fmt.Println("From database ", i)
			continue
		}
		fmt.Println("Item not found")
		time.Sleep(150 * time.Millisecond)
	}
}
