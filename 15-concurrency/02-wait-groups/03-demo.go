package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go f1() // scheduling the execution of "f1" through the scheduler
	f2()
	wg.Wait() // block until the counter becomes 0 (default = 0)
}

func f1() {
	fmt.Println("f1 invoked")
	wg.Done() // decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
