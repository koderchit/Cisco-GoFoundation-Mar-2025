package main

import "fmt"

func main() {
	var sayHi func()
	// fmt.Println(sayHi)

	sayHi = func() {
		fmt.Println("Hi!")
	}
	sayHi()

	var sayHello func(string)
	sayHello = func(name string) {
		fmt.Printf("Hello %s!\n", name)
	}
	sayHello("Magesh")

	var greet func(string, string)
	greet = func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}
	greet("Magesh", "Kuppan")

	var getGreetMsg func(string, string) string
	getGreetMsg = func(firstName, lastName string) string {
		return fmt.Sprintf("Hi %s %s, Have a good day!", firstName, lastName)
	}
	msg := getGreetMsg("Suresh", "Kannan")
	fmt.Println(msg)

	var divide func(int, int) (int, int)
	divide = func(multiplier, divisor int) (int, int) {
		quotient := multiplier / divisor
		remainder := multiplier % divisor
		return quotient, remainder
	}
	q, r := divide(100, 7)
	fmt.Printf("quotient = %d, remainder = %d\n", q, r)

}
