package main

import (
	"errors"
	"fmt"
)

var ErrF1 = errors.New("f1 error")
var ErrF2 = errors.New("f2 error")

func main() {

	e := f1()
	if errors.Is(e, ErrF1) {
		fmt.Println("Error occured in f1()")
	}
	if errors.Is(e, ErrF2) {
		fmt.Println("Error occured in f2()")
	}
}

func f1() error {
	errF2 := f2()
	return fmt.Errorf("%w %w", ErrF1, errF2)
}

func f2() error {
	// return ErrF2
	return nil
}
