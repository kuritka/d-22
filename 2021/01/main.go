package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(3)
	cdur, _ := time.ParseDuration("10ms")
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Hello")
			time.Sleep(cdur)
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("GO")
			time.Sleep(cdur)
		}
	}()
	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}

//Hello
//GO
//GO
//Hello
//HelloGO
//
//Hello
//Hello
//Hello
//GO
//Hello
//GO
