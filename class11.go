//rel_op     = "==" | "!=" | "<" | "<=" | ">" | ">=" .

package main

import (
	"fmt"
)

var x bool

func main() {
	a := 7
	b := 13

	x = (a == b)
	fmt.Println(x)
	x = (a != b)
	fmt.Println(x)
	x = (a > b)
	fmt.Println(x)
	x = (a < b)
	fmt.Println(x)
	x = (a >= b)
	fmt.Println(x)
	x = (a <= b)
	fmt.Println(x)

	fmt.Printf("%T\n", x)
}
