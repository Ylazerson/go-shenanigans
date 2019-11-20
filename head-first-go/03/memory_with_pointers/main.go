// B''H

// IGNORE THIS CODE FOR NOW - NEEDS FUTHER RESEARCH

// Not so simple to easily compare pointers vs non-pointers with strings in Go
// beacuse of the way strings are stored behind the scenes.

package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {

	var x, _ = new(big.Int).SetString("21888242871839275222246405745257275088548364400416034343698204186575808495617", 10)
	fmt.Println(x)

	// -- --------------------------------------
	// Create a large integer:
	var largeInt int64 = 1234567890111124444

	// -- --------------------------------------
	// Create array to hold 100000000 integers:
	var intArray [100000000]*int64

	// -- --------------------------------------
	// Set each element in the array to the &largeInt
	for index, _ := range intArray {
		intArray[index] = &largeInt
	}

	// -- --------------------------------------
	// Sleep 10 seconds so you can clearly observe the memory usage:
	time.Sleep(10 * time.Second)

}
