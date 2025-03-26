package main

import (
	"fmt"
	"sync"
)

// Custom type that encapsulates the state manipulation logic in a concurrent safe manner
type SafeCounter struct {
	sync.Mutex
	count int
}

func (sf *SafeCounter) Add(delta int) {
	sf.Lock()
	{
		sf.count += delta
	}
	sf.Unlock()
}

var sf SafeCounter

func main() {
	wg := &sync.WaitGroup{}
	for range 200 {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("Count :", sf.count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	sf.Add(1)
}
