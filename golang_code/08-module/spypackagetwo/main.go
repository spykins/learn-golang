package main

import (
	"../spypackageone"
	"fmt"
)

// run with >> GO111MODULE=off go run .
// # this search module in the path locally,
// # when GO111MODULE=on it pulls it from internet, GO111MODULE=auto...
func main() {
	fmt.Println("This is the main function")
	spypackageone.SayHello()
	CheckMain()
}
