// Quando se utiliza `` para strings ao invés de ""
// raw string
// string crua, pura...
// Não sofre nenhum tipo de tratamento e vai aceitar
// qualquer coisa que se coloque nela, incluindo espaçamento de linhas

package main

import "fmt"

var y = 42
var z = "Sou uma variável string"
var a = `Berg disse,
                "sou homem sim, mesmo que não pareça"
acredite se quiser...`
var b = "Sou\numa variável"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
}
