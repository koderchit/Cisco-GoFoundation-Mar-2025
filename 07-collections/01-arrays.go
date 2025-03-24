package main

import "fmt"

func main() {
	// var nos [5]int // memory is allocated and initialized with the 'zero' value of 'int'
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos [5]int = [5]int{3, 1, 4}

	// type inference
	// var nos = [5]int{3, 1, 4, 2, 5}

	// use :=
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// iterating an array
	fmt.Println("iterating an array")
	for i := 0; i < 5; i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	// iterating an array using 'range'
	fmt.Println("iterating an array using 'range'")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	newNos := nos // arrays are values and hence the assignment creates a copy
	newNos[0] = 9999
	fmt.Printf("newNos[0] = %d, nos[0] = %d\n", newNos[0], nos[0])

	fmt.Println("After sorting")
	sort(&nos)
	fmt.Println(nos)
}

func sort(list *[5]int) /* no return */ {
	for i := 0; i < (5 - 1); i++ {
		for j := i + 1; j < 5; j++ {
			// no need to dereference when accessing the members of a pointer to a collection/container
			if list[i] > list[j] {
				list[j], list[i] = list[i], list[j]
			}
		}
	}
}
