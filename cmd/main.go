package main

import (
	"fmt"
	"log"
	"net"

	"github.com/tweeker88/auth-oleg/internal/builder"
)

const grpcPort = 50051

func main() {
	srv := builder.BuildServer()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err = srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
