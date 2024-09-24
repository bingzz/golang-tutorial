package main

import "fmt"

type GasEngine struct { // Type definition
	mpg       uint8
	gallons   uint8
	ownerInfo Owner // Add a model
}

type ElectricEngine struct {
	mpkwh     uint8
	kwh       uint8
	ownerInfo Owner
}

type Owner struct {
	name string
	age  uint8
}

type Engine interface {
	milesLeft() uint8 // Method signature
}

func main() {
	var gasEngine GasEngine = GasEngine{mpg: 10, gallons: 30, ownerInfo: Owner{name: "Irving", age: 28}} // Initialize value
	var ElectricEngine ElectricEngine = ElectricEngine{mpkwh: 50, kwh: 100, ownerInfo: Owner{name: "Jason", age: 35}}

	gasEngine.mpg = 15 // Can be modified
	// fmt.Println(myEngine.gallons, myEngine.mpg)
	// fmt.Printf("Total miles left: %v\n", myEngine.milesLeft())
	canMakeIt(gasEngine, 20)
	canMakeIt(ElectricEngine, 255)
}

func (e GasEngine) milesLeft() uint8 { // Assigning a function for the specific type (similar to creating a class and its function)
	return e.gallons * e.mpg
}

func (e ElectricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func canMakeIt(e Engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("Can make it")
	} else {
		fmt.Println("Needs refueling")
	}
}
