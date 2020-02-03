// B"H

package main

import "fmt"

// -- ---------------------------
// Function with only a variadic parameter:
func showNames(names ...string) {

	fmt.Println("\n showNames is running:")

	fmt.Println(names)

}

// -- ---------------------------
// Function with 2 parameters; 1st is string, 2nd is variadic.
func showCars(carMake string, carModels ...string) {

	fmt.Println("\n showCars is running:")

	fmt.Println("Car make:", carMake)

	fmt.Println(carModels)

}

// -- ---------------------------
func main() {

	showNames("Moishy", "Sarah")
	showNames("Larry", "Moe", "Curly")
	showNames()

	showCars("Honda", "Accord", "Civic")
	showCars("Ford", "F150")
	showCars("Ferrari")

	// -- -------------------------------------
	var carMake string = "Toyota"
	carModels := []string{"Sienna", "Corrola"}

	// Approach 1:
	showCars(carMake, carModels...)

	// Approach 2:
	showCars("Toyota", "Sienna", "Corrola")

}
