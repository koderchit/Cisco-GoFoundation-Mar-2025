package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Category string
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float64) /* no return */ {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
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

// Overriding the Product.Format() method
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}

func main() {

	milk := NewPerishableProduct(201, "Milk", 52, "Food", "2 Days")

	fmt.Println(milk.Format())
	milk.ApplyDiscount(10)
	fmt.Println(milk.Format())
}
