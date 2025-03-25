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
		if q, r, err := divide(100, divisor); err == ErrDivideByZero {
			fmt.Println("do not attempt to divide by 0")
			continue
		} else {
			fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
			break
		}
	}
}

func divide(multiplier, divisor int) (quotient, remainder int, err error) {
	defer func() {
		log.Println("	[divide] - deferred")
	}()
	log.Println("[divide] - calculating quotient")
	if divisor == 0 {
		err = ErrDivideByZero
		return
	}
	quotient = multiplier / divisor
	log.Println("[divide] - calculating remainder")
	remainder = multiplier % divisor
	log.Println("[divide] - returning the results")
	return
}
