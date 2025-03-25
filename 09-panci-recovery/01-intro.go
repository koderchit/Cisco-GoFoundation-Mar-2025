package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			log.Println("App panicked.. Error :", e)
			return
		}
		fmt.Println("Thank you!")
	}()
	divisor := 0
	fmt.Println("Enter the divisor")
	fmt.Scanln(&divisor)
	q, r := divide(100, divisor)
	fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

func divide(multiplier, divisor int) (quotient, remainder int) {
	defer func() {
		log.Println("	[divide] - deferred")
	}()
	log.Println("[divide] - calculating quotient")
	quotient = multiplier / divisor
	log.Println("[divide] - calculating remainder")
	remainder = multiplier % divisor
	log.Println("[divide] - returning the results")
	return
}
