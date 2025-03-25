package main

import "fmt"

func main() {
	// s1 := []int{3, 1, 2}
	// s1 := []int{}

	// pre-allocate memory
	// s1 := make([]int, 0, 3)

	// initialize & pre-allocate memory
	s1 := make([]int, 2, 4)
	fmt.Printf("len(s1) = %d, cap(s1) = %d, s1 = %v\n", len(s1), cap(s1), s1)
	s1 = append(s1, 7)
	fmt.Printf("len(s1) = %d, cap(s1) = %d, s1 = %v\n", len(s1), cap(s1), s1)
	s1 = append(s1, 6)
	fmt.Printf("len(s1) = %d, cap(s1) = %d, s1 = %v\n", len(s1), cap(s1), s1)
	s1 = append(s1, 9)
	fmt.Printf("len(s1) = %d, cap(s1) = %d, s1 = %v\n", len(s1), cap(s1), s1)
	s1 = append(s1, 4)
	fmt.Printf("len(s1) = %d, cap(s1) = %d, s1 = %v\n", len(s1), cap(s1), s1)
	s1 = append(s1, 8)
	fmt.Printf("len(s1) = %d, cap(s1) = %d, s1 = %v\n", len(s1), cap(s1), s1)

	/*
		x1 := make([]int, 5, 8)
		x2 := x1[1:4]
		fmt.Printf("cap(x1) = %d, cap(x2) = %d\n", cap(x1), cap(x2))
	*/

}
