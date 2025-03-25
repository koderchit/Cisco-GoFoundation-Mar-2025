package main

import "fmt"

func main() {

	// var productRanks map[string]int = make(map[string]int)
	// var productRanks map[string]int = map[string]int{}
	// var productRanks map[string]int = map[string]int{"scribble-pad": 4, "pencil": 1, "marker": 3}

	// type inference
	// var productRanks = map[string]int{"scribble-pad": 4, "pencil": 1, "marker": 3}

	// use :=
	// productRanks := map[string]int{"scribble-pad": 4, "pencil": 1, "marker": 3}
	productRanks := map[string]int{
		"scribble-pad": 4,
		"pencil":       1,
		"marker":       3,
	}

	fmt.Println(productRanks)
	productRanks["pen"] = 5
	fmt.Println(productRanks)
	fmt.Printf("len(productRanks) = %d\n", len(productRanks))

	fmt.Println("Iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	// Removing an item
	// keyToRemove := "scribble-pad"
	// non-existent key
	keyToRemove := "fountain-pen"
	delete(productRanks, keyToRemove)
	fmt.Printf("After removing %q, productRanks = %v\n", keyToRemove, productRanks)

	// Check for the existence of a key
	// keyToCheck := "fountain-pen"
	keyToCheck := "pen"
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("Key %q does not exist\n", keyToCheck)
	}

}
