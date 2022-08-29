package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Product struct {
	PartNumber string
	UnitCost   float64
	UnitPrice  float64
}

type Order struct {
	CustomerNumber int
	PartNumber     string
	Quantity       int
	UnitCost       float64
	UnitPrice      float64
}

func extract(ch chan *Order) {
	f, _ := os.Open("./2015/11/input/orders.txt")
	defer f.Close()
	r := csv.NewReader(f)
	for record, err := r.Read(); err == nil; record, err = r.Read() {
		order := new(Order)
		order.CustomerNumber, _ = strconv.Atoi(record[1])
		order.PartNumber = record[0]
		order.Quantity, _ = strconv.Atoi(record[2])
		ch <- order // 🟦 each record goes to channe;
	}
	close(ch) // 🟦 we need to stop ange process in transform actor
}

func transform(inputCh, outputCh chan *Order) {
	f, _ := os.Open("./2015/11/input/orders.txt")
	//https://www.joeshaw.org/dont-defer-close-on-writable-files/
	defer f.Close()

	r := csv.NewReader(f)
	records, _ := r.ReadAll()
	productList := make(map[string]*Product)
	for _, record := range records {
		product := new(Product)
		product.PartNumber = record[0]
		product.UnitCost, _ = strconv.ParseFloat(record[1], 64)
		product.UnitPrice, _ = strconv.ParseFloat(record[2], 64)
		productList[product.PartNumber] = product
	}

	wg := sync.WaitGroup{}
	for o := range inputCh { // 🟦
		wg.Add(1)
		go func(wg *sync.WaitGroup, o *Order) {
			time.Sleep(3 * time.Millisecond)
			o.UnitPrice = productList[o.PartNumber].UnitPrice
			o.UnitCost = productList[o.PartNumber].UnitCost
			outputCh <- o // 🟦 isolated each actor
			wg.Done()
		}(&wg, o)
	}
	wg.Wait()
	close(outputCh) // 🟦 close it to inform range in load actor
}

func load(orders chan *Order, done chan bool) {
	f, _ := os.Create("./2015/11/dest.txt")
	defer f.Close()
	fmt.Fprintf(f, "%20s%15s%12s%12s%15s%15s\n", "Part Number", "Quantity", "Unit Cost", "Unit Price", "Total Cost", "Total Price")
	wg := sync.WaitGroup{}
	for o := range orders {
		wg.Add(1)
		go func(wg *sync.WaitGroup, o *Order) {
			time.Sleep(1 * time.Millisecond)
			fmt.Fprintf(f, "%20s %15d %12.2f %12.2f %15.2f %15.2f\n", o.PartNumber, o.Quantity, o.UnitCost, o.UnitPrice, o.UnitCost*float64(o.Quantity), o.UnitPrice*float64(o.Quantity))
			wg.Done()
		}(&wg, o)
	}
	wg.Wait()
	done <- true
}

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU())
	ch1 := make(chan *Order)
	ch2 := make(chan *Order)
	done := make(chan bool)
	go extract(ch1)
	go transform(ch1, ch2)
	go load(ch2, done)
	<-done
	fmt.Printf("\nexecution time %s", time.Since(start))
}
