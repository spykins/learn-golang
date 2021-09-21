package spypackageone

import "fmt"

func SayHello() {
	fmt.Println("This is calling say hello from spypackageone")
	// return "Hello from spypackageone"
}

func sayHelloWithName() {
	fmt.Println("Hello from ")
	// return "Hello from " + name
}
