package main

import (
	"fmt"
	"time"
)

func main() {
	// ArrayBasic()
	// ArrayInitialization()
	// ArraySlice()
	// Maps()
	// Iterations()
	ArrayTest()
}

func ArrayBasic() {
	var intArr [4]int32 // set a fixed length of int array so [0 0 0]

	intArr[0] = 3 // Assign value of that index
	intArr[1] = 9 // Assign value of that index
	intArr[2] = 7 // Assign value of that index
	intArr[3] = 4 // Assign value of that index

	// fmt.Println(intArr[0])   // First index
	fmt.Println(intArr[1:3]) // Index 1 and 2 (does not include the last index)
	fmt.Println(&intArr[2])  // Memory allocation
}

func ArrayInitialization() {
	var intArr [4]int32 = [4]int32{9, 2, 7, 15} // Explicitly initialize an array

	for i, v := range intArr { // iterate with index and value
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}
}

func ArraySlice() {
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("Original Length: %v, Capacity: %v\n", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7)                                                  // add value into the array to the last index
	fmt.Printf("Appended Length: %v, Capacity: %v\n", len(intSlice), cap(intSlice)) // length is added by 1, but the capacity will add more than 1

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8) // Set a fixed length and capacity
	fmt.Println(len(intSlice3), cap(intSlice3), intSlice3)
}

func Maps() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Eve": 45}
	fmt.Println(myMap2["Adam"])  // returns 23
	fmt.Println(myMap2["Jacob"]) // returns 0 if not exist

	var age, ok = myMap2["Abraham"] // first argument is the value, 2nd will check if it exists

	delete(myMap2, "Eve") // Delete key from the map

	fmt.Println(myMap2)
	if !ok {
		fmt.Printf("Invalid name")
	} else {
		fmt.Printf("Age: %v", age)
	}
}

func Iterations() {
	var myMap = map[string]uint8{"Irving": 28, "Craig": 30, "Real": 41} // Initialize a map

	for name, age := range myMap { // Iterate with key and value
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	for i := 0; i < 10; i++ { // Basic for loop
		fmt.Println(i)
	}

	// No while loop, use this instead:
	var idx int = 0
	for idx < 10 {
		if idx >= 10 { // break conditions
			break
		}

		fmt.Println(idx)
		idx = idx + 1
	}
}

func ArrayTest() {
	var n int = 1000000
	var testSlice = []int{}            // unallocated
	var testSlice2 = make([]int, 0, n) // allocated

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}

	return time.Since(t0)
}
