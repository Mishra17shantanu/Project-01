package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"

	"fmt"
)

func checkAndSavebody(url string, wg *sync.WaitGroup) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Println("URL is down\n", url)

	} else {
		defer resp.Body.Close()
		fmt.Printf("%s -> Status code: %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1]
			file += ".txt"
			fmt.Printf("writing Response body %s\n", file)
			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)

			}

		}
	}
	wg.Done()

}

func main() {

	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}

	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go checkAndSavebody(url, &wg)
		fmt.Println(strings.Repeat("#", 20))

	}
	fmt.Println("No of Goroutines", runtime.NumGoroutine())
	wg.Wait()
}
