// B"H

package main

import "fmt"

func main() {

	// % is called a Formatting Verb
	fmt.Printf(
		"The %s cost %d cents each.\n",
		"gumballs",
		23,
	)

	var houseCost float64 = 178.67
	var houses int64 = 35
	var name string = "123"
	var isHappy bool = true

	// -- --------------------------------
	// Basic output
	fmt.Println("-------------------------------")
	fmt.Printf("The house costs %f \n", houseCost)
	fmt.Printf("The houses %d \n", houses)
	fmt.Printf("The name %s \n", name)
	fmt.Printf("The isHappy %t \n", isHappy)

	// -- --------------------------------
	// Basic output with set minimum width
	fmt.Println("-------------------------------")
	fmt.Printf("%10f \n", houseCost)
	fmt.Printf("%10d \n", houses)
	fmt.Printf("%10s \n", name)
	fmt.Printf("%10t \n", isHappy)

	// -- --------------------------------
	// Basic output with set minimum width
	fmt.Println("-------------------------------")
	fmt.Printf(
		"%15s | %10d | %5.3f  \n",
		"Baruch",
		15625,
		349.2,
	)

	fmt.Printf(
		"%15s | %10d | %5.3f  \n",
		"Mendy Efune",
		99,
		0.5,
	)

	fmt.Printf(
		"%15s | %10d | %5.3f  \n",
		"Yossi",
		0,
		99.0,
	)

	fmt.Printf(
		"%15s | %10d | %5.3f  \n",
		"Zman",
		8988888,
		5.345,
	)

	// -- --------------------------------
	fmt.Println("-------------------------------")
	fmt.Printf("The name %#v \n", name)
	fmt.Printf("The name %T \n", name)

	// -- --------------------------------
	var funkyChars string = "\t"
	fmt.Println("-------------------------------")
	fmt.Printf("The name z%sz \n", funkyChars)
	fmt.Printf("The name z%vz \n", funkyChars)
	fmt.Printf("The name z%#vz \n", funkyChars)

}
