// B"H

package main

import "fmt"

func main() {

	// -- ---------------------------------
	// Map example 1
	fmt.Println("------------------------")
	fmt.Println("Map example 1:")
	fmt.Println("Use this 90% of the time")

	shopList := make(map[string]int)

	shopList["shoes"] = 5
	shopList["hat"] = 1
	shopList["tie"] = 2
	shopList["cups"] = 50
	shopList["cups"] = 75

	fmt.Println(shopList)

	fmt.Println(shopList["cups"])
	fmt.Println(shopList["zzzz"])
	// -- ---------------------------------

	// -- ---------------------------------
	// Map example 2
	fmt.Println("------------------------")
	fmt.Println("Map example 1, using map literals:")

	houses := map[string]float64{
		"21 main st":  78349.65,
		"943 loon st": 8989.65,
	}

	houses["943 loon st"] = 5.78
	houses["huh ave 99"] = 1.9898989

	fmt.Println(houses)
	// -- ---------------------------------

	// -- ---------------------------------
	fmt.Println("------------------------")
	fmt.Println("... having fun ...:")

	countries := make(map[string]int)

	countries["aus"]++
	countries["aus"]++

	fmt.Println(countries)
	// -- ---------------------------------

}
