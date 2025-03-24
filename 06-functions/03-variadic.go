package main

import "fmt"

func main() {
	fmt.Println(sum(10))             //=> 10
	fmt.Println(sum(10, 20))         //=> 30
	fmt.Println(sum(10, 20, 30, 40)) //=> 100
	fmt.Println(sum())               //=> 0

	list := []int{100, 700, 300, 200}
	fmt.Println(sum(list...))
}

func sum(nos ...int) int {
	var result int
	for i := 0; i < len(nos); i++ {
		result += nos[i]
	}
	return result
}
