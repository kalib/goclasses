package main

import (
	"fmt"
)

var y = 43
var bob int

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(y)

	z := x + y
	fmt.Println(z)

	b()
	fmt.Println(x)
	fmt.Println(bob)
}

func b() {
	fmt.Println("Novamente, ", y)
	fmt.Println(bob)
	bob = 10293473
	fmt.Println(bob)
}

// Valor ZERO de cada tipo:

// int = 0
// float = 0.0
// boolean = false
// string = ""
