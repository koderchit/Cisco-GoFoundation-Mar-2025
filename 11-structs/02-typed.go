package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Category string
}

func main() {
	/*
		var p1 Product
		p1.Id = 100
		p1.Name = "Pen"
		p1.Cost = 10
	*/

	/*
		var p1 Product = Product{
			Id:   100,
			Name: "Pen",
			Cost: 10,
		}
	*/
	/*
		var p1 = Product{
			Id:   100,
			Name: "Pen",
			Cost: 10,
		}
	*/

	p1 := Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	// p1 := Product{100, "Pen", 10} //not advisable

	// fmt.Println(p1)
	// fmt.Printf("%#v\n", p1)
	// fmt.Printf("%+v\n", p1)
	fmt.Println(Format(p1))

	var p2 Product
	p2 = p1 // creates a copy (coz struct instances are values)
	p1.Cost = 15
	fmt.Printf("p1.Cost = %v, p2.Cost = %v\n", p1.Cost, p2.Cost)

	px := Product{Id: 200, Name: "Pencil", Cost: 5}
	py := Product{Id: 200, Name: "Pencil", Cost: 5}
	fmt.Println(px == py) //comparison by value

	// references using pointers
	var p1Ptr *Product
	p1Ptr = &p1
	// No need to dererence to access the members of a pointer to a struct
	fmt.Println(p1Ptr.Id, p1Ptr.Name, p1Ptr.Cost)

	fmt.Println("Before applying discount")
	fmt.Println("p1 : ", Format(p1))
	ApplyDiscount( /* ? */ ) // (10%)
	fmt.Println("After applying discount")
	fmt.Println("p1 : ", Format(p1))

}

func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount( /*  */ ) /* no return */ {
	// apply the discount (percentage) on the given product
}
