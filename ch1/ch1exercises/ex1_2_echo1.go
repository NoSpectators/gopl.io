// modify the echo program to print the index and value of each of its arguments, one per line
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Println("index = ", i)
		fmt.Println("value = ", s)
	}
	//fmt.Println(s)
}

