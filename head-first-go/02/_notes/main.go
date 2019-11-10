// B"H

package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()

	// 2019-11-10 03:05:02.692168056 -0500 EST m=+0.000191694

	//fmt.Println
	fmt.Println(now)
	fmt.Println(now.Day())

	var nowDay int = now.Day()
	fmt.Println(nowDay)

	var nowDay2 int = time.Now().Day()
	fmt.Println(nowDay2)

	// -- ---------------------------------------
	var grade int = 65
	var passingGrade int = 60

	if grade >= passingGrade {
		fmt.Println("You Passed")

		var hello string = "You are doing great"
		fmt.Println(hello)
	}

	var passed bool = grade >= passingGrade
	fmt.Println(passed)

	var b4If string = "Before if"

	if passed {
		fmt.Println("You Passed")

		var congrats string = "You are doing great"
		fmt.Println(congrats)
		fmt.Println(b4If)

	}
	// -- ---------------------------------------

	// -- ---------------------------------------

	fmt.Println("Starting for loop")
	for x := 3; x <= 6; x++ {
		fmt.Println(x)
	}

	fmt.Println("Starting 2nd for loop")
	for x := 10; x >= 5; x-- {
		fmt.Println(x)
	}

	fmt.Println("Starting 3rd for loop")
	for x := 10; x >= 5; x-- {

		fmt.Println("before check")

		if x == 7 {
			continue
		}

		fmt.Println("after check")
		fmt.Println(x)
	}

	fmt.Println("Starting 4th for loop")
	for x := 10; x >= 5; x-- {

		fmt.Println("before check")

		if x == 7 {
			break
		}

		fmt.Println("after check")
		fmt.Println(x)
	}
	// -- ---------------------------------------

}
