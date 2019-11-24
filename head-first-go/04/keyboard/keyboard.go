// B"H

// jfdskjhklfjdsackage keyboard reads user input from the keyboard.
package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetFloat reads a floating-point number from the keyboard.
// It returns the number read and any error encountered.
func GetFloat() (float64, error) {

	reader := bufio.NewReader(os.Stdin)

	// -- ---------------------------------
	// Read string from keyboard:
	input, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}
	// -- ---------------------------------

	// -- ---------------------------------
	// Convert to float:
	input = strings.TrimSpace(input)

	number, err := strconv.ParseFloat(input, 64)

	if err != nil {
		return 0, err
	}
	// -- ---------------------------------

	return number, nil
}

func doNothing() {
	fmt.Println("waste of time")
}
