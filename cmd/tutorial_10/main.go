package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(SumSlice(intSlice)) // Assert that this value is an int data type

	var float32Slice = []float32{1.0, 2.0, 3.0}
	fmt.Println(SumSlice(float32Slice)) // Assert that this value is a float32 data type
}

func SumSlice[T int | float32 | float64](slice []T) T { // Generics (assert that this argument can have multiple data types)
	var sum T
	for _, v := range slice {
		sum += v
	}

	return sum
}
