package main

import "fmt"

func sqlQuote(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x) // x has type interface{} here.
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return x // (not shown)
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}

}

func main() {
	fmt.Printf("%s\n", sqlQuote(uint(5)))
	fmt.Printf("%s\n", sqlQuote(true))
	fmt.Printf("%s\n", sqlQuote("5"))
	fmt.Printf("%s\n", sqlQuote(-5))
	fmt.Printf("%s\n", sqlQuote(nil))
	fmt.Printf("%s\n", sqlQuote([]string{"foo", "bar"}))
}
