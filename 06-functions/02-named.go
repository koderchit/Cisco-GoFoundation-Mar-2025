package main

import "fmt"

func main() {
	// fmt.Println(divide(100, 7))

	q, r := divide(100, 7)
	fmt.Printf("quotient = %d, remainder = %d\n", q, r)

	// Use only the quotient
	/*
		q, _ := divide(100, 7)
		fmt.Printf("quotient = %d\n", q)
	*/
}

/*
func divide(multiplier, divisor int) (int, int) {
	quotient := multiplier / divisor
	remainder := multiplier % divisor
	return quotient, remainder
}
*/

// named results
func divide(multiplier, divisor int) (quotient, remainder int) { /* quotient & remainder are declared and initialized */
	quotient = multiplier / divisor
	remainder = multiplier % divisor
	return
}
