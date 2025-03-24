package main

import "fmt"

func main() {
	var x int
	x = 100

	// value -> address (referencing)
	var xPtr *int //pointer type
	xPtr = &x     //address/pointer of x
	fmt.Println(xPtr)

	// address -> value (deferencing)
	fmt.Println(*xPtr) // value at xPtr
}
