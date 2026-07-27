package main

import (
	"fmt"
	"sync"
)

var rwmu sync.RWMutex
var counter int

func readCounter(wg *sync.WaitGroup) {
	defer wg.Done()

	rwmu.RLock()
	fmt.Println("read counter:", counter)
	rwmu.RUnlock()
}

func writeCounter(wg *sync.WaitGroup, value int) {
	defer wg.Done()

	rwmu.Lock()
	counter = value
	fmt.Println("value set to counter:", value)
	rwmu.Unlock()
}

func main() {

	var wg sync.WaitGroup
	for range 5 {
		wg.Add(1)
		go readCounter(&wg)
	}

	wg.Add(1)
	go writeCounter(&wg, 18)

	wg.Wait()

}
