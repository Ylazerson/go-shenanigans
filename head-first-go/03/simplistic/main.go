// B"H

package main

import (
	"errors"
	"fmt"
	"log"
)

// -- -----------------------------------
func double(num float64) float64 {
	return num * 2
}

// -- -----------------------------------
func sayHello(name string, saying string) (output string, err error) {

	if name == "baruch" {
		return saying, err
	} else {
		return "", fmt.Errorf("You are not Baruch, %s", name)
	}

}

// -- -----------------------------------
func main() {

	fmt.Println(double(35.78))

	// -- -------------------------------
	// Utilizing custom errors:
	var myError error = errors.New("No way no how")
	fmt.Println(myError)
	//log.Fatal(myError)

	// -- -------------------------------
	// Utilizing custom errors with formmating:
	var myError2 error = fmt.Errorf("No way no how %s", "Jack")
	fmt.Println(myError2)
	//log.Fatal(myError2)

	// -- -------------------------------
	// Get output without error
	output, err := sayHello("baruch", "helllllllo")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)

	// -- -------------------------------
	// Get output with error
	output, err = sayHello("zalman", "helllllllo")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)

}
