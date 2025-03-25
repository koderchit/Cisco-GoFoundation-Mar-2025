package calculator

import "fmt"

func init() {
	fmt.Println("[calculator/add.go] - init invoked")
}

func Add(x, y int) int {
	// opCount++
	opCount["add"]++
	return x + y
}
