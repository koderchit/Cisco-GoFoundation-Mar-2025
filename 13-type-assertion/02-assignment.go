/*
Hint : use strconv.Atoi() to convert string to int
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(10))                                            //=> 10
	fmt.Println(sum(10, 20))                                        //=> 30
	fmt.Println(sum(10, "20"))                                      //=> 30
	fmt.Println(sum(10, "20", "abc"))                               //=> 30
	fmt.Println(sum(10, 20, 30, 40))                                //=> 100
	fmt.Println(sum(10, 20, []int{30, 40}))                         //=> 100
	fmt.Println(sum(10, 20, []any{30, "40"}))                       //=> 100
	fmt.Println(sum(10, 20, []any{30, "40", "abc"}))                //=> 100
	fmt.Println(sum(10, 20, []any{30, "40", "abc", []int{50, 60}})) //=> 210
	fmt.Println(sum())                                              //=> 0
}

func sum( /* ? */ ) int {
	/* ? */
}
