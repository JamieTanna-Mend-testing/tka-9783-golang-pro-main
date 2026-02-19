package main

import (
	"fmt"

	"pro-lib"
)

func main() {
	fmt.Println("Hello World")
	callPrint()
}

func callPrint() {
	lib.Printer()
}
