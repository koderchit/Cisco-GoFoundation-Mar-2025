package main

import "fmt"

func main() {
	var no int = 100
	fmt.Println("Before incrementing, no =", no)
	increment(&no)
	fmt.Println("After incrementing, no =", no)
}

func increment(x *int) /* no return values */ {
	fmt.Println("[increment] address of x :", &x)
	*x += 1
}
