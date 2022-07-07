package greetings

import (
	"fmt"
)

type Service interface {
	SayHello()
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) SayHello() {
	fmt.Println("Hello World")

}
