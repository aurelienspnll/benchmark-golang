package main

// func checkUrl(url string, res []string, c chan string) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return
// 	}
// 	bodyBytes, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// bodyString := string(bodyBytes)
// 	c <- string(bodyBytes)
// 	res = append(res, <-c)
// }

//-----------------------------------------------------------------
//----------------------------- WORKS -----------------------------
//-----------------------------------------------------------------
// func main() {
// 	start := time.Now()
// 	var size = 1000
// 	// c := make(chan string, size)
// 	var wg sync.WaitGroup
// 	wg.Add(size)
// 	var result []string
// 	for i := 0; i < size; i++ {
// 		j := (i % 100) + 1
// 		s := strconv.Itoa(j)
// 		url := "https://jsonplaceholder.typicode.com/todos/" + s
// 		// go checkUrl(url, &result)
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
// 			defer wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println(result[len(result)-1])
// 	fmt.Print("Temps : ")
// 	fmt.Println(time.Since(start))
// }
//-----------------------------------------------------------------
//-----------------------------------------------------------------
//-----------------------------------------------------------------

// func checkUrl(url string, res *[]string) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return
// 	}
// 	bodyBytes, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("hello")
// 	// bodyString := string(bodyBytes)
// 	body := string(bodyBytes)
// 	*res = append(*res, body)
// }
