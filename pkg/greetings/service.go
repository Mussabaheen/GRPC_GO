package greetings

import (
	"context"
	"fmt"

	sh "github.com/Mussabaheen/GRPC_GO/hellodev"
	"google.golang.org/grpc"
)

type Service interface {
	SayHello() (*sh.HelloReply, error)
}

type service struct {
	grpc grpc.ClientConn
}

func NewService(conn grpc.ClientConn) Service {
	return &service{
		grpc: conn,
	}
}

func (s *service) SayHello() (*sh.HelloReply, error) {
	fmt.Println("Hello World")
	greeterClient := sh.NewGreeterClient(&s.grpc)
	resp, err := greeterClient.SayHello(context.TODO(), &sh.HelloRequest{
		Name: "Malik Mussabeheen",
	})
	if err != nil {
		return &sh.HelloReply{}, err
	}
	return resp, nil
}
