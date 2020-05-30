package main

import (
	"fmt"
)

var y int = 42
var z = "Sou uma variável string"

var w = `Também sou uma string`

var nome = `Berg disse, "Sou homem...uuiiii"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(w)
	fmt.Printf("%T\n", w)

	z = `Agora tenho crase`
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(nome)
	fmt.Printf(`%T\n`, nome)
}

// Quando se utiliza `` para strings ao invés de ""
// raw string
// string crua, pura...
// Não sofre nenhum tipo de tratamento e vai aceitar
// qualquer coisa que se coloque nela, incluindo espaçamento de linhas
