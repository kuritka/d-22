package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"runtime"
	"strconv"
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

func extract() []*Order {
	result := []*Order{}
	f, _ := os.Open("./2015/11/input/orders.txt")
	defer f.Close()
	r := csv.NewReader(f)
	// one record  at time from a file and proceeding as non errors are encountered
	// po jednom záznamu ze souboru a pokračování podle toho, jak se vyskytnou chyby.
	// je to FOR cyklus, ktery bezi do konce souboru!!!
	// something like  for i:=0 ; i< something; i++ {...}
	for record, err := r.Read(); err == nil; record, err = r.Read() {
		order := new(Order)
		order.CustomerNumber, _ = strconv.Atoi(record[1])
		order.PartNumber = record[0]
		order.Quantity, _ = strconv.Atoi(record[2])
		result = append(result, order)
	}
	return result
}

func transform(orders []*Order) []*Order {
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

	for _, o := range orders {
		time.Sleep(3 * time.Millisecond)
		o.UnitPrice = productList[o.PartNumber].UnitPrice
		o.UnitCost = productList[o.PartNumber].UnitCost
	}
	return orders
}

func load(orders []*Order) {
	f, _ := os.Create("./2015/11/dest.txt")
	defer f.Close()
	fmt.Fprintf(f, "%20s%15s%12s%12s%15s%15s\n", "Part Number", "Quantity", "Unit Cost", "Unit Price", "Total Cost", "Total Price")
	for _, o := range orders {
		time.Sleep(1 * time.Millisecond)
		fmt.Fprintf(f, "%20s %15d %12.2f v%12.2f %15.2f %15.2f\n", o.PartNumber, o.Quantity, o.UnitCost, o.UnitPrice, o.UnitCost*float64(o.Quantity), o.UnitPrice*float64(o.Quantity))
	}
}

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(1)
	orders := extract()
	orders = transform(orders)
	load(orders)
	fmt.Printf("\nexecution time %s", time.Since(start))
}
