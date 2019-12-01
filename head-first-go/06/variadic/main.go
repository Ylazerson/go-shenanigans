// B"H

package main

import "fmt"

// -- ---------------------------
// Function with only a variadic parameter:
func showNames(names ...string) {

	fmt.Println("\n showNames is running:")

	for index, value := range names {
		fmt.Println("index is", index, "- value is", value)
	}

}

// -- ---------------------------
// Function with only a variadic parameter:
func showCars(carMake string, carModels ...string) {

	fmt.Println("\n showCars is running:")

	fmt.Println("Car make:", carMake)

	for index, value := range carModels {
		fmt.Println("index is", index, "- value is", value)
	}

}

// -- ---------------------------
func main() {

	showNames("Moishy", "Sarah")
	showNames("Larry", "Moe", "Curly")
	showNames()

	showCars("Honda", "Accord", "Civic")
	showCars("Ford", "F150")
	showCars("Ferrari")

	var carMake string = "Toyota"
	carModels := []string{"Sienna", "Corrola"}
	showCars(carMake, carModels...)
}
