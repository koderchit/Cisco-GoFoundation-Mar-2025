package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Category string
}

type Dummy struct {
	Id int
}

type PerishableProduct struct {
	Product
	// Dummy
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float64, category string, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:       id,
			Name:     name,
			Cost:     cost,
			Category: category,
		},
		Expiry: expiry,
	}
}

func main() {
	/*
		milk := PerishableProduct{
			Product: Product{
				Id:       201,
				Name:     "Milk",
				Cost:     52,
				Category: "Food",
			},
			Expiry: "2 Days",
		}
	*/
	milk := NewPerishableProduct(201, "Milk", 52, "Food", "2 Days")
	// fmt.Println(milk.Product.Id, milk.Product.Name, milk.Product.Cost, milk.Expiry)
	fmt.Println(milk.Id, milk.Name, milk.Cost, milk.Expiry)
}
