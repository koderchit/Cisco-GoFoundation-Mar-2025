/*
Anonymous functions
	- functions created in another function
	- cannot have any name
	- have to be immediately invoked
*/

package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hi!")
	}()

	func(name string) {
		fmt.Printf("Hello %s!\n", name)
	}("Magesh")

	func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}("Magesh", "Kuppan")

	msg := func(firstName, lastName string) string {
		return fmt.Sprintf("Hi %s %s, Have a good day!", firstName, lastName)
	}("Suresh", "Kannan")
	fmt.Println(msg)
}
