package main

import (
	"fmt"

	"github.com/tkmagesh/Cisco-GoFoundation-Mar-2025/10-modularity-app/calculator"
	// "github.com/tkmagesh/Cisco-GoFoundation-Mar-2025/10-modularity-app/calculator/utils"
	ut "github.com/tkmagesh/Cisco-GoFoundation-Mar-2025/10-modularity-app/calculator/utils"
)

func main() {
	fmt.Println("App executed!")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println("Op Count :", calculator.OpCount())
	// fmt.Println("is 21 odd? :", utils.IsOdd(21))
	fmt.Println("is 21 odd? :", ut.IsOdd(21))
}
