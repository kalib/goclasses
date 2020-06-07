package main

import (
	"fmt"
)

	var x int = 42
	var y string = "James Bond"
	var z bool = true

func main() {
	fmt.Println("Valor de x: ", x)
	fmt.Println("Valor de y: ", y)
	fmt.Println("Valor de z: ", z)	
	fmt.Println("Como se chama o valor atribuido para cada variavel ? ")
	fmt.Println("Resposta: Zero")
	
	s := fmt.Sprintf("Valores: ",x,y,z)
        fmt.Println(s)	

}