package main

import "fmt"

func main() {
	go f1() // scheduling the execution of "f1" through the scheduler
	f2()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
