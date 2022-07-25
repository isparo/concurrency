package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// channel with a capacity of 10 messages in queue
	exampleChn := make(chan string, 10)

	fmt.Println("Channel size: ", len(exampleChn))

	go func() {
		fmt.Println("readaer 1")

		for message := range exampleChn {
			time.Sleep(2 * time.Second)
			data := fmt.Sprintf("value: %s - buffer size: %d", message, len(exampleChn))
			fmt.Println("(reader 1) ", data)
		}

	}()

	go func() {
		fmt.Println("readaer 2")

		for message := range exampleChn {
			time.Sleep(2 * time.Second)
			data := fmt.Sprintf("value: %s - buffer size: %d", message, len(exampleChn))
			fmt.Println("(reader 2) ", data)
		}

	}()

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		fmt.Println("sender 1")

		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			exampleChn <- fmt.Sprintf("%d", i)
		}

		defer wg.Done()

	}()

	go func() {
		fmt.Println("sender 2")

		for i := 10; i < 20; i++ {
			time.Sleep(1 * time.Second)
			exampleChn <- fmt.Sprintf("%d", i)
		}

		defer wg.Done()
	}()

	go func() {
		fmt.Println("sender 3")

		for i := 20; i < 30; i++ {
			time.Sleep(1 * time.Second)
			exampleChn <- fmt.Sprintf("%d", i)
		}

		defer wg.Done()
	}()

	go func() {
		fmt.Println("sender 4")

		for i := 30; i < 40; i++ {
			time.Sleep(1 * time.Second)
			exampleChn <- fmt.Sprintf("%d", i)
		}

		defer wg.Done()
	}()

	wg.Wait()

}
