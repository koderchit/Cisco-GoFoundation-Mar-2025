/*
Accept a range from the user
Write simple program that prints the prime numbers between the given range
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
