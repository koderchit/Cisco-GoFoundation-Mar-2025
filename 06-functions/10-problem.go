/*
Create an interactive calculator

When the program is executed:

Display the following menu
1. Add
2. Subtract
3. Multiply
4. Divide
5. Exit

Accept the user choice
if user choice == 5
	then exit the application

if user choice == 1 to 4
	accept the 2 operands from the user
	perform the corresponding operation
	print the result
	display the menu again (IMPORTANT)

if user choice < 1 OR > 5
	print "invalid choice"
	display the menu again (IMPORTANT)

Note:
Ensure that the solution is made of functions and each function has ONLY ONE reason to change (SRP)

*/

package main

import "fmt"

func main() {
	var userChoice int
	for {
		getUserChoice(&userChoice)
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		if userChoice == 5 {
			break
		}
		processChoice(userChoice)
	}
	fmt.Println("Thank you!")
}

func getUserChoice(userChoice *int) {
	fmt.Println("Choose an option")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Scanln(userChoice)
}

func getOperands(n1, n2 *int) {
	fmt.Println("Enter the operands:")
	fmt.Scanln(n1, n2)
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

func processChoice(userChoice int) {
	var n1, n2, result int
	getOperands(&n1, &n2)
	switch userChoice {
	case 1:
		result = add(n1, n2)
	case 2:
		result = subtract(n1, n2)
	case 3:
		result = multiply(n1, n2)
	case 4:
		result = divide(n1, n2)
	}
	fmt.Println("Result :", result)
}
