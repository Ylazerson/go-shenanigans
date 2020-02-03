// B"H

package main

import (
	"fmt"
	"log"

	"github.com/Ylazerson/go-shenanigans/head-first-go/07/datafile"
)

// -- ------------------------------------------
// THIS IS OBVIOUSLY NOT THE CORRECT APPRAOCH TO USE
// ... but hey for now .....
const filePath = "/home/laz/tmp/ch07-votes.txt"

func main() {

	// -- --------------------------------------
	// Make-beleive config settings:
	configSettings := make(map[string]string)
	configSettings["username"] = "barky"
	configSettings["password"] = "please1234"
	configSettings["server-ip"] = "127.0.0.1"

	configSettings["username"] = "zooman"

	fmt.Println(configSettings)

	fmt.Println(configSettings["username"])
	fmt.Println(configSettings["password"])
	fmt.Println(configSettings["server-ip"])

	// -- --------------------------------------
	lines, err := datafile.GetStrings(filePath)

	if err != nil {
		log.Fatal(err)
	}
	// -- --------------------------------------

	// -- --------------------------------------
	counts := make(map[string]int)

	for _, line := range lines {
		counts[line]++
	}

	// Print the entire map at once:
	fmt.Println(counts)

	// Print the map one key/value at a time:
	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
	// -- --------------------------------------
}
