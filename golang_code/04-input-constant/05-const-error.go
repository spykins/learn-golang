package main

import "fmt"

func main() {
	const x string = "Hello world"
	fmt.Println(x)
	// x = "Some other values"   // error
	fmt.Println(x)
}
