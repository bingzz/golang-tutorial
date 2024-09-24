package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int16 = 32767 // explicitly typed variable
	// intNum = intNum + 1 // will give incorrect results when doing this since int16 has a limit
	fmt.Println(intNum)

	FloatNum()
	AddDifferentTypes()
	Division()
	Strings()
	Booleans()
	DefaultValues()
	Constants()

}

func FloatNum() {
	// var floatNum float64 = 12345678.9
	floatNum := 12345678.9 // inferred
	fmt.Println(floatNum)
}

func AddDifferentTypes() {
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2

	// ❌ cannot add numbers with different types
	// var result float32 = floatNum32 + intNum32

	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)
}

func Division() {
	// Ways to declare a variable
	// 1
	// var intNum1 int = 5
	// var intNum2 int = 2

	// 2
	// var intNum1, intNum2 = 5, 2

	// 3
	intNum1, intNum2 := 5, 2

	// Returns an integer without the decimal
	fmt.Println(intNum1 / intNum2)

	// Returns the remainder
	fmt.Println(intNum1 % intNum2)

	// Returns the actual value
	fmt.Println(float64(intNum1) / float64(intNum2)) // Parse the values
}

func Strings() {
	var strDouble = "Hello World"
	var strBackQuote = `Hello World` + "\n"
	var strConcat = "Hello" + " " + "World"

	fmt.Println(strDouble, strBackQuote, strConcat)
	fmt.Println(len("irving craig real!)(")) // Not the length of characters of a string, it uses the number of bytes

	fmt.Println(utf8.RuneCountInString("irving craig real!)(")) // Counts the number of characters of the string
}

func Booleans() {
	var myBool bool = false

	fmt.Println(myBool)
}

func DefaultValues() {
	var defaultBool bool
	var defaultInt int
	var defaultFloat float32
	var defaultStr string

	fmt.Println(defaultBool, defaultInt, defaultFloat, defaultStr)
}

func Constants() {
	const myConst string = "constant variable" // Always initialize a value for constants

	fmt.Println(myConst)
}
