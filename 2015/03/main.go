package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// watched patch starts from binary path
const watchedPatch = "./2015/03/source"

func main() {
	fmt.Println("started")
	runtime.GOMAXPROCS(runtime.NumCPU())
	for {
		d, _ := os.Open(watchedPatch)
		files, _ := d.Readdir(-1)
		for _, fi := range files {
			filePath := watchedPatch + "/" + fi.Name()
			f, _ := os.Open(filePath)
			data, _ := ioutil.ReadAll(f)
			_ = f.Close()
			_ = os.Remove(filePath)

			go func(data string) {
				//string to data reader
				reader := csv.NewReader(strings.NewReader(data))
				records, _ := reader.ReadAll()
				for _, r := range records {
					invoice := new(invoice)
					invoice.Number = r[0]
					invoice.Amount, _ = strconv.ParseFloat(r[1], 64)
					invoice.OrderNum, _ = strconv.Atoi(r[3])
					unixTime, _ := strconv.ParseInt(r[3], 10, 64)
					invoice.Date = time.Unix(unixTime, 0)
					fmt.Printf("invoice received %v \n", invoice)
				}
			}(string(data))
		}
	}

}

type invoice struct {
	Number   string
	Amount   float64
	OrderNum int
	Date     time.Time
}
