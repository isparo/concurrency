package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	exampleChn := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		fmt.Println("preparing to send")

		time.Sleep(3 * time.Second)

		exampleChn <- "value"
	}()

	go func() {
		fmt.Println("waiting for a value")

		val := <-exampleChn

		fmt.Println("Value: ", val)
		wg.Done()
	}()
	wg.Wait()
}
