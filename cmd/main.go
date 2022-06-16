package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {

	var tests []TestCase
	wg := sync.WaitGroup{}
	fmt.Println("starting test...")
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			avg, err := startSyncSpam(20)
			if err != nil {
				log.Fatalf("!!ERROR OCCURED DURING TEST CASE!! err - %s", err.Error())
			}
			r := TestCase{
				Elapsed: fmt.Sprintf("%.2f", avg),
			}
			tests = append(tests, r)
			wg.Done()
		}()
	}
	wg.Wait()

	for i, test := range tests {
		fmt.Printf("TEST #%d. Average response latency time - : %s \n", i+1, test.Elapsed)
	}
	fmt.Println("success")
	return
}

type TestCase struct {
	Elapsed string
}

func startSyncSpam(times int) (float64, error) {
	now := time.Now()
	defer func() {
		elapsed := time.Since(now)
		fmt.Printf("total elapsed time  testFn(sync type) - %.2f \n", elapsed.Seconds())
	}()

	var latencies []float64
	sumOfLatency := 0.0

	for i := 0; i < times; i++ {
		lat, err := doRequest()
		if err != nil {
			return 0.0, err
		}
		sumOfLatency += lat
		latencies = append(latencies, lat)
	}

	avgLatencyTime := sumOfLatency / float64(len(latencies))

	return avgLatencyTime, nil

}

func doRequest() (float64, error) {

	url := "https://storage.yandexcloud.net/zharpizza-bucket/static/images/3.png"

	now := time.Now()
	_, err := http.Get(url)
	if err != nil {
		return 0.0, fmt.Errorf("error occured sending get request. err - %s", err.Error())
	}
	elapsed := time.Since(now)

	return elapsed.Seconds(), nil
}
