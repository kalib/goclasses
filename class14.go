package main
import ("fmt"
)
//https://golang.org/ref/spec#Numeric_types

var x int //Utiliza o valor entre int32 e int64
var z int8 //int8 - the set of all signed  8-bit integers (-128 to 127)
var w int8
var y float64

func main() {
	x = 287890546
	z = -128
	w = 127
	y = 34.537
	fmt.Println("x =", x)
	fmt.Println("z =", z)
	fmt.Println("w =", w)
	fmt.Println("y =", y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", z)
	fmt.Printf("%T\n", w)
	fmt.Printf("%T\n", y)
}