package main

import (
	"fmt"
)

var nome string = "Daniel F. de Queiroz"
var idade int = 27
var peso float64 = 77.5
var pesopizza float64 = peso + 1

func main() {
	fmt.Println("Meu nome e", nome, "Estou aprendendo o básico de Go e git")
	fmt.Println("Possuo", idade, "anos e estou pesando aproximadamente", peso, "Kg.")
	fmt.Println("Utilizei variáveis para lhe passar estas informações pessoais com Go.")
	fmt.Println("Por exemplo, utilizei a variável nome, com o tipo: ")
	fmt.Printf("%T\n", nome)
	fmt.Println("Utilizei outras duas variáveis, para idade e peso, com os tipos: ")
	fmt.Printf("%T%v%T\n", idade, " e ", peso)
	fmt.Println("Gosto muito de pizza, portanto fui em um rodízio de pizzas ontem e acabei engordando 1 kg.")
	fmt.Println("Por isso agora estou pesando", pesopizza, "e utilizei a variavel do tipo: ")
	fmt.Printf("%T\n", pesopizza)
}
