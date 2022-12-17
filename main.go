package main

import (
	"go-grpc-practice/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	service.RegisterCalculatorServer(s, service.NewCalculatorServer())

	log.Println("gRPC server listening on port 50051")
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
