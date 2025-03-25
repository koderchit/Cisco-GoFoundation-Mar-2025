package main

import (
	"errors"
	"fmt"
)

func main() {
	// divisor := 7
	divisor := 0
	if q, r, err := divide(100, divisor); err == nil {
		fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
	} else {
		fmt.Println("Error :", err)
	}
}

/*
func divide(multiplier, divisor int) (int, int, error) {
	if divisor == 0 {
		err := errors.New("divisor cannot be zero")
		return 0, 0, err
	}
	quotient := multiplier / divisor
	remainder := multiplier % divisor
	return quotient, remainder, nil
}
*/

func divide(multiplier, divisor int) (quotient, remainder int, err error) {
	if divisor == 0 {
		err = errors.New("divisor cannot be zero")
		return
	}
	quotient = multiplier / divisor
	remainder = multiplier % divisor
	return
}
