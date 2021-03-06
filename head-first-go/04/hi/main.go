// B"H

package main

import (
	"fmt"
	"log"

	"github.com/Ylazerson/go-shenanigans/head-first-go/04/greeting"
	"github.com/Ylazerson/go-shenanigans/head-first-go/04/greeting/dansk"
	"github.com/Ylazerson/go-shenanigans/head-first-go/04/greeting/deutsch"
	"github.com/Ylazerson/go-shenanigans/head-first-go/04/keyboard"
)

func main() {

	fmt.Print("Choose a language:\n")
	fmt.Print("1 : English\n")
	fmt.Print("2 : Dansk\n")
	fmt.Print("3 : Deutsch\n")

	// -- -----------------------------
	var lang float64
	var err error

	lang, err = keyboard.GetFloat()

	if err != nil {
		log.Fatal(err)
	}
	// -- -----------------------------

	// -- -----------------------------
	if lang == 1 {
		greeting.Hello()
		greeting.Hi()
	} else if lang == 2 {
		dansk.Hej()
		dansk.GodMorgen()
	} else {
		deutsch.Hallo()
		deutsch.GutenTag()
	}
	// -- -----------------------------

}
