package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{}    // Mutual exclusion (prevent memory from causing errors)
var wg = sync.WaitGroup{} // Counters
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	var t0 = time.Now()

	for i := 0; i < len(dbData); i++ {
		wg.Add(1)    // Add Counter
		go dbCall(i) // execute concurrently (execute goroutine)
	}

	wg.Wait() // Waits until the counter sets to 0
	fmt.Printf("Total Execution Time: %v\n", time.Since(t0))
	fmt.Printf("Results: %v\n", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond) // pause for 1 second

	save(dbData[i])
	log()
	wg.Done() // Decrement the counter
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("Current Results: %v\n", results)
	m.RUnlock()
}
