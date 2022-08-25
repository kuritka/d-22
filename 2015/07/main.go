package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(4)
	f, _ := os.Create("./log.txt")
	defer f.Close()
	logCh := make(chan string, 50)
	go func() {
		for msg := range logCh {
			logTime := time.Now().Format(time.RFC3339)
			f.WriteString(logTime + " - " + msg + "\n")
		}
	}()

	mutex := make(chan bool, 1)
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex <- true
			go func() {
				msg := fmt.Sprintf("%d + %d = %d", i, j, i+j)
				logCh <- msg
				fmt.Println(msg)
				<-mutex
			}()
		}
	}
	fmt.Printf("\nexecution time %s", time.Since(start))
	fmt.Scanln()
}
