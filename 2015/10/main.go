package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	generate := func(ch chan int) {
		for i := 2; ; i++ {
			ch <- i
		}
	}

	filter := func(in, out chan int, prime int) {
		for {
			i := <-in
			if i%prime != 0 {
				out <- i
			}
		}
	}

	start := time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU())
	ch := make(chan int)
	go generate(ch)
	for {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1 // 🟢
	}

	fmt.Printf("\nexecution time %s", time.Since(start))
}
