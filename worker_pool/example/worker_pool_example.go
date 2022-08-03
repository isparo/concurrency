package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type city struct {
	name     string
	location string
}

func createCity(record city) {
	time.Sleep(10 * time.Millisecond)
}

func readData(cityChn chan []city) {
	var cities []city

	csvFile, err := os.Open("cities.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		cities = append(cities, city{
			name:     line[0],
			location: line[1],
		})
	}

	cityChn <- cities

}

func worker(cityChn chan city) {
	for val := range cityChn {
		createCity(val)
	}
}

func main() {
	startTime := time.Now()

	cities := make(chan []city)

	go readData(cities)

	const workers = 5
	jobs := make(chan city, 1000)

	for w := 1; w <= workers; w++ {
		go worker(jobs)
	}

	counter := 0
	for _, val := range <-cities {
		counter++
		jobs <- val
	}

	fmt.Println("records saved: ", counter)
	fmt.Println("total time: ", time.Since(startTime))
}
