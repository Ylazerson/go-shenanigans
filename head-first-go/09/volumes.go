// B"H

package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func main() {

	// -- --------------------------------------
	var soda Liters
	var water Milliliters

	soda = Liters(2)
	water = Milliliters(500)

	// -- --------------------------------------
	fmt.Println("Soda:", soda)
	fmt.Println("Water:", water)

	fmt.Println("Soda + float64:", soda+2.5)

	// -- --------------------------------------
	var sodaGallons Gallons
	sodaGallons = soda.ToGallons()

	fmt.Println("sodaGallons", sodaGallons)

	// -- --------------------------------------
	var waterGallons Gallons
	waterGallons = water.ToGallons()

	fmt.Println("waterGallons", waterGallons)

}
