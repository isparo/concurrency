package main

import (
	"fmt"
	"sync"
)

// shared variable
var ordersCount int

func chefPrepareOrder(lock *sync.Mutex, restaurant *sync.WaitGroup) {
	lock.Lock()
	fmt.Println("chef is preparing order")

	ordersCount++

	defer func() {
		restaurant.Done()
		lock.Unlock()
	}()
}

func waiterTakeOrder(lock *sync.Mutex, restaurant *sync.WaitGroup) {
	lock.Lock()
	fmt.Println("waiter is taking order")

	if ordersCount > 0 {
		ordersCount--
		fmt.Println("taking order: ", ordersCount)
	}

	defer func() {
		restaurant.Done()
		lock.Unlock()
	}()
}

func main() {

	var lock sync.Mutex

	var restaurant sync.WaitGroup

	for i := 0; i < 20; i++ {
		restaurant.Add(1)
		go chefPrepareOrder(&lock, &restaurant)
	}

	for i := 0; i < 10; i++ {
		restaurant.Add(1)
		go waiterTakeOrder(&lock, &restaurant)
	}

	restaurant.Wait()

}
