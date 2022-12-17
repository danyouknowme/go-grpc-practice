package main

import (
	"go-grpc-practice/service"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	credentials := insecure.NewCredentials()
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(credentials))
	if err != nil {
		log.Fatal(err)
	}

	calculatorClient := service.NewCalculatorClient(cc)
	calculatorService := service.NewCalculatorService(calculatorClient)

	err = calculatorService.Hello("Danny")
	if err != nil {
		log.Fatal(err)
	}
}
