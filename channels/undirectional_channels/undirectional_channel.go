package main

import "fmt"

func main() {
	var biDirectional chan string
	var readOnly <-chan string

	biDirectional = make(chan string)
	readOnly = biDirectional

	go func() {
		biDirectional <- "hello read only"
		close(biDirectional)
	}()

	fmt.Println(<-readOnly)
}
