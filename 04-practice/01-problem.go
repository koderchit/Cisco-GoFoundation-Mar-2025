/*
Write simple program that prints the prime numbers between 2 - 100
*/

package main

import "fmt"

func main() {
LOOP:
	for n := 2; n <= 100; n++ {
		for i := 2; i <= (n / 2); i++ {
			if n%i == 0 {
				continue LOOP
			}

		}
		fmt.Println("Prime :", n)
	}
}
