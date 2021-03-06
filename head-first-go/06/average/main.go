// B"H

// average calculates the average of several numbers.
package main

import (
	"fmt"
	"log"

	"github.com/Ylazerson/go-shenanigans/head-first-go/06/datafile"
)

// -- ------------------------------------------
// THIS IS OBVIOUSLY NOT THE CORRECT APPRAOCH TO USE
// ... but hey for now .....
const filePath = "/home/laz/tmp/ch06-data.txt"

// -- ------------------------------------------

func main() {

	// -- --------------------------------------
	numbers, err := datafile.GetFloats(filePath)

	if err != nil {
		log.Fatal(err)
	}
	// -- --------------------------------------

	// -- --------------------------------------
	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	sampleCount := float64(len(numbers))

	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
	// -- --------------------------------------

}
