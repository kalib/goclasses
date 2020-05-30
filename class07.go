package main

import (
	"fmt"
)

var y = 453

func main() {
	fmt.Println("Hello World")
	fmt.Println(y)
	fmt.Printf("%T\n", y)  // %T mostra o tipo
	fmt.Printf("%v\n", y)  // %v mostra o valor real da variável
	fmt.Printf("%x\n", y)  // %x mostra o valor em hexadecimal
	fmt.Printf("%X\n", y)  // %X mostra em hexadecimal com letras maiúsculas
	fmt.Printf("%#x\n", y) // %#x hexadecimal iniciando com 0x
	fmt.Printf("%b\n", y)  // %b mostra em binário

	fmt.Printf("%#x\n%b\n%x\n", y, y, y)

	fmt.Printf("Em hexadecimal com 0x %#x, em binário %b, e em hexadecimal %x\n", y, y, y)

	fmt.Printf("Esta é minha string\t\t%b\t%x\n", y, y)
}

// \n = pula linha
// \t = dá um tab
