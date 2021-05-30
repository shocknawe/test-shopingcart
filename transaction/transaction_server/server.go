package main

import (
	"log"
	"net"

	"github.com/shocknawe/test-shopingcart/transaction"
	"google.golang.org/grpc"
)

type server struct {
	transaction.UnimplementedTransactionServiceServer
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Unable to create listener", err)
	}

	s := grpc.NewServer()
	transaction.RegisterTransactionServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalln("Unable to create server", err)
	}
}
