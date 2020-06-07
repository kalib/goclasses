package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const x, y, z = 42, "James Bond", true
	s := fmt.Sprintf("%s %d %d", x, y, z)

	io.WriteString(os.Stdout, s)

}
