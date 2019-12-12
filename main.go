package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	// "github.com/shirou/gopsutil/mem"
)

const size int = 1000

func callJsonPlaceHolderApiWithChan(concurrencyLimit int) {
	//const concurrencyLimit = 4
	semaphoreChan := make(chan struct{}, concurrencyLimit)
	resultsChan := make(chan string)
	var result []string
	for i := 0; i < size; i++ {
		j := (i % 100) + 1
		s := strconv.Itoa(j)
		url := "https://jsonplaceholder.typicode.com/todos/" + s
		// go checkUrl(url, result, c)
		// result = append(result, <-c)
		go func() {
			semaphoreChan <- struct{}{}
			resp, err := http.Get(url)
			if err != nil {
				return
			}
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			// bodyString := string(bodyBytes)
			resultsChan <- string(bodyBytes)
			<-semaphoreChan
		}()
	}
	for j := 0; j < size; j++ {
		r := <-resultsChan
		result = append(result, r)
	}
	fmt.Println(len(result))
}

func benchmark(concurrencyLimit int) time.Duration {
	start := time.Now()
	// v, _ := mem.VirtualMemory()

	callJsonPlaceHolderApiWithChan(concurrencyLimit)

	// callJsonPlaceHolderApiWithWG(concurrencyLimit)

	// fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	return time.Since(start)
}

func main() {
	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Comma = ';'
	writer.Write([]string{"Concurrency limit", "Time (ms)"})
	for i := 2; i < 201; i++ {
		t, _ := time.ParseDuration(benchmark(i).String())
		fmt.Printf("With %d as concurrency limit ; Time = %d milliseconds\n", i, t.Milliseconds())
		data := []string{strconv.Itoa(i), strconv.FormatInt(t.Milliseconds(), 10)}
		writer.Write(data)
	}
	// fmt.Print(dataSet)
	// err := writer.Write(data)
	checkError("Cannot write to file", err)
	writer.Flush()
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

// /*
// If we only use WaitGroup, we will overload our processors and never have any returned values.
// TODO: add a chan to manage the concurrency limit.
// */
// func callJsonPlaceHolderApiWithWG(concurrencyLimit int) {
// 	// c := make(chan string, size)
// 	var wg sync.WaitGroup
// 	//wg.Add(size)
// 	var result []string
// 	for i := 0; i < size; i++ {
// 		j := (i % 100) + 1
// 		s := strconv.Itoa(j)
// 		url := "https://jsonplaceholder.typicode.com/todos/" + s
// 		// go checkUrl(url, &result)
// 		wg.Add(1)
// 		go func() {
// 			resp, err := http.Get(url)
// 			if err != nil {
// 				return
// 			}
// 			bodyBytes, err := ioutil.ReadAll(resp.Body)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			//fmt.Println("hello")
// 			// bodyString := string(bodyBytes)
// 			body := string(bodyBytes)
// 			result = append(result, body)
// 			wg.Done() // Don't use defer here
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println(len(result))
// }
