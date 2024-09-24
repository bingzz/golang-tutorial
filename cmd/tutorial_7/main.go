package main

import "fmt"

func main() {

	// Pointers()
	// Slices()

	// var thing1 = [5]float64{1, 2, 3, 4, 5}
	// fmt.Printf("Memory location of thing1 Array: %p\n", &thing1)

	// var result [5]float64 = Square(&thing1)
	// fmt.Printf("Result: %v\n", result)
	// fmt.Printf("Thing1 Result: %v\n", thing1)

	PointerTest()
}

func Pointers() {
	var p *int32 = new(int32) // Initialize pointer
	var i int32

	fmt.Printf("The value p points to is: %v\n", *p) // Dereferencing the pointer
	fmt.Printf("The value of i is: %v\n", i)

	// *p = 10 // Reassign the value
	p = &i
	*p = 6
	fmt.Printf("The value p points to is: %v\n", *p) // Dereferencing the pointer
	fmt.Printf("The value of i is: %v\n", i)
}

func Slices() {
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice

	sliceCopy[2] = 4

	fmt.Println(slice)
	fmt.Println(sliceCopy)
}

func Square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("Memory location of thing2 Array: %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}

	return *thing2 // value
}

func PointerTest() {
	var a = 1  // initialize a variable
	var b = &a // get its memory from a
	var c = &b
	var d = &c
	var e = &d
	var f = &e

	*****f = 10 // applies to the source, which will apply to the pointed values
	**c = 25
	fmt.Println(a, *b, **c, ***d, ****e, *****f)
}
