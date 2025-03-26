package main

import (
	"fmt"
	"sort"
	"strings"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f, Units = %d, Category = %q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

/*
Create a "Sort" api (function) to sort the products by any attribute
	IMPORTANT: DO NOT write your own logic for sorting. instead use sort.Sort() function
*/

type Products []Product

func (products Products) String() string {
	var builder strings.Builder
	for _, product := range products {
		builder.WriteString(fmt.Sprintf("%s\n", product))
	}
	return builder.String()
}

const (
	Id       = "Id"
	Name     = "Name"
	Cost     = "Cost"
	Units    = "Units"
	Category = "Category"
)

// utility method
func (products Products) Sort(attrName string) {
	switch attrName {
	case Id:
		sort.Slice(products, func(i, j int) bool {
			return products[i].Id < products[j].Id
		})
	case Name:
		sort.Slice(products, func(i, j int) bool {
			return products[i].Name < products[j].Name
		})
	case Cost:
		sort.Slice(products, func(i, j int) bool {
			return products[i].Cost < products[j].Cost
		})
	case Units:
		sort.Slice(products, func(i, j int) bool {
			return products[i].Units < products[j].Units
		})
	case Category:
		sort.Slice(products, func(i, j int) bool {
			return products[i].Category < products[j].Category
		})
	}
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}
	fmt.Println("Initial List")
	fmt.Println(products)

	fmt.Println("Sort by Id")
	// sort.Sort(products)
	products.Sort(Id)
	fmt.Println(products)

	fmt.Println("Sort by Name")
	// sort.Sort(ByName{products})
	products.Sort(Name)
	fmt.Println(products)

	fmt.Println("Sort by Cost")
	// sort.Sort(ByCost{products})
	products.Sort(Cost)
	fmt.Println(products)

	fmt.Println("Sort by Units")
	// sort.Sort(ByUnits{products})
	products.Sort(Units)
	fmt.Println(products)

	fmt.Println("Sort by Category")
	// sort.Sort(ByCategory{products})
	products.Sort(Category)
	fmt.Println(products)
}
