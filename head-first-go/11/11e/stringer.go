// B"H

package main

import "fmt"

// -- -------------------------------------
/*

https://golang.org/pkg/fmt/#example_Stringer

From Go's fmt package:

type Stringer interface {
    String() string
}

*/
// -- -------------------------------------

// -- -------------------------------------
type Ounces float64

// -- -------------------------------------
type Gallons float64

// Satisfies the fmt.Stringer interface:
func (g Gallons) String() string {
	return fmt.Sprintf("Here is your %0.2f gal", g)
}

// -- -------------------------------------
type Liters float64

// Satisfies the fmt.Stringer interface:
func (l Liters) String() string {
	return fmt.Sprintf("Here is your %0.2f L", l)
}

// -- -------------------------------------
type Milliliters float64

// Satisfies the fmt.Stringer interface:
func (m Milliliters) String() string {
	return fmt.Sprintf("Here is your %0.2f mL", m)
}

// -- -------------------------------------
func main() {

	fmt.Println(Ounces(12.09248342))

	fmt.Println(Gallons(12.09248342))

	fmt.Println(Liters(12.09248342))

	fmt.Println(Milliliters(12.09248342))
}
