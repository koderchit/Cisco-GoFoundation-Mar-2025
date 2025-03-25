package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {
	defer func() {
		if e := recover(); e != nil {
			log.Println("App panicked.. Error :", e)
			return
		}
		fmt.Println("Thank you!")
	}()
	var divisor int
	for {
		fmt.Println("Enter the divisor")
		fmt.Scanln(&divisor)
		if q, r, err := divideAdapter(100, divisor); err != nil {
			fmt.Println(err)
			continue
		} else {
			fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
			break
		}
	}
}

// adapter function to convert the panic into an error
func divideAdapter(multiplier, divisor int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(multiplier, divisor)
	return
}

// 3rd party api (cannot be changed)
func divide(multiplier, divisor int) (quotient, remainder int) {
	log.Println("[divide] - calculating quotient")
	if divisor == 0 {
		panic(ErrDivideByZero)
	}
	quotient = multiplier / divisor
	log.Println("[divide] - calculating remainder")
	remainder = multiplier % divisor
	log.Println("[divide] - returning the results")
	return
}
