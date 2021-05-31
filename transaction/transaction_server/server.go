package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/shocknawe/test-shopingcart/transaction"
	"google.golang.org/grpc"
)

type server struct {
	transaction.UnimplementedTransactionServiceServer
}

func (s *server) GetTotal(cts context.Context, request *transaction.TransactionRequest) (*transaction.TransactionResponse, error) {
	lineItems := request.GetTransaction().GetLineItems()
	fmt.Println(lineItems)

	response := &transaction.TransactionResponse{
		Total: 0,
	}

	return response, nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Unable to create listener", err)
	}
	log.Println("Listening to port 50051...")

	s := grpc.NewServer()
	transaction.RegisterTransactionServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalln("Unable to create server", err)
	}
}
