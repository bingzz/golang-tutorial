package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_CHICKEN_PRICE float32 = 5
const MAX_TOFU_PRICES float32 = 3

func main() {
	// createChannel()
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	var chickenChannel = make(chan string) // Channel string
	var tofuChannel = make(chan string)

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}

	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20

		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website // assign channel with this matched condition
			break
		}
	}
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20

		if tofuPrice <= MAX_TOFU_PRICES {
			tofuChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	// fmt.Printf("Found a deal on chicken at %s\n", <-chickenChannel)
	// fmt.Printf("Found a deal on tofu at %s\n", <-tofuChannel)

	select { // Switch statements for channel
	case website := <-chickenChannel:
		fmt.Printf("Found a deal on chicken at %s\n", website)
	case website := <-tofuChannel:
		fmt.Printf("Found a deal on tofu at %s\n", website)
	}
}

func createChannel() {
	var c = make(chan int, 5) // Hold a single int value
	go process(c)             // Goroutine execution

	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

func process(c chan int) {
	defer close(c) // Execute before the function exits
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Exiting process")
}
