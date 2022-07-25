package main

import (
	"fmt"
	"time"
)

func mainw() {
	fmt.Println("Init...")

	// calling a function
	go sendAlert()

	//using an anonymous function
	go func() {
		fmt.Println("send alert")
	}()

	time.Sleep(time.Second * 3)

	fmt.Println("End...")
}

func sendAlert() {
	fmt.Println("send alert")
}
