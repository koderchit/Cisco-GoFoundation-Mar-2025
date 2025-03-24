package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// iterating an slice
	fmt.Println("iterating an slice")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	// iterating an slice using 'range'
	fmt.Println("iterating an slice using 'range'")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	// append new items
	nos = append(nos, 10)
	fmt.Println(nos)

	nos = append(nos, 40, 20, 30)
	fmt.Println(nos)

	hundreds := []int{300, 100, 500}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	newNos := nos // behavior of assignment operation is different from an array.. why?
	newNos[0] = 9999
	fmt.Printf("newNos[0] = %d and nos[0] = %d\n", newNos[0], nos[0])

	fmt.Println("After sorting,")
	sort(nos)
	fmt.Println(nos)

	// slicing
	fmt.Println("Slicing")
	fmt.Println("nos[3:6] :", nos[3:6]) //3,4,5
	fmt.Println("nos[:6] :", nos[:6])   //0..5
	fmt.Println("nos[6:] :", nos[6:])   //6..end of the slice
}

func sort(list []int) /* no return */ {
	for i := 0; i < (len(list) - 1); i++ {
		for j := i + 1; j < len(list); j++ {
			// no need to dereference when accessing the members of a pointer to a collection/container
			if list[i] > list[j] {
				list[j], list[i] = list[i], list[j]
			}
		}
	}
}
