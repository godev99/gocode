package main

import (
	"fmt"
	"os"
)

func main() {

	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("index : %d	", i)
		fmt.Printf("value : %s\n", os.Args[i])
	}

}
