package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("This will not be repeated")
}

func main() {

	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Goroutouein #", i)
			once.Do(initialize)
		}()
	}

	wg.Wait()

}
