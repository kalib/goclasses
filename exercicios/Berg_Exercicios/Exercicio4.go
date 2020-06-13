package main

import (
	"fmt"
)
	var nome string  = "Berg"
	var idade int = 39
	var peso float64 = 82.3
	var pesopizza int

func main() {
	    fmt.Println("Onde você encontrar <nome>, deverá ser impresso o valor da variável: ", nome)	
		fmt.Println("Onde você encontrar <idade>, deverá ser impresso o valor da variavel: ", idade)
		fmt.Println("Onde você encontrar <peso>, deverá ser impresso o valor da variavel: ", peso)
		fmt.Println("Onde você encontrar <pesopizza>, deverá ser impresso o valor da variavel: ", pesopizza)
		fmt.Println()
		fmt.Printf("Onde você encontrar <tipo_variável_nome>, deverá ser impresso o tipo da variavel:  %T\n ", nome)
		fmt.Printf("Onde você encontrar <tipo_variável_idade>, deverá ser impresso o tipo da variavel:  %T\n ", idade)
		fmt.Printf("Onde você encontrar <tipo_variável_peso>, deverá ser impresso o tipo da variavel:  %T\n ", peso)
		fmt.Printf("Onde você encontrar <tipo_variável_pesopizza>, deverá ser impresso o tipo da variavel:  %T\n ", pesopizza)
		fmt.Println()
		fmt.Println("Meu nome é ", nome,". Estou aprendendo o básico de Go e git.")
		fmt.Println("Possuo ", idade,"anos e estou pesando aproximadamente ", peso,".")
		fmt.Println("Utilizei variáveis para lhe passar estas informações pessoais com Go.")
		fmt.Printf("Por exemplo, utilizei a variável nome, com o tipo %T\n", nome)
		fmt.Printf("Utilizei outras duas variáveis, para idade e peso, com os tipos %T e %T ", idade, peso)
		fmt.Println()
		fmt.Println("Gosto muito de pizza, portanto fui em um rodízio de pizzas ontem e acabei engordando 1 kg.")
		var pesopizza = peso + 1
		fmt.Print(pesopizza)
		fmt.Printf(" %T" , pesopizza, )
}