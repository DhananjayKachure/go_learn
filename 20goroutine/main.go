package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var signals = []string{"test"}
var mut sync.Mutex

func main() {
	fmt.Println("Let's do Go routines")
	url := []string{
		"https://gobyexample.com/tickers",
		"https://www.w3schools.com/js/js_iterables.asp",
		"https://www.allmath.com/bodmas-calculator.php",
		"https://www.geeksforgeeks.org/creating-a-command-line-calculator-in-golang-step-by-step-tutorial",
	}

	// Loop through URLs and launch goroutines to fetch status
	for i := 0; i < len(url); i++ {
		wg.Add(1)
		go getStatus(url[i])
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Lock signals before reading it

	fmt.Println(signals)
}

func getStatus(endpoint string) {
	defer wg.Done()

	// Fetch the status of the URL
	res, err := http.Get(endpoint)
	if err != nil {
		log.Printf("Error fetching %s: %v", endpoint, err)
		return
	}

	// Lock before appending to shared signals slice
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()

	// Log the status code
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}
