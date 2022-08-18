package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"sync"
	"time"
)

// {"status":"success","data":{"id":2,"employee_name":"Garrett Winters","employee_salary":170750,"employee_age":63,"profile_image":""},"message":"Successfully! Record has been fetched."}

type empl struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID     int    `json:"id"`
		Name   string `json:"employee_name"`
		Salary int    `json:"employee_salary"`
		Age    int    `json:"employee_age"`
	} `json:"data"`
}

func main() {
	runtime.GOMAXPROCS(2)
	start := time.Now()
	wg := sync.WaitGroup{}
	for i := 5; i <= 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, id int) {
			resp, _ := http.Get(fmt.Sprintf("https://dummy.restapiexample.com/api/v1/employee/%d", id))
			defer func() { _ = resp.Body.Close() }()
			body, _ := ioutil.ReadAll(resp.Body)
			e := empl{}
			_ = json.Unmarshal(body, &e)
			fmt.Println(e.Data.Name, " ", e.Data.Age)
			wg.Done()
		}(&wg, i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("\nexecution time %s", elapsed)
}

//Rhona Davidson   55
//Colleen Hurst   39
//Herrod Chandler   59
//Sonya Frost   23
//Airi Satou   33
//Brielle Williamson   61
//
//execution time 1.196298925s
