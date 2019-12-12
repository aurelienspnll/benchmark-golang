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
)

func benchmark(concurrencyLimit int) time.Duration {
	start := time.Now()
	const size = 1000
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
