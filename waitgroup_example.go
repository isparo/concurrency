package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(3)

	go doSomething(3, "alert 1", &wg)
	go doSomething(1, "alert 2", &wg)
	go doSomething(4, "alert 3", &wg)

	//wg.Wait()

	time.Sleep(time.Second * 10)
	fmt.Println("process completed")
}

func doSomething(num int, name string, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Second * (time.Duration(num)))

	fmt.Println("Goroutine executed: ", name)

}
