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
	/*


		func(firstName, lastName string) {
			fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
		}("Magesh", "Kuppan")

		msg := func(firstName, lastName string) string {
			return fmt.Sprintf("Hi %s %s, Have a good day!", firstName, lastName)
		}("Suresh", "Kannan")
		fmt.Println(msg)
	*/
}
