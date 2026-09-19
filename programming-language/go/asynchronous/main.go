package main

import (
	"fmt"
	"math/rand" // generates random value lib
	"sync"
	"time"
)

func getConfigure(s string) string {
	// `Intn` is `math/rand` in the package of random integer generator and value bind
	// time.Duration(rand.Intn(n)) * time.<Clocktype>
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    return fmt.Sprintf("Connected to %s", s)
}

func connectAsync(devices []string, result chan string) {
	var wg sync.WaitGroup

	for _, device := range devices {
		wg.Add(1)
		go func(d string) {
			defer wg.Done()
			result <- getConfigure(d)
		}(device)
	}

	wg.Wait()
	close(result)
}

func main() {
	// array or slice literal syntax: []<DataType>{<element1>, <element2>, ...}
	devices := []string{"Laptop", "Smartphone", "Tablet"}
	
	// `make` is used to create a channel with a specified buffer size
	// make(chan <DataType>, <bufferSize>)
	resultCh := make(chan string, len(devices))

	go connectAsync(devices, resultCh)

	fmt.Println(<-resultCh)

	for msg := range resultCh {
		fmt.Println(msg)
	}
}