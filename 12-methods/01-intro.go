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

func main() {

	p1 := Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	fmt.Println("Before applying discount")

	// fmt.Println("p1 : ", Format(p1))
	fmt.Println("p1 : ", p1.Format())

	// ApplyDiscount(&p1, 10) // (10%)
	p1.ApplyDiscount(10)

	fmt.Println("After applying discount")

	// fmt.Println("p1 : ", Format(p1))
	fmt.Println("p1 : ", p1.Format())

}
