package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be zero")

func main() {
	// divisor := 7
	divisor := 0
	if q, r, err := divide(100, divisor); err == nil {
		fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
	} else if err == ErrDivideByZero {
		fmt.Println("Please do not attempt to divide by zero")
	} else {
		fmt.Println("Unknown error :", err)
	}
}

func divide(multiplier, divisor int) (quotient, remainder int, err error) {
	if divisor == 0 {
		err = ErrDivideByZero
		return
	}
	quotient = multiplier / divisor
	remainder = multiplier % divisor
	return
}
