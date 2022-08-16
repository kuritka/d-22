package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]item{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		id := rnd.Intn(7) + 1
		go func(id int) {
			if i, ok := queryCacheFast(id); ok {
				fmt.Println("From cache ", i)
			}
		}(id)

		go func(id int) {
			if i, ok := queryDatabaseSlow(id); ok {
				fmt.Println("From database ", i)
			}
		}(id)

		fmt.Println("item not found")
	}
	time.Sleep(time.Second * 2)
	elapsed := time.Since(start)
	fmt.Printf("\nexecution time %s", elapsed)
}
