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
	primes := genPrimes(start, end)
	fmt.Printf("Prime numbers between %d and %d\n", start, end)
	for _, primeNo := range primes {
		fmt.Println("Prime :", primeNo)
	}
}

func genPrimes(start, end int) []int {
	primes := make([]int, 0)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
	return primes
}
func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
