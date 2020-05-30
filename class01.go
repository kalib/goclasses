package main

import (
	"fmt"
)

var y = 43

func main() {
	x := 42
	fmt.Println(x)

	fmt.Println(y)

	z := x + y
	fmt.Println(z)

	b()
	fmt.Println(x)
}

func b() {
	fmt.Println("Novamente, ", y)
}
