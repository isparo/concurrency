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

func main() {
	startTime := time.Now()

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

	counter := 0
	for _, line := range csvLines {
		counter++
		createCity(city{
			name:     line[0],
			location: line[1],
		})
	}

	fmt.Println("records saved: ", counter)
	fmt.Println("total time: ", time.Since(startTime))
}
