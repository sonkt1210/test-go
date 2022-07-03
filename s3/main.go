package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"test-go/s3/database"
	"test-go/s3/transactionpb"
)

var (
	conns   []net.Conn
	connCh  = make(chan net.Conn)
	closeCh = make(chan net.Conn)
	msgCh   = make(chan string)
)

type server struct{}

func main() {
	db := database.Init()
	fmt.Println("connect db", db)

	lis, err := net.Listen("tcp", "0.0.0.0:50069")
	if err != nil {
		log.Fatalf("create listen %v", err)
	}

	s := grpc.NewServer()

	transactionpb.RegisterTransactionServiceServer(s, &server{})

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("create server %v", err)
	}
}
