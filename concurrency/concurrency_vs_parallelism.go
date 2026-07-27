package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// func printNumbers() {
// 	for i := range 5 {
// 		fmt.Println(i)
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }

// func printLetters() {
// 	for _, letter := range "ABCDE" {
// 		fmt.Println(string(letter))
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }

func HeavyTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d is starting\n", id)
	for range 100_000_000 {

	}
	fmt.Println(time.Now())
	fmt.Printf("Task %d is finished\n", id)
}

func main() {

	// go printNumbers()
	// go printLetters()
	// time.Sleep(3 * time.Second)

	numThreads := 4
	runtime.GOMAXPROCS(numThreads)
	var wg sync.WaitGroup

	for i := range numThreads {
		wg.Add(1)
		go HeavyTask(i, &wg)
	}

	wg.Wait()
}
