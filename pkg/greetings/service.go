package greetings

import "fmt"

type Service interface {
	SayHello()
}

func SayHello() {
	fmt.Println("Hello World")
}
