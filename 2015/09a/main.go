package main

import (
	"fmt"
	"time"
)

type PurchaseOrder struct {
	Number int
	Value  float64
}

func main() {
	consumer := func(results <-chan *PurchaseOrder) {
		for result := range results {
			fmt.Printf("Received: %v\n", result)
		}
		fmt.Println("Done receiving!")
	}

	savePO := func(po *PurchaseOrder, callback chan *PurchaseOrder) {
		fmt.Println("Purchase ordered, making web service call...")
		time.Sleep(1 * time.Second)
		callback <- po
	}

	publisher := func(data []*PurchaseOrder) <-chan *PurchaseOrder {
		results := make(chan *PurchaseOrder, len(data))
		defer close(results) // 🟢
		ch := make(chan *PurchaseOrder, len(data))
		for _, v := range data {
			// long running task
			go savePO(v, ch)
		}
		for i := 0; i < len(data); i++ { // 🟢
			results <- <-ch
		}
		return results
	}

	start := time.Now()
	result := publisher([]*PurchaseOrder{{1234, 50.12}, {1235, 13.1}, {1236, 11.1}})
	consumer(result)
	fmt.Printf("\nexecution time %s", time.Since(start))
}
