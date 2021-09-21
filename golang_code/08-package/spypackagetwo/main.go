package main

import (
	"../spypackageone"
	"fmt"
)

func main() {
	fmt.Println("This is the main function")
	spypackageone.SayHello()
	CheckMain()
}
