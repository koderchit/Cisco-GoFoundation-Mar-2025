package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		r := add(100, 200)
		ch <- r
	}()
	fmt.Println("Awating result....")
	result := <-ch
	fmt.Println("Result :", result)

}

// 3rd party api (cannot change this)
func add(x, y int) int {
	return x + y
}
