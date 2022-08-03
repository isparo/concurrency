package main

import (
	"fmt"
	"time"
)

func doProcessOne(initChn chan string, pross1Chan chan string) {
	time.Sleep(150 * time.Millisecond)
	pross1Chan <- concatStage(<-initChn, "first step")
	close(pross1Chan)
}

func doProcessTwo(pross1Chan chan string, pross2Chan chan string) {
	time.Sleep(200 * time.Millisecond)
	pross2Chan <- concatStage(<-pross1Chan, "second  step")
	close(pross2Chan)
}

func doProcessThree(pross2Chan chan string, pross3Chan chan string) {
	time.Sleep(20 * time.Millisecond)
	pross3Chan <- concatStage(<-pross2Chan, "third step")
	close(pross3Chan)
}

func doProcessFour(pross3Chan chan string, resultChn chan string) {
	time.Sleep(98 * time.Millisecond)
	resultChn <- concatStage(<-pross3Chan, "result step")
	close(resultChn)
}

func requestValue(initChn chan string) {
	time.Sleep(300 * time.Millisecond)
	initChn <- "start point"
}

func concatStage(current, next string) string {
	return current + " -> " + next
}

func main() {
	startTime := time.Now()

	fmt.Println("staring process")

	initChn := make(chan string)
	pross1Chan := make(chan string)
	pross2Chan := make(chan string)
	pross3Chan := make(chan string)
	resultChn := make(chan string)

	go doProcessOne(initChn, pross1Chan)
	go doProcessTwo(pross1Chan, pross2Chan)
	go doProcessThree(pross2Chan, pross3Chan)
	go doProcessFour(pross3Chan, resultChn)

	requestValue(initChn)
	close(initChn)

	result := <-resultChn
	fmt.Println("the result is: ", result)
	fmt.Println("total time: ", time.Since(startTime))
}
