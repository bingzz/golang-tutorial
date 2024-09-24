package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World"
	PrintMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = IntDivision(numerator, denominator) // function that returns multiple values

	// Handle the error
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// 	return
	// }

	// if remainder == 0 {
	// 	fmt.Printf("Integer Division Result: %v", result)
	// } else {
	// 	fmt.Printf("Integer Division Result: %v, Remainder: %v", result, remainder)
	// }

	// Switch statements
	switch { // General if statements
	case err != nil:
		fmt.Printf(err.Error())
		// break // not needed
	case remainder == 0:
		fmt.Printf("Integer Division Result: %v", result)
		// break
	default:
		fmt.Printf("Integer Division Result: %v, Remainder: %v", result, remainder)
	}

	switch remainder { // Specific variable
	case 0:
		fmt.Printf("division was exact")
	case 1, 2:
		fmt.Printf("division was close")
	default:
		fmt.Printf("divison was not close")
	}
}

func PrintMe(printValue string) {
	fmt.Println(printValue)
}

func IntDivision(numerator int, denominator int) (int, int, error) { // functionName (arguments): (multiple return values)
	var err error // Default value of error is nil
	if denominator == 0 {
		err = errors.New("Cannot divide by Zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, err
}
