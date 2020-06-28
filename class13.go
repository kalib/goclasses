package main
import ("fmt"
)

var x int
var y float64

func main() {
	x = 28
	y = 34.537
	fmt.Println("x =", x)
	fmt.Println("y =", y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}