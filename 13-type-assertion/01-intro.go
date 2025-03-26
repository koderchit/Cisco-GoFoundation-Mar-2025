package main

import "fmt"

type Product struct {
	Id   int
	Name string
}

func main() {
	// var x interface{}
	var x any
	x = 100
	x = "Officia et do et dolor reprehenderit in esse consectetur irure enim."
	x = 99.99
	x = true
	x = struct{}{}
	fmt.Println(x)

	// x = 100
	// y := x * 2
	x = "Aliqua minim duis magna fugiat cupidatat sint."

	// not type safe
	// y := x.(int) * 2

	// type safe (type assertion using 'if-else')
	if val, ok := x.(int); ok {
		y := val * 2
		fmt.Println(y)
	} else {
		fmt.Println("value in x is not an int")
	}

	// type assertion using type-switch
	// x = 100
	// x = "Id ipsum amet culpa est fugiat fugiat cupidatat eu quis quis minim ullamco sunt sint."
	// x = 99.993
	// x = true
	x = Product{100, "Pen"}
	// x = struct{}{}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 2 =", val*2)
	case float64:
		fmt.Println("x is a float64, x * 0.02 =", val*0.02)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case Product:
		fmt.Printf("x is a Product : id = %d, name = %q\n", val.Id, val.Name)
	default:
		fmt.Println("x is an unknown type")
	}

}
