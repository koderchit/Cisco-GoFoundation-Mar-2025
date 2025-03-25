package main

import "fmt"

func main() {
	/*
		var p1 struct {
			Id   int
			Name string
			Cost float64
		}
		p1.Id = 100
		p1.Name = "Pen"
		p1.Cost = 10
	*/

	var p1 struct {
		Id   int
		Name string
		Cost float64
	} = struct {
		Id   int
		Name string
		Cost float64
	}{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	// fmt.Println(p1)
	// fmt.Printf("%#v\n", p1)
	// fmt.Printf("%+v\n", p1)
	fmt.Println(Format(p1))
}

func Format(p struct {
	Id   int
	Name string
	Cost float64
}) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}
