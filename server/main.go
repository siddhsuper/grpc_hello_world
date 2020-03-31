package main

import (
	"context"

	"log"
	"net"

	grpc "google.golang.org/grpc"

	"github.com/grpc_hello_world/helloworld"
)

const (
	port = ":8888"
)

type server struct {
}

// SayHello implements helloworld.GreeterServer: Sends a greeting
func (s *server) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	log.Printf("Received: %v", req.GetName())
	return &helloworld.HelloResponse{Message: "Hello " + req.GetName()}, nil
}

func main() {
	srv := grpc.NewServer()
	var greeterServer server
	helloworld.RegisterGreeterServer(srv, &greeterServer)
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("could not listen to %s %v", port, err)
	}
	log.Println("GRPC Server Running")
	log.Fatal(srv.Serve(l))
}
