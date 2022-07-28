package main

import (
	"fmt"
	"time"
)

func main() {

	chan1 := make(chan string)
	chan2 := make(chan string)
	quit := make(chan bool)

	go func() {
		time.Sleep(3 * time.Second)
		chan1 <- "hello 1"
	}()

	go func() {
		time.Sleep(4 * time.Second)
		chan2 <- "hello 2"
		close(quit)
	}()

	for {
		select {
		case <-chan1:
			// do something
			fmt.Println("from chan 1")
		case <-chan2:
			// do other thing
			fmt.Println("from chan 2")

		case <-quit:
			return
		}
	}
}
