package main

import (
	"fmt"
	"sync"
)

// MUTEX WITHOUT STRUCT
func main() {
	var counter int
	var wg sync.WaitGroup
	// var mu sync.Mutex

	numGoroutines := 5
	wg.Add(numGoroutines)

	increment := func() {
		defer wg.Done()
		for range 1000 {
			// mu.Lock()
			counter++ //critical section
			// mu.Unlock()
		}
	}

	for range numGoroutines {
		go increment()
	}

	wg.Wait()
	fmt.Println("Finishing. final counter value: ", counter)
}
