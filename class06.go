package main

import (
	"fmt"
)

var y string
var z int

func main() {
	fmt.Println("Imprimindo o valor de y", y, "e pronto.")
	fmt.Printf("%T\n", y)

	y = "Bond, James Bond"
	fmt.Println("Meu nome é", y, ">]")
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
