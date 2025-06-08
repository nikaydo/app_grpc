package grpc

import (
	"log"

	auth "github.com/nikaydo/grpc-contract/gen/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run() auth.AuthClient {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("grpc start")
	return auth.NewAuthClient(conn)
}
