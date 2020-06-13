package main

import (
	"fmt"
)

var a int

type serie int

var b serie

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// Não poderia fazer essa operacao porque sao tipos distintos. Somente um teste.
	//a = b
	//fmt.Println(a)

}
