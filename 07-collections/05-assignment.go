/*
Refactor the following using functions

isPrime() => to check if the given number is a prime or not
genPrimes(start, end) => returns the generated prime numbers
main() => accept the range and prints the prime numbers
*/

package main

import "fmt"

func main() {
	var start, end int
	fmt.Println("Enter the start and the end (seperated by space)")
	fmt.Scanln(&start, &end)
	fmt.Printf("Prime numbers between %d and %d\n", start, end)
LOOP:
	for n := start; n <= end; n++ {
		for i := 2; i <= (n / 2); i++ {
			if n%i == 0 {
				continue LOOP
			}

		}
		fmt.Println("Prime :", n)
	}
}
