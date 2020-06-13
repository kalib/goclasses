package main

import (
	"fmt"
)

var a int
type bar int
var b bar
type cerva string
var c cerva

func main() {
	a = 42
	b = 43
	c = "44"

	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	a = int(b)
	//b = int(c) //C é originalmente "string" por tanto não pode ser convertida para "int"
	fmt.Println(a)
	fmt.Printf("%T\t%T\t%T\n", a, b, c)
	fmt.Println(b)
	fmt.Println(c)
}
