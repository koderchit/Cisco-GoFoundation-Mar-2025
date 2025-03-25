package calculator

/*
// private
var opCount int

// init
func init() {
	fmt.Println("[calculator/calc.go] - init invoked")
}

// public
func OpCount() int {
	return opCount
}
*/

// private
var opCount map[string]int

// init
func init() {
	opCount = make(map[string]int)
}

func OpCount() map[string]int {
	return opCount
}
