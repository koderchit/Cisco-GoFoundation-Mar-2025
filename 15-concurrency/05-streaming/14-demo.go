package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genPrimes(2, 100)
	for primeNo := range ch {
		fmt.Println("Prime No :", primeNo)
	}
	fmt.Println("Done!")
}

func genPrimes(start, end int) chan int {
	ch := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			if isPrime(no) {
				ch <- no
				time.Sleep(500 * time.Millisecond)
			}
		}
		close(ch)
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
