// B"H

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func double(num float64) float64 {
	return num * 2
}

func main() {

	// -- -----------------------------------
	// Get a standard input:
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	fmt.Print("Enter a number: ")

	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	// -- -----------------------------------
	// Trim space and convert to float64:
	input = strings.TrimSpace(input)

	inputNum, err := strconv.ParseFloat(input, 64)

	if err != nil {
		log.Fatal(err)
	}

	// -- -----------------------------------
	fmt.Println(double(inputNum))

}
