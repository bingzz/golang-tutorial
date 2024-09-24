package main

import (
	"fmt"
	"strings"
)

func main() {
	// StringIterate()
	AppendStrings()
	// Runes()
}

func StringIterate() {
	var myString = "resume" // translated into binary code
	// var myString = []rune("resume") // int32

	var indexed = myString[0]
	// fmt.Println(indexed)
	fmt.Printf("%v, %T\n", indexed, indexed) // displayed as a utf8 encoding

	for i, v := range myString {
		fmt.Println(i, v)
	}
}

func Runes() {
	// var myRune = "a" // string
	var myRune = 'a' // rune

	fmt.Printf("My Rune: %v\n", myRune)
}

func AppendStrings() {
	var strSlice = []string{"i", "r", "v", "i", "n", "g"}
	// var concatenated = ""
	var strBuilder strings.Builder

	for i := range strSlice {
		strBuilder.WriteString(strSlice[i]) // use the strBuilder instead to prevent from repetitive iteration (much faster)
		// concatenated += strSlice[i] // This will generate a string every iteration
	}
	var catStr = strBuilder.String()

	fmt.Printf("%v\n", catStr)
}
