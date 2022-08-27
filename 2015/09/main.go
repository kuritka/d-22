package main

import (
	"fmt"
	"time"
)

type PurchaseOrder struct {
	Number int
	Value  float64
}

func SavePO(po *PurchaseOrder, callback chan *PurchaseOrder) {
	fmt.Println("Purchase ordered, making web service call...")
	time.Sleep(1 * time.Second)
	callback <- po
}

func main() {
	start := time.Now()

	chanOwner := func(orders []*PurchaseOrder) <-chan *PurchaseOrder {
		results := make(chan *PurchaseOrder, len(orders))
		go func() {
			defer close(results)
			ch := make(chan *PurchaseOrder, len(orders))
			for _, v := range orders {
				go SavePO(v, ch)
			}
			for i := 0; i < len(orders); i++ {
				results <- <-ch
			}
		}()
		return results
	}

	consumer := func(results <-chan *PurchaseOrder) {
		for result := range results {
			fmt.Printf("Received: %v\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner([]*PurchaseOrder{{1234, 50.12}, {1235, 13.1}, &PurchaseOrder{1236, 11.1}})
	consumer(results)

	fmt.Printf("\nexecution time %s", time.Since(start))
}
