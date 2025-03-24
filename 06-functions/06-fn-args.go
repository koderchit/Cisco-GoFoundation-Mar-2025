package main

import "fmt"

/*
func main() {
	exec("f1") //=> f1 is executed
	exec("F2") //=> f2 is executed
}

func exec(fnName string) {
	// execute a function based on the argument
	switch fnName {
	case "f1":
		f1()
	case "f2":
		f2()
	default:
		fmt.Println("Invalid argument")
	}
}
*/

func main() {
	exec(f1)
	exec(f2)
	exec(func() {
		fmt.Println("anonymous fn invoked")
	})
}

func exec(fn func()) {
	fn()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
