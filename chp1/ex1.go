package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf("command : %s\n", os.Args[0])
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("result : %s\n", s)
}
