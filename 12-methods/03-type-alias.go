package main

import "fmt"

type MyStr string // type alias

func (s MyStr) Length() int {
	return len(s)
}

func main() {
	str := MyStr("Amet ut dolor culpa ipsum fugiat non sint.")
	fmt.Println(str.Length())
}
